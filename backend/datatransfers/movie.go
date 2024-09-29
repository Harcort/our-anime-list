package datatransfers

import (
	"time"
)

type MovieUpdate struct {
	Title       string `uri:"watchlist" json:"title"`
	Description string `json:"description"`
}

type MovieInfo struct {
	Title       string    `uri:"watchlist" json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MovieCreate struct {
	Title       string `uri:"watchlist" json:"title"`
	Description string `json:"description"`
}
