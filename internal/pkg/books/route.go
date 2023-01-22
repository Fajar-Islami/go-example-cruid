package books

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	bookscontroller "tugas_akhir_example/internal/pkg/books/controller"

	booksrepository "tugas_akhir_example/internal/pkg/books/repository"

	booksusecase "tugas_akhir_example/internal/pkg/books/usecase"
)

func BooksRoute(r fiber.Router, containerConf *container.Container) {
	repo := booksrepository.NewBooksRepository(containerConf.Mysqldb)
	usecase := booksusecase.NewBooksUseCase(repo)
	controller := bookscontroller.NewBooksController(usecase)

	r.Get("", controller.GetAllBooks)
	r.Get("/:id_books", controller.GetBooksByID)
	r.Post("", controller.CreateBooks)
	r.Put("/:id_books", controller.UpdateBooksByID)
	r.Delete("/:id_books", controller.DeleteBooksByID)
}
