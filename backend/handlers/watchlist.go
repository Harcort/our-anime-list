package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"our-anime-list/backend/datatransfers"
	"our-anime-list/backend/models"
)

func (m *module) RetrieveWatchlist(name string) (watchlist models.Watchlist, err error) {
	if watchlist, err = m.db.watchlistOrmer.GetOneByName(name); err != nil {
		return models.Watchlist{}, fmt.Errorf("cannot find watchlist with name %s", name)
	}
	return
}

func (m *module) UpdateWatchlist(id uint, watchlist datatransfers.WatchlistUpdate) (err error) {
	if err = m.db.watchlistOrmer.UpdateWatchlist(models.Watchlist{
		ID:           id,
		Name:         watchlist.Name,
		ListOfMovies: watchlist.ListOfMovies,
	}); err != nil {
		return errors.New("cannot update movie")
	}
	return
}

func (m *module) CreateWatchlist(c *gin.Context, watchlist datatransfers.WatchlistCreate) (id uint, err error) {
	var movieIDs []uint
	for _, movie := range watchlist.ListOfMovies {
		movieIDs = append(movieIDs, movie.ID)
	}
	movies, err := m.db.movieOrmer.GetMoviesByIds(movieIDs)
	watchlistId, err := m.db.watchlistOrmer.InsertWatchlist(models.Watchlist{
		Name:         watchlist.Name,
		ListOfMovies: movies,
		UserID:       c.GetUint("user_id"),
	})
	if err != nil {
		return 0, errors.New("cannot update movie")
	}
	return watchlistId, nil
}
