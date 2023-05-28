package petalscript // PetalScript
import "github.com/antlr4-go/antlr/v4"

// BasePetalScriptListener is a complete listener for a parse tree produced by PetalScript.
type BasePetalScriptListener struct{}

var _ PetalScriptListener = &BasePetalScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePetalScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePetalScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePetalScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePetalScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasePetalScriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasePetalScriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePetalScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePetalScriptListener) ExitStatement(ctx *StatementContext) {}

// EnterLet_statement is called when production let_statement is entered.
func (s *BasePetalScriptListener) EnterLet_statement(ctx *Let_statementContext) {}

// ExitLet_statement is called when production let_statement is exited.
func (s *BasePetalScriptListener) ExitLet_statement(ctx *Let_statementContext) {}

// EnterOutput_statement is called when production output_statement is entered.
func (s *BasePetalScriptListener) EnterOutput_statement(ctx *Output_statementContext) {}

// ExitOutput_statement is called when production output_statement is exited.
func (s *BasePetalScriptListener) ExitOutput_statement(ctx *Output_statementContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BasePetalScriptListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BasePetalScriptListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterFilter_call is called when production filter_call is entered.
func (s *BasePetalScriptListener) EnterFilter_call(ctx *Filter_callContext) {}

// ExitFilter_call is called when production filter_call is exited.
func (s *BasePetalScriptListener) ExitFilter_call(ctx *Filter_callContext) {}

// EnterSearch_call is called when production search_call is entered.
func (s *BasePetalScriptListener) EnterSearch_call(ctx *Search_callContext) {}

// ExitSearch_call is called when production search_call is exited.
func (s *BasePetalScriptListener) ExitSearch_call(ctx *Search_callContext) {}

// EnterSort_call is called when production sort_call is entered.
func (s *BasePetalScriptListener) EnterSort_call(ctx *Sort_callContext) {}

// ExitSort_call is called when production sort_call is exited.
func (s *BasePetalScriptListener) ExitSort_call(ctx *Sort_callContext) {}

// EnterCount_call is called when production count_call is entered.
func (s *BasePetalScriptListener) EnterCount_call(ctx *Count_callContext) {}

// ExitCount_call is called when production count_call is exited.
func (s *BasePetalScriptListener) ExitCount_call(ctx *Count_callContext) {}

// EnterSum_call is called when production sum_call is entered.
func (s *BasePetalScriptListener) EnterSum_call(ctx *Sum_callContext) {}

// ExitSum_call is called when production sum_call is exited.
func (s *BasePetalScriptListener) ExitSum_call(ctx *Sum_callContext) {}

// EnterAggregate is called when production aggregate is entered.
func (s *BasePetalScriptListener) EnterAggregate(ctx *AggregateContext) {}

// ExitAggregate is called when production aggregate is exited.
func (s *BasePetalScriptListener) ExitAggregate(ctx *AggregateContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePetalScriptListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePetalScriptListener) ExitIdentifier(ctx *IdentifierContext) {}
