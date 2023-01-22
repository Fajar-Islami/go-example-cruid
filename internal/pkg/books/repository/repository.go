package booksrepository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type BooksRepository interface {
	GetAllBooks(ctx context.Context, limit, page int) (res []daos.Book, err error)
	GetBooksByID(ctx context.Context, booksid string) (res daos.Book, err error)
	CreateBooks(ctx context.Context, data daos.Book) (res uint, err error)
	UpdateBooksByID(ctx context.Context, booksid string, data daos.Book) (res string, err error)
	DeleteBooksByID(ctx context.Context, booksid string) (res string, err error)
}

type BooksRepositoryImpl struct {
	db *gorm.DB
}

func NewBooksRepository(db *gorm.DB) BooksRepository {
	return &BooksRepositoryImpl{
		db: db,
	}
}
func (alr *BooksRepositoryImpl) GetAllBooks(ctx context.Context, limit, page int) (res []daos.Book, err error) {
	db := alr.db

	if err := db.Find(&res).WithContext(ctx).Limit(limit).Offset(page).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *BooksRepositoryImpl) GetBooksByID(ctx context.Context, booksid string) (res daos.Book, err error) {
	if err := alr.db.First(&res, booksid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *BooksRepositoryImpl) CreateBooks(ctx context.Context, data daos.Book) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (alr *BooksRepositoryImpl) UpdateBooksByID(ctx context.Context, booksid string, data daos.Book) (res string, err error) {
	var dataBooks daos.Book
	if err = alr.db.Where("id = ? ", booksid).First(&dataBooks).WithContext(ctx).Error; err != nil {
		return "Update books failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataBooks).Updates(&data).Where("id = ? ", booksid).Error; err != nil {
		return "Update books failed", err
	}

	return res, nil
}

func (alr *BooksRepositoryImpl) DeleteBooksByID(ctx context.Context, booksid string) (res string, err error) {
	var dataBooks daos.Book
	if err = alr.db.Where("id = ?", booksid).First(&dataBooks).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataBooks).Delete(&dataBooks).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
