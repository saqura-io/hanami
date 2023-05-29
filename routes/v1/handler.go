package v1

import (
	"github.com/gofiber/fiber/v2"
	transpiler "github.com/saqura-io/hanami/handlers"
)

func SetupRoutes(r fiber.Router) {
	r.Post("/transpile", transpiler.TranspileToSQL)
	r.Post("/execute", transpiler.RunAgainstClickHouse)
}
