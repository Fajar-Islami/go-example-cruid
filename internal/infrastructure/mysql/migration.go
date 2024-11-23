package mysql

import (
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&entity.Book{},
	)

	var count int64
	if mysqlDB.Migrator().HasTable(&entity.Book{}) {
		mysqlDB.Model(&entity.Book{}).Count(&count)
		if count < 1 {
			mysqlDB.CreateInBatches(booksSeed, len(booksSeed))
		}
	}

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
