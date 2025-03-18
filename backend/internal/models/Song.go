package models

import (
	"time"
)

type Song struct {
	ID          int       `db:"id" json:"id"`
	Group       string    `db:"group" json:"group" binding:"required"`
	SongName    string    `db:"song" json:"song" binding:"required"`
	ReleaseDate string    `db:"release_date" json:"releaseDate"`
	Text        string    `db:"text" json:"-"`
	Link        string    `db:"link" json:"link"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

type SongLyricsPage struct {
	Verses []string `json:"verses"`
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
	Total  int      `json:"total"`
}