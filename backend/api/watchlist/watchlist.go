package watchlist

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	MovieID     int    `json:"movie_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Watchlist struct {
	WatchlistID  int     `json:"watchlist_id"`
	UserID       int     `json:"user_id"`
	ListOfMovies []Movie `json:"list_of_movies"`
}

func GetWatchlist(writer http.ResponseWriter, request *http.Request) {
	var req Watchlist
	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		fmt.Printf("Error decoding request: %v", err)
	}
	exampleWatchlist := Watchlist{
		WatchlistID: 1,
		UserID:      1,
		ListOfMovies: []Movie{
			{
				MovieID:     1,
				Title:       "The Dark Knight",
				Description: "A movie about Batman",
			},
		},
	}
	err = json.NewEncoder(writer).Encode(exampleWatchlist)
	if err != nil {
		fmt.Printf("Error encoding response: %v", err)
	}
}
