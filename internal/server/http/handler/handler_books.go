package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	bookscontroller "tugas_akhir_example/internal/pkg/controller"
)

func BooksRoute(r fiber.Router, containerConf *container.Container) {
	controller := bookscontroller.NewBooksController(containerConf.BooksUsc)

	booksAPI := r.Group("/books")
	booksAPI.Get("", controller.GetAllBooks)
	booksAPI.Get("/:id_books", controller.GetBooksByID)
	booksAPI.Post("", MiddlewareAuth, controller.CreateBooks)
	booksAPI.Put("/:id_books", controller.UpdateBooksByID)
	booksAPI.Delete("/:id_books", controller.DeleteBooksByID)
}
