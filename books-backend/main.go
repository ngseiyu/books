package main

import (
	"book-management/internal/database"
	"book-management/internal/handlers"
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed views/*
var staticFiles embed.FS

func main() {

	if err := database.InitDatabase(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")
	books := api.Group("/books")
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	books.Post("/", handlers.CreateBook)
	books.Post("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)

	staticServer := filesystem.New(filesystem.Config{
		Root:       http.FS(staticFiles),
		PathPrefix: "views",
		Browse:     true,
	})
	app.Use("/", staticServer)

	// app.Static("/", "./views")
	app.Get("/*", func(c *fiber.Ctx) error {
		// return c.SendFile("./views/index.html")
		indexFile, err := staticFiles.ReadFile("views/index.html")
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("could not load index.html")
		}
		return c.Type("html").Send(indexFile)
	})

	log.Fatal(app.Listen(":8080"))
}
