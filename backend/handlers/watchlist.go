package handlers

import (
	"errors"
	"fmt"

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
func (m *module) CreateWatchlist(watchlist datatransfers.WatchlistCreate) (id uint, err error) {
	watchlistId, err := m.db.watchlistOrmer.InsertWatchlist(models.Watchlist{
		Name:         watchlist.Name,
		ListOfMovies: watchlist.ListOfMovies,
	})
	if err != nil {
		return -1, errors.New("cannot update movie")
	}
	return watchlistId, nil
}
