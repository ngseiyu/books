package database

import (
	"book-management/internal/models"

	"gorm.io/gorm"
)

func populateBooks(books []models.Book, db *gorm.DB) error {

	if len(books) == 0 {
		return nil
	}

	return db.CreateInBatches(books, len(books)).Error
}

func GetBooks(db *gorm.DB) ([]models.BookList, error) {

	books := []models.BookList{}
	result := db.Model(&models.Book{}).
		Order("id DESC").
		Find(&books).
		Error

	return books, result
}

func GetBook(db *gorm.DB, id int) (models.BookDetails, error) {

	book := models.BookDetails{}
	result := db.Model(&models.Book{}).
		Where("id = ?", id).
		First(&book).
		Error

	return book, result
}

func CreateBook(db *gorm.DB, book *models.Book) error {

	return db.Model(&models.Book{}).
		Create(&book).
		Error
}

func UpdateBook(db *gorm.DB, book models.Book) error {

	return db.Model(&models.Book{}).
		Where("id = ?", book.ID).
		Updates(book).
		Error
}

func DeleteBook(db *gorm.DB, id int) error {

	return db.Model(&models.Book{}).
		Where("id = ?", id).
		Unscoped().
		Delete(&models.Book{}).
		Error
}
