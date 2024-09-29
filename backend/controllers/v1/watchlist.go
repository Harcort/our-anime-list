package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"our-anime-list/backend/constants"
	"our-anime-list/backend/datatransfers"
	"our-anime-list/backend/handlers"
	"our-anime-list/backend/models"
)

func GETWatchlist(c *gin.Context) {
	var err error
	var watchlistInfo datatransfers.WatchlistInfo
	if err = c.ShouldBindUri(&watchlistInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	var watchlist models.Watchlist
	if watchlist, err = handlers.Handler.RetrieveWatchlist(watchlistInfo.Name); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Error: "cannot find watchlist"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: datatransfers.WatchlistInfo{
		Name:         watchlist.Name,
		ListOfMovies: watchlist.ListOfMovies,
		CreatedAt:    watchlist.CreatedAt,
		UpdatedAt:    watchlist.UpdatedAt,
	}})
}

func PUTWatchlist(c *gin.Context) {
	var err error
	var watchlist datatransfers.WatchlistUpdate
	if err = c.ShouldBind(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.UpdateWatchlist(uint(c.GetInt(constants.IsAuthenticatedKey)), watchlist); err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Error: "failed updating watchlist"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: watchlist})
}

func POSTWatchlist(c *gin.Context) {
	var err error
	var watchlist datatransfers.WatchlistCreate
	if err = c.ShouldBind(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	_, err = handlers.Handler.CreateWatchlist(watchlist)
	if err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Error: "failed updating watchlist"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: watchlist})
}
