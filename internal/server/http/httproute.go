package http

import (
	"tugas_akhir_example/internal/pkg/books"

	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	booksAPI := api.Group("/books")
	books.BooksRoute(booksAPI, containerConf)

}
