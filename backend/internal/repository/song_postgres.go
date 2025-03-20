package repository

import (
	"fmt"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) CreateSong(song models.Song) error {
	logrus.WithFields(logrus.Fields{
        "group": song.Group,
        "song":  song.SongName,
    }).Debug("Inserting a song into the database")
	query := "INSERT INTO songs (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	_, err := r.db.Exec(query, song.Group, song.SongName, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		logrus.WithError(err).Error("Error inserting song")
		return err
	}

	logrus.Info("The song has been successfully saved")
	return nil
}

func (r *SongPostgres) DeleteSongById(id int) error {
	logrus.WithField("id", id).Debug("Deleting a song from the database")
	query := "DELETE FROM songs WHERE id = $1"
    result, err := r.db.Exec(query, id)
    if err != nil {
        logrus.WithError(err).Error("Database error during deletion")
        return err
    }
    
    affected, _ := result.RowsAffected()
    if affected == 0 {
        notFoundErr := fmt.Errorf("song with id %d not found", id)
        logrus.WithError(notFoundErr).Warn("Attempt to delete non-existing song")
        return notFoundErr
    }
    
    logrus.Info("Song successfully deleted")
    return nil
}