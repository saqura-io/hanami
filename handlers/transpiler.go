package transpiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/saqura-io/hanami/listeners"
	"github.com/saqura-io/hanami/petalscript"
)

type transpilationRequest struct {
	PetalScript string `json:"petalscript"`
}

func TranspileToSQL(c *fiber.Ctx) error {
	var req transpilationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	is := antlr.NewInputStream(req.PetalScript)
	lexer := petalscript.NewPetalScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := petalscript.NewPetalScript(stream)
	l := listeners.NewPetalListener()

	antlr.ParseTreeWalkerDefault.Walk(l, p.Program())

	return c.JSON(fiber.Map{
		"transpiled": l.Result,
	})
}
