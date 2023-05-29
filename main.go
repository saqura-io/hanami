package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saqura-io/hanami/config"
	"github.com/saqura-io/hanami/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app)

	port := config.Get("SERVER_PORT", "3002")
	err := app.Listen(":" + port)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
