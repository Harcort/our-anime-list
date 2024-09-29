package router

import (
	"github.com/gin-gonic/gin"

	"our-anime-list/backend/controllers/middleware"
	"our-anime-list/backend/controllers/v1"
	"our-anime-list/backend/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin)
			auth.POST("/signup", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1.GETUser)
			user.PUT("", utils.AuthOnly, v1.PUTUser)
		}
		watchlist := v1route.Group("/watchlist")
		{
			watchlist.GET(":id", utils.AuthOnly, v1.GETWatchlist)
			watchlist.PUT("", utils.AuthOnly, v1.PUTWatchlist)
			watchlist.POST("", utils.AuthOnly, v1.POSTWatchlist)

		}
		movie := v1route.Group("/movie")
		{
			movie.GET(":id", utils.AuthOnly, v1.GETMovie)
			movie.PUT("", utils.AuthOnly, v1.PUTMovie)
			movie.POST("", utils.AuthOnly, v1.POSTMovie)
		}

	}
	return
}
