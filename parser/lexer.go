// Code generated from LitebufLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var (
	_ = fmt.Printf
	_ = sync.Once{}
	_ = unicode.IsLetter
)

type LitebufLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LitebufLexerStaticData struct {
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

func litebufLexerInit() {
	staticData := &LitebufLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'namespace'", "'import'", "'struct'", "'union'", "'enum'", "'root'",
		"'srv'", "'chan'", "'{'", "'}'", "':'", "'='", "'->'", "'bool'", "'i8'",
		"'u8'", "'str'",
	}
	staticData.SymbolicNames = []string{
		"", "NAMESPACE", "IMPORT", "STRUCT", "UNION", "ENUM", "ROOT", "SERVICE",
		"CHANNEL", "LBRACE", "RBRACE", "COLON", "EQUALS", "RARROW", "BOOLEAN",
		"INT8", "UINT8", "STRING", "IDENTIFIER", "ORDER", "URL", "WS",
	}
	staticData.RuleNames = []string{
		"NAMESPACE", "IMPORT", "STRUCT", "UNION", "ENUM", "ROOT", "SERVICE",
		"CHANNEL", "LBRACE", "RBRACE", "COLON", "EQUALS", "RARROW", "BOOLEAN",
		"INT8", "UINT8", "STRING", "IDENTIFIER", "ORDER", "URL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 121, 8,
		17, 10, 17, 12, 17, 124, 9, 17, 1, 18, 4, 18, 127, 8, 18, 11, 18, 12, 18,
		128, 1, 19, 1, 19, 5, 19, 133, 8, 19, 10, 19, 12, 19, 136, 9, 19, 1, 19,
		1, 19, 1, 20, 4, 20, 141, 8, 20, 11, 20, 12, 20, 142, 1, 20, 1, 20, 0,
		0, 21, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32,
		32, 149, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 43, 1, 0, 0, 0, 3, 53, 1, 0,
		0, 0, 5, 60, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 9, 73, 1, 0, 0, 0, 11, 78,
		1, 0, 0, 0, 13, 83, 1, 0, 0, 0, 15, 87, 1, 0, 0, 0, 17, 92, 1, 0, 0, 0,
		19, 94, 1, 0, 0, 0, 21, 96, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0, 25, 100, 1,
		0, 0, 0, 27, 103, 1, 0, 0, 0, 29, 108, 1, 0, 0, 0, 31, 111, 1, 0, 0, 0,
		33, 114, 1, 0, 0, 0, 35, 118, 1, 0, 0, 0, 37, 126, 1, 0, 0, 0, 39, 130,
		1, 0, 0, 0, 41, 140, 1, 0, 0, 0, 43, 44, 5, 110, 0, 0, 44, 45, 5, 97, 0,
		0, 45, 46, 5, 109, 0, 0, 46, 47, 5, 101, 0, 0, 47, 48, 5, 115, 0, 0, 48,
		49, 5, 112, 0, 0, 49, 50, 5, 97, 0, 0, 50, 51, 5, 99, 0, 0, 51, 52, 5,
		101, 0, 0, 52, 2, 1, 0, 0, 0, 53, 54, 5, 105, 0, 0, 54, 55, 5, 109, 0,
		0, 55, 56, 5, 112, 0, 0, 56, 57, 5, 111, 0, 0, 57, 58, 5, 114, 0, 0, 58,
		59, 5, 116, 0, 0, 59, 4, 1, 0, 0, 0, 60, 61, 5, 115, 0, 0, 61, 62, 5, 116,
		0, 0, 62, 63, 5, 114, 0, 0, 63, 64, 5, 117, 0, 0, 64, 65, 5, 99, 0, 0,
		65, 66, 5, 116, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68, 5, 117, 0, 0, 68, 69,
		5, 110, 0, 0, 69, 70, 5, 105, 0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 110,
		0, 0, 72, 8, 1, 0, 0, 0, 73, 74, 5, 101, 0, 0, 74, 75, 5, 110, 0, 0, 75,
		76, 5, 117, 0, 0, 76, 77, 5, 109, 0, 0, 77, 10, 1, 0, 0, 0, 78, 79, 5,
		114, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 116,
		0, 0, 82, 12, 1, 0, 0, 0, 83, 84, 5, 115, 0, 0, 84, 85, 5, 114, 0, 0, 85,
		86, 5, 118, 0, 0, 86, 14, 1, 0, 0, 0, 87, 88, 5, 99, 0, 0, 88, 89, 5, 104,
		0, 0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 110, 0, 0, 91, 16, 1, 0, 0, 0, 92,
		93, 5, 123, 0, 0, 93, 18, 1, 0, 0, 0, 94, 95, 5, 125, 0, 0, 95, 20, 1,
		0, 0, 0, 96, 97, 5, 58, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 5, 61, 0, 0,
		99, 24, 1, 0, 0, 0, 100, 101, 5, 45, 0, 0, 101, 102, 5, 62, 0, 0, 102,
		26, 1, 0, 0, 0, 103, 104, 5, 98, 0, 0, 104, 105, 5, 111, 0, 0, 105, 106,
		5, 111, 0, 0, 106, 107, 5, 108, 0, 0, 107, 28, 1, 0, 0, 0, 108, 109, 5,
		105, 0, 0, 109, 110, 5, 56, 0, 0, 110, 30, 1, 0, 0, 0, 111, 112, 5, 117,
		0, 0, 112, 113, 5, 56, 0, 0, 113, 32, 1, 0, 0, 0, 114, 115, 5, 115, 0,
		0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 114, 0, 0, 117, 34, 1, 0, 0, 0,
		118, 122, 7, 0, 0, 0, 119, 121, 7, 1, 0, 0, 120, 119, 1, 0, 0, 0, 121,
		124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 36, 1,
		0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 127, 7, 2, 0, 0, 126, 125, 1, 0, 0,
		0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		38, 1, 0, 0, 0, 130, 134, 5, 34, 0, 0, 131, 133, 8, 3, 0, 0, 132, 131,
		1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0,
		0, 0, 135, 137, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 138, 5, 34, 0, 0,
		138, 40, 1, 0, 0, 0, 139, 141, 7, 4, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0,
		0, 0, 144, 145, 6, 20, 0, 0, 145, 42, 1, 0, 0, 0, 5, 0, 122, 128, 134,
		142, 1, 6, 0, 0,
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

// LitebufLexerInit initializes any static state used to implement LitebufLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLitebufLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LitebufLexerInit() {
	staticData := &LitebufLexerStaticData
	staticData.once.Do(litebufLexerInit)
}

// NewLitebufLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLitebufLexer(input antlr.CharStream) *LitebufLexer {
	LitebufLexerInit()
	l := new(LitebufLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LitebufLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "LitebufLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LitebufLexer tokens.
const (
	LitebufLexerNAMESPACE  = 1
	LitebufLexerIMPORT     = 2
	LitebufLexerSTRUCT     = 3
	LitebufLexerUNION      = 4
	LitebufLexerENUM       = 5
	LitebufLexerROOT       = 6
	LitebufLexerSERVICE    = 7
	LitebufLexerCHANNEL    = 8
	LitebufLexerLBRACE     = 9
	LitebufLexerRBRACE     = 10
	LitebufLexerCOLON      = 11
	LitebufLexerEQUALS     = 12
	LitebufLexerRARROW     = 13
	LitebufLexerBOOLEAN    = 14
	LitebufLexerINT8       = 15
	LitebufLexerUINT8      = 16
	LitebufLexerSTRING     = 17
	LitebufLexerIDENTIFIER = 18
	LitebufLexerORDER      = 19
	LitebufLexerURL        = 20
	LitebufLexerWS         = 21
)
