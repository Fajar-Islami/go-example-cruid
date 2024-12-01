package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	usermodel "tugas_akhir_example/internal/pkg/model"
	userrepository "tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersUseCase interface {
	Login(ctx context.Context, params usermodel.Login) (res usermodel.LoginRes, err *helper.ErrorStruct)
	CreateUsers(ctx context.Context, data usermodel.CreateUser) (res uint, err *helper.ErrorStruct)
}

type UsersUseCaseImpl struct {
	userrepository userrepository.UsersRepository
}

func NewUsersUseCase(userrepository userrepository.UsersRepository) UsersUseCase {
	return &UsersUseCaseImpl{
		userrepository: userrepository,
	}

}

func (alc *UsersUseCaseImpl) Login(ctx context.Context, params usermodel.Login) (res usermodel.LoginRes, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.userrepository.GetUsersByEmail(ctx, params.Email)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Users"),
		}
	}

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllUsers : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	isValid := utils.CheckPasswordHash(params.Password, resRepo.Password)
	if !isValid {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("invalid account"),
		}
	}

	tokenInit := utils.NewToken(utils.DataClaims{
		ID:    fmt.Sprint(resRepo.ID),
		Email: resRepo.Email,
	})

	token, errToken := tokenInit.Create()
	if errToken != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errToken,
		}
	}

	res = usermodel.LoginRes{
		Email: resRepo.Email,
		Name:  resRepo.Name,
		Token: token,
	}

	return res, nil
}
func (alc *UsersUseCaseImpl) CreateUsers(ctx context.Context, params usermodel.CreateUser) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(params); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// TODO PENGECEKAN EMAIL SUDAH TERPAKAI ATAU BELUM

	hashPass, errHash := utils.HashPassword(params.Password)
	if errHash != nil {
		log.Println(errHash)
		err = &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errHash,
		}
		return
	}

	resRepo, errRepo := alc.userrepository.CreateUsers(ctx, entity.User{
		Email:    params.Email,
		Name:     params.Name,
		Password: hashPass,
	})
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateUsers : %s", errRepo.Error()), errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
