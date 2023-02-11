package controller

import (
	"log"
	booksdto "tugas_akhir_example/internal/pkg/dto"
	booksusecase "tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type BooksController interface {
	GetAllBooks(ctx *fiber.Ctx) error
	GetBooksByID(ctx *fiber.Ctx) error
	CreateBooks(ctx *fiber.Ctx) error
	UpdateBooksByID(ctx *fiber.Ctx) error
	DeleteBooksByID(ctx *fiber.Ctx) error
}

type BooksControllerImpl struct {
	booksusecase booksusecase.BooksUseCase
}

func NewBooksController(booksusecase booksusecase.BooksUseCase) BooksController {
	return &BooksControllerImpl{
		booksusecase: booksusecase,
	}
}

func (uc *BooksControllerImpl) GetAllBooks(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(booksdto.BooksFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.booksusecase.GetAllBooks(c, booksdto.BooksFilter{
		Title: filter.Title,
		Limit: filter.Limit,
		Page:  filter.Page,
	})

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) GetBooksByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	booksid := ctx.Params("id_books")
	if booksid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.booksusecase.GetBooksByID(c, booksid)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) CreateBooks(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(booksdto.BooksReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.booksusecase.CreateBooks(c, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) UpdateBooksByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	booksid := ctx.Params("id_books")
	if booksid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	data := new(booksdto.BooksReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.booksusecase.UpdateBooksByID(c, booksid, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) DeleteBooksByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	booksid := ctx.Params("id_books")
	if booksid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.booksusecase.DeleteBooksByID(c, booksid)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}
