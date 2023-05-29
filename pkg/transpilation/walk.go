package traverser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/saqura-io/hanami/listeners"
	"github.com/saqura-io/hanami/petalscript"
)

func Walk(input string) string {
	is := antlr.NewInputStream(input)
	lexer := petalscript.NewPetalScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := petalscript.NewPetalScript(stream)
	l := listeners.NewPetalListener()

	antlr.ParseTreeWalkerDefault.Walk(l, p.Program())

	return l.Result
}
