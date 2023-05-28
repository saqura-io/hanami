package petalscript // PetalScript
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PetalScript struct {
	*antlr.BaseParser
}

var PetalScriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func petalscriptParserInit() {
	staticData := &PetalScriptParserStaticData
	staticData.LiteralNames = []string{
		"", "'let'", "'from'", "'equals'", "'by'", "'as'", "'in'", "'for'",
		"'filter'", "'search'", "'sort'", "'count'", "'sum'", "'order'", "'output'",
		"'caseInsensitive'", "'exactMatch'", "'='", "'('", "')'", "'['", "']'",
		"';'", "':'", "','", "", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "DECLARATOR", "FROM", "EQUALSEXPR", "BY", "AS", "IN", "FOR", "FILTER",
		"SEARCH", "SORT", "COUNT", "SUM", "ORDER", "OUTPUT", "CASE", "EXACT",
		"EQUALS", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "SEMICOLON", "COLON",
		"COMMA", "BOOL", "TRUE", "FALSE", "ID", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "let_statement", "output_statement", "function_call",
		"filter_call", "search_call", "sort_call", "count_call", "sum_call",
		"aggregate", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 159, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 4, 0, 26, 8, 0, 11, 0, 12, 0, 27, 1, 1, 1, 1, 3,
		1, 32, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 56, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 79, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 89,
		8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 95, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 105, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 124, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 138, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 5, 10, 150, 8, 10, 10, 10, 12, 10, 153, 9, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 0, 0, 160, 0, 25, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 35,
		1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8, 55, 1, 0, 0, 0, 10, 57, 1, 0, 0, 0, 12,
		72, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 117, 1, 0, 0, 0, 18, 131, 1, 0,
		0, 0, 20, 145, 1, 0, 0, 0, 22, 156, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25,
		24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0,
		0, 28, 1, 1, 0, 0, 0, 29, 32, 3, 4, 2, 0, 30, 32, 3, 6, 3, 0, 31, 29, 1,
		0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 34, 5, 22, 0, 0, 34,
		3, 1, 0, 0, 0, 35, 36, 5, 1, 0, 0, 36, 37, 3, 22, 11, 0, 37, 38, 5, 17,
		0, 0, 38, 39, 3, 8, 4, 0, 39, 5, 1, 0, 0, 0, 40, 41, 5, 14, 0, 0, 41, 42,
		5, 18, 0, 0, 42, 43, 5, 5, 0, 0, 43, 44, 5, 23, 0, 0, 44, 45, 5, 29, 0,
		0, 45, 46, 5, 24, 0, 0, 46, 47, 3, 22, 11, 0, 47, 48, 5, 19, 0, 0, 48,
		7, 1, 0, 0, 0, 49, 56, 3, 10, 5, 0, 50, 56, 3, 12, 6, 0, 51, 56, 3, 14,
		7, 0, 52, 56, 3, 16, 8, 0, 53, 56, 3, 18, 9, 0, 54, 56, 3, 20, 10, 0, 55,
		49, 1, 0, 0, 0, 55, 50, 1, 0, 0, 0, 55, 51, 1, 0, 0, 0, 55, 52, 1, 0, 0,
		0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 9, 1, 0, 0, 0, 57, 58, 5,
		8, 0, 0, 58, 59, 5, 18, 0, 0, 59, 60, 5, 2, 0, 0, 60, 61, 5, 23, 0, 0,
		61, 62, 5, 29, 0, 0, 62, 63, 5, 24, 0, 0, 63, 64, 5, 4, 0, 0, 64, 65, 5,
		23, 0, 0, 65, 66, 5, 29, 0, 0, 66, 67, 5, 24, 0, 0, 67, 68, 5, 3, 0, 0,
		68, 69, 5, 23, 0, 0, 69, 70, 5, 29, 0, 0, 70, 71, 5, 19, 0, 0, 71, 11,
		1, 0, 0, 0, 72, 73, 5, 9, 0, 0, 73, 74, 5, 18, 0, 0, 74, 75, 5, 6, 0, 0,
		75, 78, 5, 23, 0, 0, 76, 79, 5, 29, 0, 0, 77, 79, 3, 22, 11, 0, 78, 76,
		1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 5, 24, 0, 0,
		81, 82, 5, 7, 0, 0, 82, 83, 5, 23, 0, 0, 83, 88, 5, 29, 0, 0, 84, 85, 5,
		24, 0, 0, 85, 86, 5, 15, 0, 0, 86, 87, 5, 23, 0, 0, 87, 89, 5, 25, 0, 0,
		88, 84, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 94, 1, 0, 0, 0, 90, 91, 5,
		24, 0, 0, 91, 92, 5, 16, 0, 0, 92, 93, 5, 23, 0, 0, 93, 95, 5, 25, 0, 0,
		94, 90, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 5,
		19, 0, 0, 97, 13, 1, 0, 0, 0, 98, 99, 5, 10, 0, 0, 99, 100, 5, 18, 0, 0,
		100, 101, 5, 2, 0, 0, 101, 104, 5, 23, 0, 0, 102, 105, 5, 29, 0, 0, 103,
		105, 3, 22, 11, 0, 104, 102, 1, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106,
		1, 0, 0, 0, 106, 107, 5, 24, 0, 0, 107, 108, 5, 4, 0, 0, 108, 109, 5, 23,
		0, 0, 109, 110, 5, 29, 0, 0, 110, 111, 5, 24, 0, 0, 111, 112, 5, 13, 0,
		0, 112, 113, 5, 23, 0, 0, 113, 114, 5, 29, 0, 0, 114, 115, 6, 7, -1, 0,
		115, 116, 5, 19, 0, 0, 116, 15, 1, 0, 0, 0, 117, 118, 5, 11, 0, 0, 118,
		119, 5, 18, 0, 0, 119, 120, 5, 2, 0, 0, 120, 123, 5, 23, 0, 0, 121, 124,
		5, 29, 0, 0, 122, 124, 3, 22, 11, 0, 123, 121, 1, 0, 0, 0, 123, 122, 1,
		0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 5, 24, 0, 0, 126, 127, 5, 4, 0,
		0, 127, 128, 5, 23, 0, 0, 128, 129, 5, 29, 0, 0, 129, 130, 5, 19, 0, 0,
		130, 17, 1, 0, 0, 0, 131, 132, 5, 12, 0, 0, 132, 133, 5, 18, 0, 0, 133,
		134, 5, 2, 0, 0, 134, 137, 5, 23, 0, 0, 135, 138, 5, 29, 0, 0, 136, 138,
		3, 22, 11, 0, 137, 135, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 139, 1,
		0, 0, 0, 139, 140, 5, 24, 0, 0, 140, 141, 5, 4, 0, 0, 141, 142, 5, 23,
		0, 0, 142, 143, 5, 29, 0, 0, 143, 144, 5, 19, 0, 0, 144, 19, 1, 0, 0, 0,
		145, 146, 5, 20, 0, 0, 146, 151, 3, 22, 11, 0, 147, 148, 5, 24, 0, 0, 148,
		150, 3, 22, 11, 0, 149, 147, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 151, 1, 0,
		0, 0, 154, 155, 5, 21, 0, 0, 155, 21, 1, 0, 0, 0, 156, 157, 5, 28, 0, 0,
		157, 23, 1, 0, 0, 0, 10, 27, 31, 55, 78, 88, 94, 104, 123, 137, 151,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PetalScriptInit initializes any static state used to implement PetalScript. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPetalScript(). You can call this function if you wish to initialize the static state ahead
// of time.
func PetalScriptInit() {
	staticData := &PetalScriptParserStaticData
	staticData.once.Do(petalscriptParserInit)
}

// NewPetalScript produces a new parser instance for the optional input antlr.TokenStream.
func NewPetalScript(input antlr.TokenStream) *PetalScript {
	PetalScriptInit()
	this := new(PetalScript)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PetalScriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PetalScript.g4"

	return this
}

// PetalScript tokens.
const (
	PetalScriptEOF        = antlr.TokenEOF
	PetalScriptDECLARATOR = 1
	PetalScriptFROM       = 2
	PetalScriptEQUALSEXPR = 3
	PetalScriptBY         = 4
	PetalScriptAS         = 5
	PetalScriptIN         = 6
	PetalScriptFOR        = 7
	PetalScriptFILTER     = 8
	PetalScriptSEARCH     = 9
	PetalScriptSORT       = 10
	PetalScriptCOUNT      = 11
	PetalScriptSUM        = 12
	PetalScriptORDER      = 13
	PetalScriptOUTPUT     = 14
	PetalScriptCASE       = 15
	PetalScriptEXACT      = 16
	PetalScriptEQUALS     = 17
	PetalScriptLPAREN     = 18
	PetalScriptRPAREN     = 19
	PetalScriptLBRACKET   = 20
	PetalScriptRBRACKET   = 21
	PetalScriptSEMICOLON  = 22
	PetalScriptCOLON      = 23
	PetalScriptCOMMA      = 24
	PetalScriptBOOL       = 25
	PetalScriptTRUE       = 26
	PetalScriptFALSE      = 27
	PetalScriptID         = 28
	PetalScriptSTRING     = 29
	PetalScriptWS         = 30
	PetalScriptCOMMENT    = 31
)

// PetalScript rules.
const (
	PetalScriptRULE_program          = 0
	PetalScriptRULE_statement        = 1
	PetalScriptRULE_let_statement    = 2
	PetalScriptRULE_output_statement = 3
	PetalScriptRULE_function_call    = 4
	PetalScriptRULE_filter_call      = 5
	PetalScriptRULE_search_call      = 6
	PetalScriptRULE_sort_call        = 7
	PetalScriptRULE_count_call       = 8
	PetalScriptRULE_sum_call         = 9
	PetalScriptRULE_aggregate        = 10
	PetalScriptRULE_identifier       = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *PetalScript) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PetalScriptRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PetalScriptDECLARATOR || _la == PetalScriptOUTPUT {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	Let_statement() ILet_statementContext
	Output_statement() IOutput_statementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PetalScriptSEMICOLON, 0)
}

func (s *StatementContext) Let_statement() ILet_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILet_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILet_statementContext)
}

func (s *StatementContext) Output_statement() IOutput_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutput_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutput_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *PetalScript) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PetalScriptRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PetalScriptDECLARATOR:
		{
			p.SetState(29)
			p.Let_statement()
		}

	case PetalScriptOUTPUT:
		{
			p.SetState(30)
			p.Output_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(33)
		p.Match(PetalScriptSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILet_statementContext is an interface to support dynamic dispatch.
type ILet_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECLARATOR() antlr.TerminalNode
	Identifier() IIdentifierContext
	EQUALS() antlr.TerminalNode
	Function_call() IFunction_callContext

	// IsLet_statementContext differentiates from other interfaces.
	IsLet_statementContext()
}

type Let_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_statementContext() *Let_statementContext {
	var p = new(Let_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_let_statement
	return p
}

func InitEmptyLet_statementContext(p *Let_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_let_statement
}

func (*Let_statementContext) IsLet_statementContext() {}

func NewLet_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_statementContext {
	var p = new(Let_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_let_statement

	return p
}

func (s *Let_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_statementContext) DECLARATOR() antlr.TerminalNode {
	return s.GetToken(PetalScriptDECLARATOR, 0)
}

func (s *Let_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Let_statementContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PetalScriptEQUALS, 0)
}

func (s *Let_statementContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Let_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterLet_statement(s)
	}
}

func (s *Let_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitLet_statement(s)
	}
}

func (p *PetalScript) Let_statement() (localctx ILet_statementContext) {
	localctx = NewLet_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PetalScriptRULE_let_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(PetalScriptDECLARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Identifier()
	}
	{
		p.SetState(37)
		p.Match(PetalScriptEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Function_call()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutput_statementContext is an interface to support dynamic dispatch.
type IOutput_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUTPUT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AS() antlr.TerminalNode
	COLON() antlr.TerminalNode
	STRING() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Identifier() IIdentifierContext
	RPAREN() antlr.TerminalNode

	// IsOutput_statementContext differentiates from other interfaces.
	IsOutput_statementContext()
}

type Output_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutput_statementContext() *Output_statementContext {
	var p = new(Output_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_output_statement
	return p
}

func InitEmptyOutput_statementContext(p *Output_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_output_statement
}

func (*Output_statementContext) IsOutput_statementContext() {}

func NewOutput_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Output_statementContext {
	var p = new(Output_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_output_statement

	return p
}

func (s *Output_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Output_statementContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(PetalScriptOUTPUT, 0)
}

func (s *Output_statementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptLPAREN, 0)
}

func (s *Output_statementContext) AS() antlr.TerminalNode {
	return s.GetToken(PetalScriptAS, 0)
}

func (s *Output_statementContext) COLON() antlr.TerminalNode {
	return s.GetToken(PetalScriptCOLON, 0)
}

func (s *Output_statementContext) STRING() antlr.TerminalNode {
	return s.GetToken(PetalScriptSTRING, 0)
}

func (s *Output_statementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, 0)
}

func (s *Output_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Output_statementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptRPAREN, 0)
}

func (s *Output_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Output_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Output_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterOutput_statement(s)
	}
}

func (s *Output_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitOutput_statement(s)
	}
}

func (p *PetalScript) Output_statement() (localctx IOutput_statementContext) {
	localctx = NewOutput_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PetalScriptRULE_output_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(PetalScriptOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(PetalScriptLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(PetalScriptAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Identifier()
	}
	{
		p.SetState(47)
		p.Match(PetalScriptRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Filter_call() IFilter_callContext
	Search_call() ISearch_callContext
	Sort_call() ISort_callContext
	Count_call() ICount_callContext
	Sum_call() ISum_callContext
	Aggregate() IAggregateContext

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Filter_call() IFilter_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilter_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilter_callContext)
}

func (s *Function_callContext) Search_call() ISearch_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearch_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearch_callContext)
}

func (s *Function_callContext) Sort_call() ISort_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISort_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISort_callContext)
}

func (s *Function_callContext) Count_call() ICount_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICount_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICount_callContext)
}

func (s *Function_callContext) Sum_call() ISum_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISum_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISum_callContext)
}

func (s *Function_callContext) Aggregate() IAggregateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *PetalScript) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PetalScriptRULE_function_call)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PetalScriptFILTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Filter_call()
		}

	case PetalScriptSEARCH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Search_call()
		}

	case PetalScriptSORT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Sort_call()
		}

	case PetalScriptCOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)
			p.Count_call()
		}

	case PetalScriptSUM:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(53)
			p.Sum_call()
		}

	case PetalScriptLBRACKET:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(54)
			p.Aggregate()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilter_callContext is an interface to support dynamic dispatch.
type IFilter_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILTER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FROM() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	BY() antlr.TerminalNode
	EQUALSEXPR() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsFilter_callContext differentiates from other interfaces.
	IsFilter_callContext()
}

type Filter_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_callContext() *Filter_callContext {
	var p = new(Filter_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_filter_call
	return p
}

func InitEmptyFilter_callContext(p *Filter_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_filter_call
}

func (*Filter_callContext) IsFilter_callContext() {}

func NewFilter_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_callContext {
	var p = new(Filter_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_filter_call

	return p
}

func (s *Filter_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_callContext) FILTER() antlr.TerminalNode {
	return s.GetToken(PetalScriptFILTER, 0)
}

func (s *Filter_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptLPAREN, 0)
}

func (s *Filter_callContext) FROM() antlr.TerminalNode {
	return s.GetToken(PetalScriptFROM, 0)
}

func (s *Filter_callContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOLON)
}

func (s *Filter_callContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOLON, i)
}

func (s *Filter_callContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptSTRING)
}

func (s *Filter_callContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptSTRING, i)
}

func (s *Filter_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOMMA)
}

func (s *Filter_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, i)
}

func (s *Filter_callContext) BY() antlr.TerminalNode {
	return s.GetToken(PetalScriptBY, 0)
}

func (s *Filter_callContext) EQUALSEXPR() antlr.TerminalNode {
	return s.GetToken(PetalScriptEQUALSEXPR, 0)
}

func (s *Filter_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptRPAREN, 0)
}

func (s *Filter_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterFilter_call(s)
	}
}

func (s *Filter_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitFilter_call(s)
	}
}

func (p *PetalScript) Filter_call() (localctx IFilter_callContext) {
	localctx = NewFilter_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PetalScriptRULE_filter_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(PetalScriptFILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(PetalScriptLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(PetalScriptFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(PetalScriptBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(PetalScriptEQUALSEXPR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(PetalScriptRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearch_callContext is an interface to support dynamic dispatch.
type ISearch_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEARCH() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	IN() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	FOR() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Identifier() IIdentifierContext
	CASE() antlr.TerminalNode
	AllBOOL() []antlr.TerminalNode
	BOOL(i int) antlr.TerminalNode
	EXACT() antlr.TerminalNode

	// IsSearch_callContext differentiates from other interfaces.
	IsSearch_callContext()
}

type Search_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearch_callContext() *Search_callContext {
	var p = new(Search_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_search_call
	return p
}

func InitEmptySearch_callContext(p *Search_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_search_call
}

func (*Search_callContext) IsSearch_callContext() {}

func NewSearch_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Search_callContext {
	var p = new(Search_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_search_call

	return p
}

func (s *Search_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Search_callContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(PetalScriptSEARCH, 0)
}

func (s *Search_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptLPAREN, 0)
}

func (s *Search_callContext) IN() antlr.TerminalNode {
	return s.GetToken(PetalScriptIN, 0)
}

func (s *Search_callContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOLON)
}

func (s *Search_callContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOLON, i)
}

func (s *Search_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOMMA)
}

func (s *Search_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, i)
}

func (s *Search_callContext) FOR() antlr.TerminalNode {
	return s.GetToken(PetalScriptFOR, 0)
}

func (s *Search_callContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptSTRING)
}

func (s *Search_callContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptSTRING, i)
}

func (s *Search_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptRPAREN, 0)
}

func (s *Search_callContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Search_callContext) CASE() antlr.TerminalNode {
	return s.GetToken(PetalScriptCASE, 0)
}

func (s *Search_callContext) AllBOOL() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptBOOL)
}

func (s *Search_callContext) BOOL(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptBOOL, i)
}

func (s *Search_callContext) EXACT() antlr.TerminalNode {
	return s.GetToken(PetalScriptEXACT, 0)
}

func (s *Search_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Search_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Search_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterSearch_call(s)
	}
}

func (s *Search_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitSearch_call(s)
	}
}

func (p *PetalScript) Search_call() (localctx ISearch_callContext) {
	localctx = NewSearch_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PetalScriptRULE_search_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(PetalScriptSEARCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(PetalScriptLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(PetalScriptIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PetalScriptSTRING:
		{
			p.SetState(76)
			p.Match(PetalScriptSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PetalScriptID:
		{
			p.SetState(77)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(80)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(PetalScriptFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(84)
			p.Match(PetalScriptCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(PetalScriptCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(PetalScriptCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(PetalScriptBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PetalScriptCOMMA {
		{
			p.SetState(90)
			p.Match(PetalScriptCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(PetalScriptEXACT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(PetalScriptCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(PetalScriptBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(96)
		p.Match(PetalScriptRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISort_callContext is an interface to support dynamic dispatch.
type ISort_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Getter signatures
	SORT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FROM() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	BY() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	ORDER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsSort_callContext differentiates from other interfaces.
	IsSort_callContext()
}

type Sort_callContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	_STRING antlr.Token
}

func NewEmptySort_callContext() *Sort_callContext {
	var p = new(Sort_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_sort_call
	return p
}

func InitEmptySort_callContext(p *Sort_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_sort_call
}

func (*Sort_callContext) IsSort_callContext() {}

func NewSort_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sort_callContext {
	var p = new(Sort_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_sort_call

	return p
}

func (s *Sort_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Sort_callContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Sort_callContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Sort_callContext) SORT() antlr.TerminalNode {
	return s.GetToken(PetalScriptSORT, 0)
}

func (s *Sort_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptLPAREN, 0)
}

func (s *Sort_callContext) FROM() antlr.TerminalNode {
	return s.GetToken(PetalScriptFROM, 0)
}

func (s *Sort_callContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOLON)
}

func (s *Sort_callContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOLON, i)
}

func (s *Sort_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOMMA)
}

func (s *Sort_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, i)
}

func (s *Sort_callContext) BY() antlr.TerminalNode {
	return s.GetToken(PetalScriptBY, 0)
}

func (s *Sort_callContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptSTRING)
}

func (s *Sort_callContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptSTRING, i)
}

func (s *Sort_callContext) ORDER() antlr.TerminalNode {
	return s.GetToken(PetalScriptORDER, 0)
}

func (s *Sort_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptRPAREN, 0)
}

func (s *Sort_callContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Sort_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sort_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sort_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterSort_call(s)
	}
}

func (s *Sort_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitSort_call(s)
	}
}

func (p *PetalScript) Sort_call() (localctx ISort_callContext) {
	localctx = NewSort_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PetalScriptRULE_sort_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(PetalScriptSORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(PetalScriptLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(PetalScriptFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PetalScriptSTRING:
		{
			p.SetState(102)

			var _m = p.Match(PetalScriptSTRING)

			localctx.(*Sort_callContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PetalScriptID:
		{
			p.SetState(103)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(106)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(PetalScriptBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)

		var _m = p.Match(PetalScriptSTRING)

		localctx.(*Sort_callContext)._STRING = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(PetalScriptORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)

		var _m = p.Match(PetalScriptSTRING)

		localctx.(*Sort_callContext)._STRING = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(PetalScriptRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICount_callContext is an interface to support dynamic dispatch.
type ICount_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FROM() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsCount_callContext differentiates from other interfaces.
	IsCount_callContext()
}

type Count_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCount_callContext() *Count_callContext {
	var p = new(Count_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_count_call
	return p
}

func InitEmptyCount_callContext(p *Count_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_count_call
}

func (*Count_callContext) IsCount_callContext() {}

func NewCount_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Count_callContext {
	var p = new(Count_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_count_call

	return p
}

func (s *Count_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Count_callContext) COUNT() antlr.TerminalNode {
	return s.GetToken(PetalScriptCOUNT, 0)
}

func (s *Count_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptLPAREN, 0)
}

func (s *Count_callContext) FROM() antlr.TerminalNode {
	return s.GetToken(PetalScriptFROM, 0)
}

func (s *Count_callContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOLON)
}

func (s *Count_callContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOLON, i)
}

func (s *Count_callContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, 0)
}

func (s *Count_callContext) BY() antlr.TerminalNode {
	return s.GetToken(PetalScriptBY, 0)
}

func (s *Count_callContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptSTRING)
}

func (s *Count_callContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptSTRING, i)
}

func (s *Count_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptRPAREN, 0)
}

func (s *Count_callContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Count_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Count_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Count_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterCount_call(s)
	}
}

func (s *Count_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitCount_call(s)
	}
}

func (p *PetalScript) Count_call() (localctx ICount_callContext) {
	localctx = NewCount_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PetalScriptRULE_count_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(PetalScriptCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(PetalScriptLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(PetalScriptFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PetalScriptSTRING:
		{
			p.SetState(121)
			p.Match(PetalScriptSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PetalScriptID:
		{
			p.SetState(122)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(125)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(PetalScriptBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(PetalScriptRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISum_callContext is an interface to support dynamic dispatch.
type ISum_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUM() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FROM() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsSum_callContext differentiates from other interfaces.
	IsSum_callContext()
}

type Sum_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySum_callContext() *Sum_callContext {
	var p = new(Sum_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_sum_call
	return p
}

func InitEmptySum_callContext(p *Sum_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_sum_call
}

func (*Sum_callContext) IsSum_callContext() {}

func NewSum_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sum_callContext {
	var p = new(Sum_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_sum_call

	return p
}

func (s *Sum_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Sum_callContext) SUM() antlr.TerminalNode {
	return s.GetToken(PetalScriptSUM, 0)
}

func (s *Sum_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptLPAREN, 0)
}

func (s *Sum_callContext) FROM() antlr.TerminalNode {
	return s.GetToken(PetalScriptFROM, 0)
}

func (s *Sum_callContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOLON)
}

func (s *Sum_callContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOLON, i)
}

func (s *Sum_callContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, 0)
}

func (s *Sum_callContext) BY() antlr.TerminalNode {
	return s.GetToken(PetalScriptBY, 0)
}

func (s *Sum_callContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptSTRING)
}

func (s *Sum_callContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptSTRING, i)
}

func (s *Sum_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PetalScriptRPAREN, 0)
}

func (s *Sum_callContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Sum_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sum_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sum_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterSum_call(s)
	}
}

func (s *Sum_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitSum_call(s)
	}
}

func (p *PetalScript) Sum_call() (localctx ISum_callContext) {
	localctx = NewSum_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PetalScriptRULE_sum_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(PetalScriptSUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(PetalScriptLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(PetalScriptFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PetalScriptSTRING:
		{
			p.SetState(135)
			p.Match(PetalScriptSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PetalScriptID:
		{
			p.SetState(136)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(139)
		p.Match(PetalScriptCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(PetalScriptBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(PetalScriptCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(PetalScriptSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(PetalScriptRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateContext is an interface to support dynamic dispatch.
type IAggregateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	RBRACKET() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAggregateContext differentiates from other interfaces.
	IsAggregateContext()
}

type AggregateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateContext() *AggregateContext {
	var p = new(AggregateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_aggregate
	return p
}

func InitEmptyAggregateContext(p *AggregateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_aggregate
}

func (*AggregateContext) IsAggregateContext() {}

func NewAggregateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateContext {
	var p = new(AggregateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_aggregate

	return p
}

func (s *AggregateContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(PetalScriptLBRACKET, 0)
}

func (s *AggregateContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *AggregateContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AggregateContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(PetalScriptRBRACKET, 0)
}

func (s *AggregateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PetalScriptCOMMA)
}

func (s *AggregateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PetalScriptCOMMA, i)
}

func (s *AggregateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterAggregate(s)
	}
}

func (s *AggregateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitAggregate(s)
	}
}

func (p *PetalScript) Aggregate() (localctx IAggregateContext) {
	localctx = NewAggregateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PetalScriptRULE_aggregate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(PetalScriptLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Identifier()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PetalScriptCOMMA {
		{
			p.SetState(147)
			p.Match(PetalScriptCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Identifier()
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(154)
		p.Match(PetalScriptRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PetalScriptRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PetalScriptRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(PetalScriptID, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PetalScriptListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *PetalScript) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PetalScriptRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(PetalScriptID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
