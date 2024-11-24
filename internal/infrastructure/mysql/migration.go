package mysql

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&entity.Book{},
		&entity.User{},
	)
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "Failed Database Migrated", err)
	}

	// var count int64
	// if mysqlDB.Migrator().HasTable(&entity.Book{}) {
	// 	mysqlDB.Model(&entity.Book{}).Count(&count)
	// 	if count < 1 {
	// 		mysqlDB.CreateInBatches(booksSeed, len(booksSeed))
	// 	}
	// }

	helper.Logger(helper.LoggerLevelInfo, "Database Migrated", nil)
}
