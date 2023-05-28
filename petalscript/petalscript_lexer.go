package petalscript

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PetalScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PetalScriptLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func petalscriptlexerLexerInit() {
	staticData := &PetalScriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"DECLARATOR", "FROM", "EQUALSEXPR", "BY", "AS", "IN", "FOR", "FILTER",
		"SEARCH", "SORT", "COUNT", "SUM", "ORDER", "OUTPUT", "CASE", "EXACT",
		"EQUALS", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "SEMICOLON", "COLON",
		"COMMA", "BOOL", "TRUE", "FALSE", "ID", "STRING", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 224, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 180, 8, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5,
		27, 195, 8, 27, 10, 27, 12, 27, 198, 9, 27, 1, 28, 1, 28, 5, 28, 202, 8,
		28, 10, 28, 12, 28, 205, 9, 28, 1, 28, 1, 28, 1, 29, 4, 29, 210, 8, 29,
		11, 29, 12, 29, 211, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 218, 8, 30, 10,
		30, 12, 30, 221, 9, 30, 1, 30, 1, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 1, 0, 5,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3,
		0, 10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13,
		13, 228, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 1, 63, 1, 0, 0, 0, 3, 67, 1, 0, 0, 0, 5, 72, 1, 0, 0, 0, 7,
		79, 1, 0, 0, 0, 9, 82, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 88, 1, 0, 0,
		0, 15, 92, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0, 19, 106, 1, 0, 0, 0, 21, 111,
		1, 0, 0, 0, 23, 117, 1, 0, 0, 0, 25, 121, 1, 0, 0, 0, 27, 127, 1, 0, 0,
		0, 29, 134, 1, 0, 0, 0, 31, 150, 1, 0, 0, 0, 33, 161, 1, 0, 0, 0, 35, 163,
		1, 0, 0, 0, 37, 165, 1, 0, 0, 0, 39, 167, 1, 0, 0, 0, 41, 169, 1, 0, 0,
		0, 43, 171, 1, 0, 0, 0, 45, 173, 1, 0, 0, 0, 47, 175, 1, 0, 0, 0, 49, 179,
		1, 0, 0, 0, 51, 181, 1, 0, 0, 0, 53, 186, 1, 0, 0, 0, 55, 192, 1, 0, 0,
		0, 57, 199, 1, 0, 0, 0, 59, 209, 1, 0, 0, 0, 61, 215, 1, 0, 0, 0, 63, 64,
		5, 108, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66, 5, 116, 0, 0, 66, 2, 1, 0,
		0, 0, 67, 68, 5, 102, 0, 0, 68, 69, 5, 114, 0, 0, 69, 70, 5, 111, 0, 0,
		70, 71, 5, 109, 0, 0, 71, 4, 1, 0, 0, 0, 72, 73, 5, 101, 0, 0, 73, 74,
		5, 113, 0, 0, 74, 75, 5, 117, 0, 0, 75, 76, 5, 97, 0, 0, 76, 77, 5, 108,
		0, 0, 77, 78, 5, 115, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 98, 0, 0, 80,
		81, 5, 121, 0, 0, 81, 8, 1, 0, 0, 0, 82, 83, 5, 97, 0, 0, 83, 84, 5, 115,
		0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5, 105, 0, 0, 86, 87, 5, 110, 0, 0, 87,
		12, 1, 0, 0, 0, 88, 89, 5, 102, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91, 5,
		114, 0, 0, 91, 14, 1, 0, 0, 0, 92, 93, 5, 102, 0, 0, 93, 94, 5, 105, 0,
		0, 94, 95, 5, 108, 0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 101, 0, 0, 97,
		98, 5, 114, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100, 5, 115, 0, 0, 100, 101,
		5, 101, 0, 0, 101, 102, 5, 97, 0, 0, 102, 103, 5, 114, 0, 0, 103, 104,
		5, 99, 0, 0, 104, 105, 5, 104, 0, 0, 105, 18, 1, 0, 0, 0, 106, 107, 5,
		115, 0, 0, 107, 108, 5, 111, 0, 0, 108, 109, 5, 114, 0, 0, 109, 110, 5,
		116, 0, 0, 110, 20, 1, 0, 0, 0, 111, 112, 5, 99, 0, 0, 112, 113, 5, 111,
		0, 0, 113, 114, 5, 117, 0, 0, 114, 115, 5, 110, 0, 0, 115, 116, 5, 116,
		0, 0, 116, 22, 1, 0, 0, 0, 117, 118, 5, 115, 0, 0, 118, 119, 5, 117, 0,
		0, 119, 120, 5, 109, 0, 0, 120, 24, 1, 0, 0, 0, 121, 122, 5, 111, 0, 0,
		122, 123, 5, 114, 0, 0, 123, 124, 5, 100, 0, 0, 124, 125, 5, 101, 0, 0,
		125, 126, 5, 114, 0, 0, 126, 26, 1, 0, 0, 0, 127, 128, 5, 111, 0, 0, 128,
		129, 5, 117, 0, 0, 129, 130, 5, 116, 0, 0, 130, 131, 5, 112, 0, 0, 131,
		132, 5, 117, 0, 0, 132, 133, 5, 116, 0, 0, 133, 28, 1, 0, 0, 0, 134, 135,
		5, 99, 0, 0, 135, 136, 5, 97, 0, 0, 136, 137, 5, 115, 0, 0, 137, 138, 5,
		101, 0, 0, 138, 139, 5, 73, 0, 0, 139, 140, 5, 110, 0, 0, 140, 141, 5,
		115, 0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 110, 0, 0, 143, 144, 5,
		115, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5, 116, 0, 0, 146, 147, 5,
		105, 0, 0, 147, 148, 5, 118, 0, 0, 148, 149, 5, 101, 0, 0, 149, 30, 1,
		0, 0, 0, 150, 151, 5, 101, 0, 0, 151, 152, 5, 120, 0, 0, 152, 153, 5, 97,
		0, 0, 153, 154, 5, 99, 0, 0, 154, 155, 5, 116, 0, 0, 155, 156, 5, 77, 0,
		0, 156, 157, 5, 97, 0, 0, 157, 158, 5, 116, 0, 0, 158, 159, 5, 99, 0, 0,
		159, 160, 5, 104, 0, 0, 160, 32, 1, 0, 0, 0, 161, 162, 5, 61, 0, 0, 162,
		34, 1, 0, 0, 0, 163, 164, 5, 40, 0, 0, 164, 36, 1, 0, 0, 0, 165, 166, 5,
		41, 0, 0, 166, 38, 1, 0, 0, 0, 167, 168, 5, 91, 0, 0, 168, 40, 1, 0, 0,
		0, 169, 170, 5, 93, 0, 0, 170, 42, 1, 0, 0, 0, 171, 172, 5, 59, 0, 0, 172,
		44, 1, 0, 0, 0, 173, 174, 5, 58, 0, 0, 174, 46, 1, 0, 0, 0, 175, 176, 5,
		44, 0, 0, 176, 48, 1, 0, 0, 0, 177, 180, 3, 51, 25, 0, 178, 180, 3, 53,
		26, 0, 179, 177, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 50, 1, 0, 0, 0,
		181, 182, 5, 116, 0, 0, 182, 183, 5, 114, 0, 0, 183, 184, 5, 117, 0, 0,
		184, 185, 5, 101, 0, 0, 185, 52, 1, 0, 0, 0, 186, 187, 5, 102, 0, 0, 187,
		188, 5, 97, 0, 0, 188, 189, 5, 108, 0, 0, 189, 190, 5, 115, 0, 0, 190,
		191, 5, 101, 0, 0, 191, 54, 1, 0, 0, 0, 192, 196, 7, 0, 0, 0, 193, 195,
		7, 1, 0, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0,
		0, 0, 196, 197, 1, 0, 0, 0, 197, 56, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0,
		199, 203, 5, 34, 0, 0, 200, 202, 8, 2, 0, 0, 201, 200, 1, 0, 0, 0, 202,
		205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 206,
		1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207, 5, 34, 0, 0, 207, 58, 1, 0,
		0, 0, 208, 210, 7, 3, 0, 0, 209, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213,
		214, 6, 29, 0, 0, 214, 60, 1, 0, 0, 0, 215, 219, 5, 35, 0, 0, 216, 218,
		8, 4, 0, 0, 217, 216, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0,
		0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0,
		222, 223, 6, 30, 0, 0, 223, 62, 1, 0, 0, 0, 6, 0, 179, 196, 203, 211, 219,
		1, 6, 0, 0,
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

// PetalScriptLexerInit initializes any static state used to implement PetalScriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPetalScriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PetalScriptLexerInit() {
	staticData := &PetalScriptLexerLexerStaticData
	staticData.once.Do(petalscriptlexerLexerInit)
}

// NewPetalScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPetalScriptLexer(input antlr.CharStream) *PetalScriptLexer {
	PetalScriptLexerInit()
	l := new(PetalScriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PetalScriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PetalScriptLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PetalScriptLexer tokens.
const (
	PetalScriptLexerDECLARATOR = 1
	PetalScriptLexerFROM       = 2
	PetalScriptLexerEQUALSEXPR = 3
	PetalScriptLexerBY         = 4
	PetalScriptLexerAS         = 5
	PetalScriptLexerIN         = 6
	PetalScriptLexerFOR        = 7
	PetalScriptLexerFILTER     = 8
	PetalScriptLexerSEARCH     = 9
	PetalScriptLexerSORT       = 10
	PetalScriptLexerCOUNT      = 11
	PetalScriptLexerSUM        = 12
	PetalScriptLexerORDER      = 13
	PetalScriptLexerOUTPUT     = 14
	PetalScriptLexerCASE       = 15
	PetalScriptLexerEXACT      = 16
	PetalScriptLexerEQUALS     = 17
	PetalScriptLexerLPAREN     = 18
	PetalScriptLexerRPAREN     = 19
	PetalScriptLexerLBRACKET   = 20
	PetalScriptLexerRBRACKET   = 21
	PetalScriptLexerSEMICOLON  = 22
	PetalScriptLexerCOLON      = 23
	PetalScriptLexerCOMMA      = 24
	PetalScriptLexerBOOL       = 25
	PetalScriptLexerTRUE       = 26
	PetalScriptLexerFALSE      = 27
	PetalScriptLexerID         = 28
	PetalScriptLexerSTRING     = 29
	PetalScriptLexerWS         = 30
	PetalScriptLexerCOMMENT    = 31
)
