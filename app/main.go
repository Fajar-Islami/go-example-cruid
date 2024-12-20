package main

import (
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/container"

	rest "tugas_akhir_example/internal/server/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	containerConf := container.InitContainer()
	// defer mysql.CloseDatabaseConnection(containerConf.Mysqldb)

	app := fiber.New()
	app.Use(logger.New())

	rest.HTTPRouteInit(app, containerConf)

	port := fmt.Sprintf("%s:%d", containerConf.Apps.Host, containerConf.Apps.HttpPort)
	if err := app.Listen(port); err != nil {
		helper.Logger(helper.LoggerLevelFatal, "error", err)
	}
}
