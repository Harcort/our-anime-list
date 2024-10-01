package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"

	"our-anime-list/backend/models"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestSelect(t *testing.T) {
	db, mock := NewMockDB()

	//err := handlers.MigrateTables(db)
	//if err != nil {
	//	t.Fatalf("Error in migrating tables: %v", err)
	//}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "John Doe")

	mock.ExpectQuery("^SELECT (.+) FROM \"watchlists\"").WillReturnRows(rows)

	var watchList []models.Watchlist
	if err := db.Find(&watchList).Error; err != nil {
		t.Fatalf("Error in finding users: %v", err)
	}

	if len(watchList) != 1 || watchList[0].Name != "John Doe" {
		t.Fatalf("Unexpected user data retrieved: %v", watchList)
	}
}
