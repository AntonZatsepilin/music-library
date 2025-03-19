package models

type Song struct {
    ID          int    `db:"id" json:"id"`
    Group       string `db:"group_name" json:"group" binding:"required"`
    SongName    string `db:"song_name" json:"song" binding:"required"`
    ReleaseDate string `db:"release_date" json:"releaseDate"`
    Text        string `db:"text" json:"text"`
    Link        string `db:"link" json:"link"`
}

type SongDetail struct {
    ReleaseDate string `json:"releaseDate"`
    Text        string `json:"text"`
    Link        string `json:"link"`
}

type CreateSongRequest struct {
    Group string `json:"group"`
    Song  string `json:"song"`
}