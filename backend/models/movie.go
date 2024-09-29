package models

import (
	"gorm.io/gorm"
	"time"
)

type orm struct {
	db *gorm.DB
}

type Movie struct {
	ID          uint      `gorm:"primaryKey" json:"movie_id"`
	Title       string    `json:"title" json:"title"`
	Description string    `json:"description" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"-"`
}

type MovieOrmer interface {
	GetOneByID(id uint) (movie Movie, err error)
	GetOneByName(name string) (movie Movie, err error)
	InsertMovie(movie Movie) (id uint, err error)
	UpdateMovie(movie Movie) (err error)
}

func NewMovieOrmer(db *gorm.DB) MovieOrmer {
	//_ = db.AutoMigrate(&Movie{})		// builds table when enabled
	return &orm{db}
}

func (o *orm) GetOneByID(id uint) (movie Movie, err error) {
	result := o.db.Model(&Movie{}).Where("id = ?", id).First(&movie)
	return movie, result.Error
}

func (o *orm) GetOneByName(name string) (movie Movie, err error) {
	result := o.db.Model(&Movie{}).Where("name = ?", name).First(&movie)
	return movie, result.Error
}

func (o *orm) InsertMovie(movie Movie) (id uint, err error) {
	result := o.db.Model(&Movie{}).Create(&movie)
	return movie.ID, result.Error
}

func (o *orm) UpdateMovie(movie Movie) (err error) {
	// By default, only non-empty fields are updated. See https://gorm.io/docs/update.html#Updates-multiple-columns
	result := o.db.Model(&Movie{}).Model(&movie).Updates(&movie)
	return result.Error
}
