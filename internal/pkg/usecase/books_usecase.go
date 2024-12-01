package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	booksmodel "tugas_akhir_example/internal/pkg/model"
	booksrepository "tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BooksUseCase interface {
	GetAllBooks(ctx context.Context, params booksmodel.BooksFilter) (res []booksmodel.BooksResp, err *helper.ErrorStruct)
	GetBooksByID(ctx context.Context, booksid string) (res booksmodel.BooksResp, err *helper.ErrorStruct)
	CreateBooks(ctx context.Context, data booksmodel.BooksReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateBooksByID(ctx context.Context, booksid string, data booksmodel.BooksReqUpdate) (res string, err *helper.ErrorStruct)
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

func (alc *BooksUseCaseImpl) GetAllBooks(ctx context.Context, params booksmodel.BooksFilter) (res []booksmodel.BooksResp, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.booksrepository.GetAllBooks(ctx, entity.FilterBooks{
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
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, booksmodel.BooksResp{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Author:      v.Author,
		})
	}

	return res, nil
}
func (alc *BooksUseCaseImpl) GetBooksByID(ctx context.Context, booksid string) (res booksmodel.BooksResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.booksrepository.GetBooksByID(ctx, booksid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Books"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = booksmodel.BooksResp{
		ID:          resRepo.ID,
		Title:       resRepo.Title,
		Description: resRepo.Description,
		Author:      resRepo.Author,
	}

	return res, nil
}
func (alc *BooksUseCaseImpl) CreateBooks(ctx context.Context, data booksmodel.BooksReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.booksrepository.CreateBooks(ctx, entity.Book{
		Title:       data.Title,
		Description: data.Description,
		Author:      data.Author,
		UserID:      data.UserID,
	})
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *BooksUseCaseImpl) UpdateBooksByID(ctx context.Context, booksid string, data booksmodel.BooksReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.booksrepository.UpdateBooksByID(ctx, booksid, entity.Book{
		Title:       data.Title,
		Description: data.Description,
		Author:      data.Author,
	})

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()), errRepo)
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
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
