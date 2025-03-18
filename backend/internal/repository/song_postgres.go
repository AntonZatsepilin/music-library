package repository

import (
	"log"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) CreateSong(song models.Song) error {
	log.Printf("Inserting song: %+v", song)
	query := "INSERT INTO songs (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	_, err := r.db.Exec(query, song.Group, song.SongName, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		log.Printf("Error inserting song: %v", err)
		return err
	}
	return nil
}
