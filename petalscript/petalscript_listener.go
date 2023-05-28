package petalscript // PetalScript
import "github.com/antlr4-go/antlr/v4"

// PetalScriptListener is a complete listener for a parse tree produced by PetalScript.
type PetalScriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLet_statement is called when entering the let_statement production.
	EnterLet_statement(c *Let_statementContext)

	// EnterOutput_statement is called when entering the output_statement production.
	EnterOutput_statement(c *Output_statementContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterFilter_call is called when entering the filter_call production.
	EnterFilter_call(c *Filter_callContext)

	// EnterSearch_call is called when entering the search_call production.
	EnterSearch_call(c *Search_callContext)

	// EnterSort_call is called when entering the sort_call production.
	EnterSort_call(c *Sort_callContext)

	// EnterCount_call is called when entering the count_call production.
	EnterCount_call(c *Count_callContext)

	// EnterSum_call is called when entering the sum_call production.
	EnterSum_call(c *Sum_callContext)

	// EnterAggregate is called when entering the aggregate production.
	EnterAggregate(c *AggregateContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLet_statement is called when exiting the let_statement production.
	ExitLet_statement(c *Let_statementContext)

	// ExitOutput_statement is called when exiting the output_statement production.
	ExitOutput_statement(c *Output_statementContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitFilter_call is called when exiting the filter_call production.
	ExitFilter_call(c *Filter_callContext)

	// ExitSearch_call is called when exiting the search_call production.
	ExitSearch_call(c *Search_callContext)

	// ExitSort_call is called when exiting the sort_call production.
	ExitSort_call(c *Sort_callContext)

	// ExitCount_call is called when exiting the count_call production.
	ExitCount_call(c *Count_callContext)

	// ExitSum_call is called when exiting the sum_call production.
	ExitSum_call(c *Sum_callContext)

	// ExitAggregate is called when exiting the aggregate production.
	ExitAggregate(c *AggregateContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
