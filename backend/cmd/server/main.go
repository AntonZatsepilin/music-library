package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/AntonZatsepilin/music-library.git/docs"
	"github.com/AntonZatsepilin/music-library.git/internal/handler"
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
	"github.com/AntonZatsepilin/music-library.git/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           Music Library API
// @version         1.0.0
// @description     API for managing music library
// @host            localhost:8080
// @BasePath        /
func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	musicInfoAPI := viper.GetString("musicInfoAPI")

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	infoClient := service.NewMusicInfoClient(musicInfoAPI)
	services := service.NewService(repos, infoClient)
	handlers := handler.NewHandler(services)

	srv := new(models.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("music-library-app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logrus.Print("music-library-app Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
