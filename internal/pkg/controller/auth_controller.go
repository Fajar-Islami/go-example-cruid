package controller

import (
	authmodel "tugas_akhir_example/internal/pkg/model"
	authUsc "tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authUsc authUsc.UsersUseCase
}

func NewAuthController(authUsc authUsc.UsersUseCase) AuthController {
	return &AuthControllerImpl{
		authUsc: authUsc,
	}
}

func (uc *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(authmodel.Login)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// @TODO IMRPOVE FORMAT RESPONSE
	res, err := uc.authUsc.Login(c, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// @TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(authmodel.CreateUser)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.authUsc.CreateUsers(c, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// @TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
	})
}
