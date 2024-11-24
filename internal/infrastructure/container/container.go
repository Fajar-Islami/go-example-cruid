package container

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/mysql"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var v *viper.Viper

type (
	Container struct {
		Mysqldb  *gorm.DB
		Apps     *Apps
		BooksUsc usecase.BooksUseCase
		UserUsc  usecase.UsersUseCase
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

func loadEnv() {
	projectDirName := "go-example-cruid"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + `/.env`)
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()), err)
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()), err)
	}

	err = v.ReadInConfig()
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("failed init config : %s", err.Error()), err)
	}

	helper.Logger(helper.LoggerLevelInfo, "Succeed read configuration file", err)
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()), err)
	}
	helper.Logger(helper.LoggerLevelInfo, "Succeed when unmarshal configuration file", err)
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	utils.InitJWT(apps.SecretJwt)
	mysqldb := mysql.DatabaseInit(v)

	bookRepo := repository.NewBooksRepository(mysqldb)
	userRepo := repository.NewUsersRepository(mysqldb)

	bookUsc := usecase.NewBooksUseCase(bookRepo)
	userUsc := usecase.NewUsersUseCase(userRepo)

	return &Container{
		Apps:     &apps,
		Mysqldb:  mysqldb,
		BooksUsc: bookUsc,
		UserUsc:  userUsc,
	}

}
