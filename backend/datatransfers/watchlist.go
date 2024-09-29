package datatransfers

import (
	"our-anime-list/backend/models"
	"time"
)

type WatchlistUpdate struct {
	Name         string         `uri:"watchlist" json:"name"`
	ListOfMovies []models.Movie `json:"movies"`
}

type WatchlistInfo struct {
	Name         string         `uri:"watchlist" json:"name"`
	ListOfMovies []models.Movie `json:"movies"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type WatchlistCreate struct {
	Name         string         `uri:"watchlist" json:"name"`
	ListOfMovies []models.Movie `json:"movies"`
}
