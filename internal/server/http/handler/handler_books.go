package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	bookscontroller "tugas_akhir_example/internal/pkg/controller"

	booksrepository "tugas_akhir_example/internal/pkg/repository"

	booksusecase "tugas_akhir_example/internal/pkg/usecase"
)

func BooksRoute(r fiber.Router, containerConf *container.Container) {
	repo := booksrepository.NewBooksRepository(containerConf.Mysqldb)
	usecase := booksusecase.NewBooksUseCase(repo)
	controller := bookscontroller.NewBooksController(usecase)

	booksAPI := r.Group("/books")
	booksAPI.Get("", controller.GetAllBooks)
	booksAPI.Get("/:id_books", controller.GetBooksByID)
	booksAPI.Post("", controller.CreateBooks)
	booksAPI.Put("/:id_books", controller.UpdateBooksByID)
	booksAPI.Delete("/:id_books", controller.DeleteBooksByID)
}
