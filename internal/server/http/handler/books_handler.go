package handler

import (
	"github.com/gofiber/fiber/v2"

	bookscontroller "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
)

func BooksRoute(r fiber.Router, BooksUsc usecase.BooksUseCase) {
	controller := bookscontroller.NewBooksController(BooksUsc)

	booksAPI := r.Group("/books")
	booksAPI.Get("", controller.GetAllBooks)
	booksAPI.Get("/:id_books", controller.GetBooksByID)
	booksAPI.Post("", MiddlewareAuth, controller.CreateBooks)
	booksAPI.Put("/:id_books", controller.UpdateBooksByID)
	booksAPI.Delete("/:id_books", controller.DeleteBooksByID)
}
