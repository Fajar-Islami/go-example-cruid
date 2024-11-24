package handler

import (
	"github.com/gofiber/fiber/v2"

	authcontroller "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
)

func AuthRoute(r fiber.Router, UserUsc usecase.UsersUseCase) {
	controller := authcontroller.NewAuthController(UserUsc)

	booksAPI := r.Group("/auth")
	booksAPI.Post("/register", controller.Register)
	booksAPI.Post("/login", controller.Login)
}
