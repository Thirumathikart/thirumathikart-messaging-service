package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	"github.com/thirumathikart/thirumathikart-messaging-service/routes"
	"github.com/thirumathikart/thirumathikart-messaging-service/utils"
)

func main() {
	config.InitConfig()
	server := echo.New()
	utils.InitLogger(server)
	server.Use(middleware.CORS())

	routes.Init(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
