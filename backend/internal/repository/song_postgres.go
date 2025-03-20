package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

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

	_, err := r.GetSongById(id)
	if err != nil {
		return err
	}
	
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

func (r *SongPostgres) UpdateSongById(id int, input models.UpdateSongRequest) error {

	_, err := r.GetSongById(id)
	if err != nil {
		return err
	}


	logrus.WithField("id", id).Debug("Updating a song in the database")

	queryGroupName := "UPDATE songs SET group_name=$1 WHERE id=$2"
	querySongName := "UPDATE songs SET song_name=$1 WHERE id=$2"
	queryReleaseDate := "UPDATE songs SET release_date=$1 WHERE id=$2"
	queryText := "UPDATE songs SET text=$1 WHERE id=$2"
	queryLink := "UPDATE songs SET link=$1 WHERE id=$2"

	wg := sync.WaitGroup{}

	if input.Group != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := r.db.Exec(queryGroupName, input.Group, id)
			if err != nil {
				logrus.WithError(err).Error("Error updating group name")
			}

			affected, _ := result.RowsAffected()
			if affected == 0 {
				notFoundErr := fmt.Errorf("1: song with id %d not found", id)
				logrus.WithError(notFoundErr).Warn("1: Attempt to update non-existing song")
			}
			
		}()
	}

		if input.Song != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := r.db.Exec(querySongName, input.Song, id)
			if err != nil {
				logrus.WithError(err).Error("Error updating song name")
			}
			affected, _ := result.RowsAffected()
			if affected == 0 {
				notFoundErr := fmt.Errorf("2: song with id %d not found", id)
				logrus.WithError(notFoundErr).Warn("2: Attempt to update non-existing song")
			}
		}()
		}

		if input.ReleaseDate != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := r.db.Exec(queryReleaseDate, input.ReleaseDate, id)
			if err != nil {
				logrus.WithError(err).Error("Error updating release date")
			}
			affected, _ := result.RowsAffected()
			if affected == 0 {
				notFoundErr := fmt.Errorf("3: song with id %d not found", id)
				logrus.WithError(notFoundErr).Warn("3: Attempt to update non-existing song")
			}
		}()
	}

	if input.Text != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := r.db.Exec(queryText, input.Text, id)
			if err != nil {
				logrus.WithError(err).Error("Error updating text")
			}
			affected, _ := result.RowsAffected()
			if affected == 0 {
				notFoundErr := fmt.Errorf("4: song with id %d not found", id)
				logrus.WithError(notFoundErr).Warn("4: Attempt to update non-existing song")				
			}
		}()
	}

	if input.Link != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := r.db.Exec(queryLink, input.Link, id)
			if err != nil {
				logrus.WithError(err).Error("Error updating link")
			}
			affected, _ := result.RowsAffected()
			if affected == 0 {
				notFoundErr := fmt.Errorf("5: song with id %d not found", id)
				logrus.WithError(notFoundErr).Warn("5: Attempt to update non-existing song")
			}
		}()
	}

	wg.Wait()

	return nil
}

func (r *SongPostgres) GetSongById(id int) (models.Song, error) {
    var song models.Song
    err := r.db.Get(&song, "SELECT * FROM songs WHERE id = $1", id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return song, fmt.Errorf("song with id %d not found", id)
        }
        return song, err
    }
    return song, nil
}

