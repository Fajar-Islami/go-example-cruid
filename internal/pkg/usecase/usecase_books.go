package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	booksdto "tugas_akhir_example/internal/pkg/dto"
	booksrepository "tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentfilepath = "internal/pkg/usecase/usecase.go"

type BooksUseCase interface {
	GetAllBooks(ctx context.Context, params booksdto.BooksFilter) (res []booksdto.BooksResp, err *helper.ErrorStruct)
	GetBooksByID(ctx context.Context, booksid string) (res booksdto.BooksResp, err *helper.ErrorStruct)
	CreateBooks(ctx context.Context, data booksdto.BooksReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateBooksByID(ctx context.Context, booksid string, data booksdto.BooksReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteBooksByID(ctx context.Context, booksid string) (res string, err *helper.ErrorStruct)
}

type BooksUseCaseImpl struct {
	booksrepository booksrepository.BooksRepository
}

func NewBooksUseCase(booksrepository booksrepository.BooksRepository) BooksUseCase {
	return &BooksUseCaseImpl{
		booksrepository: booksrepository,
	}

}

func (alc *BooksUseCaseImpl) GetAllBooks(ctx context.Context, params booksdto.BooksFilter) (res []booksdto.BooksResp, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.booksrepository.GetAllBooks(ctx, daos.FilterBooks{
		Limit:  params.Limit,
		Offset: params.Page,
		Title:  params.Title,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Books"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, booksdto.BooksResp{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Author:      v.Author,
		})
	}

	return res, nil
}
func (alc *BooksUseCaseImpl) GetBooksByID(ctx context.Context, booksid string) (res booksdto.BooksResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.booksrepository.GetBooksByID(ctx, booksid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Books"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = booksdto.BooksResp{
		ID:          resRepo.ID,
		Title:       resRepo.Title,
		Description: resRepo.Description,
		Author:      resRepo.Author,
	}

	return res, nil
}
func (alc *BooksUseCaseImpl) CreateBooks(ctx context.Context, data booksdto.BooksReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.booksrepository.CreateBooks(ctx, daos.Book{
		Title:       data.Title,
		Description: data.Description,
		Author:      data.Author,
	})
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *BooksUseCaseImpl) UpdateBooksByID(ctx context.Context, booksid string, data booksdto.BooksReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.booksrepository.UpdateBooksByID(ctx, booksid, daos.Book{
		Title:       data.Title,
		Description: data.Description,
		Author:      data.Author,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *BooksUseCaseImpl) DeleteBooksByID(ctx context.Context, booksid string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.booksrepository.DeleteBooksByID(ctx, booksid)
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
