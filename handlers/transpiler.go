package transpiler

import (
	"github.com/gofiber/fiber/v2"
	clickhouse "github.com/saqura-io/hanami/db"
	traverser "github.com/saqura-io/hanami/pkg/transpilation"
)

type transpilationRequest struct {
	PetalScript string `json:"petalscript"`
}

func RunAgainstClickHouse(c *fiber.Ctx) error {
	var req transpilationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	transpiled := traverser.Walk(req.PetalScript)
	ch := clickhouse.Connect()

	rows, err := ch.GetRowsByQuery(transpiled)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rows)
}

func TranspileToSQL(c *fiber.Ctx) error {
	var req transpilationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	result := traverser.Walk(req.PetalScript)

	return c.JSON(fiber.Map{
		"transpiled": result,
	})
}
