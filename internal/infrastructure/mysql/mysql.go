package mysql

import (
	"fmt"
	"tugas_akhir_example/internal/helper"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct {
	Username           string `mapstructure:"mysql_username"`
	Password           string `mapstructure:"mysql_password"`
	DbName             string `mapstructure:"mysql_Dbname"`
	Host               string `mapstructure:"mysql_host"`
	Port               int    `mapstructure:"mysql_port"`
	Schema             string `mapstructure:"mysql_schema"`
	LogMode            bool   `mapstructure:"mysql_logMode"`
	MaxLifetime        int    `mapstructure:"mysql_maxLifetime"`
	MinIdleConnections int    `mapstructure:"mysql_minIdleConnections"`
	MaxOpenConnections int    `mapstructure:"mysql_maxOpenConnections"`
}

func DatabaseInit(v *viper.Viper) *gorm.DB {
	var mysqlConfig MysqlConf
	err := v.Unmarshal(&mysqlConfig)
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("failed init database mysql : %s", err.Error()), err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("Cannot conenct to database : %s", err.Error()), err)
		panic(err)
	}

	_, err = db.DB()
	if err != nil {
		// helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("Cannot conenct to database : %s", err.Error()), err)
		panic(err)
	}

	// TODO POOLING CONNECTION

	helper.Logger(helper.LoggerLevelInfo, "⇨ MySQL status is connected", nil)
	RunMigration(db)

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("Failed to close connection to database : %s", err.Error()), err)
	}

	dbSQL.Close()

}
