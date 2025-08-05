package handlers

import (
	"book-management/internal/database"
	"book-management/internal/models"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBooks(c *fiber.Ctx) error {

	db := database.DB

	books := []models.BookList{}
	books, err := database.GetBooks(db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(books)
}

func GetBook(c *fiber.Ctx) error {

	db := database.DB

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid parameters",
		})
	}

	book := models.BookDetails{}
	book, err = database.GetBook(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "book id not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {

	db := database.DB

	book := models.Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "unable to parse request body",
		})
	}

	if err := database.CreateBook(db, &book); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(strings.ToLower(err.Error()), "unique constraint failed") {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"error": "book title already exists",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {

	db := database.DB

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid parameters",
		})
	}

	bookUpdate := models.Book{}
	if err := c.BodyParser(&bookUpdate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "unable to parse request body",
		})
	}
	bookUpdate.ID = uint(id)
	if err := database.UpdateBook(db, bookUpdate); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"error": "book id not found",
			})
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(strings.ToLower(err.Error()), "unique constraint failed") {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"error": "book title already exists",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	book := models.BookDetails{}
	book, err = database.GetBook(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "book id not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {

	db := database.DB

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid parameters",
		})
	}

	if err := database.DeleteBook(db, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(id)
}
