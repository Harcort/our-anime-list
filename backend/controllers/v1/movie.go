package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"our-anime-list/backend/constants"
	"our-anime-list/backend/datatransfers"
	"our-anime-list/backend/handlers"
	"our-anime-list/backend/models"
)

func GETMovie(c *gin.Context) {
	var err error
	var movieInfo datatransfers.MovieInfo
	if err = c.ShouldBindUri(&movieInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	var movie models.Movie
	if movie, err = handlers.Handler.RetrieveMovie(movieInfo.Title); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find movie"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: datatransfers.MovieInfo{
		Title:       movie.Title,
		Description: movie.Description,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}})
}

func PUTMovie(c *gin.Context) {
	var err error
	var movie datatransfers.MovieUpdate
	if err = c.ShouldBind(&movie); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.UpdateMovie(uint(c.GetInt(constants.IsAuthenticatedKey)), movie); err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Error: "failed updating movie"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: movie})
}

func POSTMovie(c *gin.Context) {
	var err error
	var movie datatransfers.MovieCreate
	if err = c.ShouldBind(&movie); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	_, err = handlers.Handler.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Error: "failed updating movie"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: movie})
}
