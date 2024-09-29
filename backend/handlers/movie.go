package handlers

import (
	"errors"
	"fmt"

	"our-anime-list/backend/datatransfers"
	"our-anime-list/backend/models"
)

func (m *module) RetrieveMovie(name string) (movie models.Movie, err error) {
	if movie, err = m.db.movieOrmer.GetOneByName(name); err != nil {
		return models.Movie{}, fmt.Errorf("cannot find movie with name %s", name)
	}
	return
}

func (m *module) UpdateMovie(id uint, movie datatransfers.MovieUpdate) (err error) {
	if err = m.db.movieOrmer.UpdateMovie(models.Movie{
		ID:          id,
		Title:       movie.Title,
		Description: movie.Description,
	}); err != nil {
		return errors.New("cannot update movie")
	}
	return
}

func (m *module) CreateMovie(id uint, movie datatransfers.MovieUpdate) (err error) {
	if err = m.db.movieOrmer.UpdateMovie(models.Movie{
		ID:          id,
		Title:       movie.Title,
		Description: movie.Description,
	}); err != nil {
		return errors.New("cannot create movie")
	}
	return
}
