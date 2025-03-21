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

type UpdateSongRequest struct {
    Group       string `json:"group" db:"group_name"`
    Song        string `json:"song" db:"song_name"`
    ReleaseDate string `json:"releaseDate" db:"release_date"`
    Text        string `json:"text" db:"text"`
    Link        string `json:"link" db:"link"`
}

type LyricsRequest struct {
    Page  int `json:"page"`
    Limit int `json:"limit"`
}

type LyricResponse struct {
    Verses []string `json:"verses"`
    Total  int      `json:"total"`
    Page   int      `json:"page"`
    Limit  int      `json:"limit"`
}

type SongFilter struct {
    Group       *string `form:"group"`
    Song        *string `form:"song"`
    ReleaseDate *string `form:"releaseDate"`
    Text        *string `form:"text"`
    Link        *string `form:"link"`
}

type SongsResponse struct {
    Data  []Song `json:"data"`
    Total int    `json:"total"`
    Page  int    `json:"page"`
    Limit int    `json:"limit"`
}