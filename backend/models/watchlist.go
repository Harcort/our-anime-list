package models

import (
	"gorm.io/gorm"
	"time"
)

type watchlistOrm struct {
	db *gorm.DB
}

type Watchlist struct {
	ID           uint      `gorm:"primaryKey" json:"-"`
	Name         string    `gorm:"uniqueIndex" json:"name"`
	ListOfMovies []Movie   `gorm:"many2many:watchlist_movie;" json:"listOfMovies"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"-"`
}

type WatchlistOrmer interface {
	GetOneByID(id uint) (watchlist Watchlist, err error)
	GetOneByName(name string) (watchlist Watchlist, err error)
	InsertWatchlist(watchlist Watchlist) (id uint, err error)
	UpdateWatchlist(watchlist Watchlist) (err error)
}

func NewWatchlistOrmer(db *gorm.DB) WatchlistOrmer {
	//_ = db.AutoMigrate(&Watchlist{})		// builds table when enabled
	return &watchlistOrm{db}
}

func (o *watchlistOrm) GetOneByID(id uint) (watchlist Watchlist, err error) {
	result := o.db.Model(&Watchlist{}).Where("id = ?", id).First(&watchlist)
	return watchlist, result.Error
}

func (o *watchlistOrm) GetOneByName(name string) (watchlist Watchlist, err error) {
	result := o.db.Model(&Watchlist{}).Where("name = ?", name).First(&watchlist)
	return watchlist, result.Error
}

func (o *watchlistOrm) InsertWatchlist(watchlist Watchlist) (id uint, err error) {
	result := o.db.Model(&Watchlist{}).Create(&watchlist)
	return watchlist.ID, result.Error
}

func (o *watchlistOrm) UpdateWatchlist(watchlist Watchlist) (err error) {
	// By default, only non-empty fields are updated. See https://gorm.io/docs/update.html#Updates-multiple-columns
	result := o.db.Model(&Watchlist{}).Model(&watchlist).Updates(&watchlist)
	return result.Error
}
