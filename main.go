package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saqura-io/hanami/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app)

	port := 3002 // TODO: Add ENV config reader
	err := app.Listen(":" + strconv.Itoa(port))

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
