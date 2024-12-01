package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type BooksRepository interface {
	GetAllBooks(ctx context.Context, params entity.FilterBooks) (res []entity.Book, err error)
	GetBooksByID(ctx context.Context, booksid string) (res entity.Book, err error)
	CreateBooks(ctx context.Context, data entity.Book) (res uint, err error)
	UpdateBooksByID(ctx context.Context, booksid string, data entity.Book) (res string, err error)
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
func (r *BooksRepositoryImpl) GetAllBooks(ctx context.Context, params entity.FilterBooks) (res []entity.Book, err error) {
	db := r.db

	filter := map[string][]any{
		"title like ? or description like ? or author like ?": {fmt.Sprint("%" + params.Title), "%ab ", "%ab"},
	}

	// if params.Title != "" {
	// 	db = db.Where("title like ?", "%"+params.Title)
	// }

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	// db = db.Where(map[string]interface{}{"created_at BETWEEN ? AND ?": []string{"2000-01-01 00:00:00", "2000-01-01 00:00:00"}})

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *BooksRepositoryImpl) GetBooksByID(ctx context.Context, booksid string) (res entity.Book, err error) {
	if err := r.db.First(&res, booksid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *BooksRepositoryImpl) CreateBooks(ctx context.Context, data entity.Book) (res uint, err error) {
	result := r.db.Debug().Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (r *BooksRepositoryImpl) UpdateBooksByID(ctx context.Context, booksid string, data entity.Book) (res string, err error) {
	var dataBooks entity.Book
	if err = r.db.Where("id = ? ", booksid).First(&dataBooks).WithContext(ctx).Error; err != nil {
		return "Update books failed", gorm.ErrRecordNotFound
	}

	if err := r.db.Model(dataBooks).Updates(&data).Where("id = ? ", booksid).Error; err != nil {
		return "Update books failed", err
	}

	return res, nil
}

func (r *BooksRepositoryImpl) DeleteBooksByID(ctx context.Context, booksid string) (res string, err error) {
	var dataBooks entity.Book
	if err = r.db.Where("id = ?", booksid).First(&dataBooks).WithContext(ctx).Error; err != nil {
		return "Delete book failed", gorm.ErrRecordNotFound
	}

	if err := r.db.Model(dataBooks).Delete(&dataBooks).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
