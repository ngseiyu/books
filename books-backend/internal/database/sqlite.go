package database

import (
	"book-management/internal/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {

	var err error

	_, err = os.Stat("app.db")
	firstConn := os.IsNotExist(err)

	DB, err = gorm.Open(sqlite.Open("app.db"))
	if err != nil {
		return errors.New("failed to connect to the database: " + err.Error())
	}

	if err := migrate(DB); err != nil {
		return errors.New("failed to initialise the database: " + err.Error())
	}

	// retrieve dummy data and populate database if is first time connecting / database not found
	// can replace with running a sql script
	if firstConn {
		resp, err := http.Get("https://fakerapi.it/api/v1/books")
		if err != nil {
			return errors.New("failed to fetch books: " + err.Error())
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.New("failed to read books: " + err.Error())
		}

		respData := struct {
			Status string        `json:"status"`
			Code   int           `json:"code"`
			Data   []models.Book `json:"data"`
		}{}

		// books := []models.Book{}
		if err := json.Unmarshal(respBody, &respData); err != nil {
			return errors.New("failed to parse book data: " + err.Error())
		}

		if err := populateBooks(respData.Data, DB); err != nil {
			return errors.New("failed to populdate book data: " + err.Error())
		}
	}

	return nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Book{},
	)
}
