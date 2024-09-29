package handlers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"our-anime-list/backend/config"
	"our-anime-list/backend/datatransfers"
	"our-anime-list/backend/models"
)

var Handler HandlerFunc

type HandlerFunc interface {
	AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error)
	RegisterUser(credentials datatransfers.UserSignup) (err error)

	RetrieveUser(username string) (user models.User, err error)
	UpdateUser(id uint, user datatransfers.UserUpdate) (err error)

	CreateWatchlist(watchlist datatransfers.WatchlistCreate) (id uint, err error)
	RetrieveWatchlist(name string) (watchlist models.Watchlist, err error)
	UpdateWatchlist(id uint, watchlist datatransfers.WatchlistUpdate) (err error)

	//CreateMovie(movie datatransfers.MovieCreate) (id uint, err error)
	RetrieveMovie(name string) (movie models.Movie, err error)
	UpdateMovie(id uint, movie datatransfers.MovieUpdate) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn           *gorm.DB
	userOrmer      models.UserOrmer
	watchlistOrmer models.WatchlistOrmer
	movieOrmer     models.MovieOrmer
}

func migrateTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Watchlist{},
	)
	return err
}

func InitializeHandler() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
		return
	}
	err = migrateTables(db)
	if err != nil {
		log.Println("[INIT] failed migrating tables")
		return
	}

	log.Println("[INIT] connected to PostgreSQL")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:           db,
			userOrmer:      models.NewUserOrmer(db),
			watchlistOrmer: models.NewWatchlistOrmer(db),
			movieOrmer:     models.NewMovieOrmer(db),
		},
	}
	return
}
