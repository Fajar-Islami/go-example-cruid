package container

import (
	"fmt"
	"os"
	"path/filepath"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/mysql"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var v *viper.Viper

const currentfilepath = "internal/infrastructure/container/container.go"

type (
	Container struct {
		Mysqldb *gorm.DB
		Apps    *Apps
	}

	Apps struct {
		Name      string `mapstructure:"name"`
		Host      string `mapstructure:"host"`
		Version   string `mapstructure:"version"`
		Address   string `mapstructure:"address"`
		HttpPort  int    `mapstructure:"httpport"`
		SecretJwt string `mapstructure:"secretJwt"`
	}
)

func init() {
	v = viper.New()

	v.AutomaticEnv()
	v.SetConfigFile(".env")

	path, err := os.Executable()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()))
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()))
	}

	err = v.ReadInConfig()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed init config : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read configuration file")
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()))
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed when unmarshal configuration file")
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	mysqldb := mysql.DatabaseInit(v)

	return &Container{
		Apps:    &apps,
		Mysqldb: mysqldb,
	}

}
