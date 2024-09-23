package server

import "our-anime-list/backend/api/watchlist"

func Router() {
	router.HandleFunc("/watchlist", watchlist.GetWatchlist).Methods("POST")
}
