// Code generated from LitebufParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LitebufParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var (
	_ = fmt.Printf
	_ = strconv.Itoa
	_ = sync.Once{}
)

type LitebufParser struct {
	*antlr.BaseParser
}

var LitebufParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func litebufParserInit() {
	staticData := &LitebufParserStaticData
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
		"source", "namespace", "include", "def", "struct", "union", "field",
		"enum", "variant", "order", "root", "service", "fn", "channel", "type",
		"primitive", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 154, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 3, 0, 37, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		43, 8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 4, 2, 57, 8, 2, 11, 2, 12, 2, 58, 1, 2, 3, 2, 62, 8, 2,
		1, 3, 1, 3, 1, 3, 3, 3, 67, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 73, 8,
		4, 11, 4, 12, 4, 74, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 83, 8, 5,
		11, 5, 12, 5, 84, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 4, 7, 97, 8, 7, 11, 7, 12, 7, 98, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 105,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 113, 8, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 5, 11, 119, 8, 11, 10, 11, 12, 11, 122, 9, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 131, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 137, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 144, 8, 13, 1, 14, 1, 14, 3, 14, 148, 8, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 0, 1, 1, 0, 14, 17, 158, 0, 34, 1, 0, 0, 0, 2, 49,
		1, 0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 66, 1, 0, 0, 0, 8, 68, 1, 0, 0, 0, 10,
		78, 1, 0, 0, 0, 12, 88, 1, 0, 0, 0, 14, 92, 1, 0, 0, 0, 16, 102, 1, 0,
		0, 0, 18, 106, 1, 0, 0, 0, 20, 109, 1, 0, 0, 0, 22, 114, 1, 0, 0, 0, 24,
		125, 1, 0, 0, 0, 26, 138, 1, 0, 0, 0, 28, 147, 1, 0, 0, 0, 30, 149, 1,
		0, 0, 0, 32, 151, 1, 0, 0, 0, 34, 36, 3, 2, 1, 0, 35, 37, 3, 4, 2, 0, 36,
		35, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 44, 1, 0, 0, 0, 38, 43, 3, 6, 3,
		0, 39, 43, 3, 20, 10, 0, 40, 43, 3, 22, 11, 0, 41, 43, 3, 26, 13, 0, 42,
		38, 1, 0, 0, 0, 42, 39, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0,
		0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 0, 0, 1, 48, 1, 1, 0, 0, 0,
		49, 50, 5, 1, 0, 0, 50, 51, 3, 32, 16, 0, 51, 3, 1, 0, 0, 0, 52, 61, 5,
		2, 0, 0, 53, 62, 5, 20, 0, 0, 54, 56, 5, 9, 0, 0, 55, 57, 5, 20, 0, 0,
		56, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 5, 10, 0, 0, 61, 53, 1, 0, 0, 0, 61,
		54, 1, 0, 0, 0, 62, 5, 1, 0, 0, 0, 63, 67, 3, 8, 4, 0, 64, 67, 3, 10, 5,
		0, 65, 67, 3, 14, 7, 0, 66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 65,
		1, 0, 0, 0, 67, 7, 1, 0, 0, 0, 68, 69, 5, 3, 0, 0, 69, 70, 3, 32, 16, 0,
		70, 72, 5, 9, 0, 0, 71, 73, 3, 12, 6, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1,
		0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		77, 5, 10, 0, 0, 77, 9, 1, 0, 0, 0, 78, 79, 5, 4, 0, 0, 79, 80, 3, 32,
		16, 0, 80, 82, 5, 9, 0, 0, 81, 83, 3, 12, 6, 0, 82, 81, 1, 0, 0, 0, 83,
		84, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0,
		0, 86, 87, 5, 10, 0, 0, 87, 11, 1, 0, 0, 0, 88, 89, 3, 32, 16, 0, 89, 90,
		5, 11, 0, 0, 90, 91, 3, 28, 14, 0, 91, 13, 1, 0, 0, 0, 92, 93, 5, 5, 0,
		0, 93, 94, 3, 32, 16, 0, 94, 96, 5, 9, 0, 0, 95, 97, 3, 16, 8, 0, 96, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 100, 1, 0, 0, 0, 100, 101, 5, 10, 0, 0, 101, 15, 1, 0, 0, 0, 102, 104,
		3, 32, 16, 0, 103, 105, 3, 18, 9, 0, 104, 103, 1, 0, 0, 0, 104, 105, 1,
		0, 0, 0, 105, 17, 1, 0, 0, 0, 106, 107, 5, 12, 0, 0, 107, 108, 5, 19, 0,
		0, 108, 19, 1, 0, 0, 0, 109, 112, 5, 6, 0, 0, 110, 113, 3, 32, 16, 0, 111,
		113, 3, 6, 3, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 21, 1,
		0, 0, 0, 114, 115, 5, 7, 0, 0, 115, 116, 3, 32, 16, 0, 116, 120, 5, 9,
		0, 0, 117, 119, 3, 24, 12, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0,
		0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 123, 124, 5, 10, 0, 0, 124, 23, 1, 0, 0, 0, 125, 126,
		3, 32, 16, 0, 126, 130, 5, 11, 0, 0, 127, 131, 3, 32, 16, 0, 128, 131,
		3, 28, 14, 0, 129, 131, 3, 6, 3, 0, 130, 127, 1, 0, 0, 0, 130, 128, 1,
		0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 136, 5, 13, 0,
		0, 133, 137, 3, 32, 16, 0, 134, 137, 3, 28, 14, 0, 135, 137, 3, 6, 3, 0,
		136, 133, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137,
		25, 1, 0, 0, 0, 138, 139, 5, 8, 0, 0, 139, 143, 3, 32, 16, 0, 140, 144,
		3, 32, 16, 0, 141, 144, 3, 28, 14, 0, 142, 144, 3, 6, 3, 0, 143, 140, 1,
		0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 27, 1, 0, 0,
		0, 145, 148, 3, 32, 16, 0, 146, 148, 3, 30, 15, 0, 147, 145, 1, 0, 0, 0,
		147, 146, 1, 0, 0, 0, 148, 29, 1, 0, 0, 0, 149, 150, 7, 0, 0, 0, 150, 31,
		1, 0, 0, 0, 151, 152, 5, 18, 0, 0, 152, 33, 1, 0, 0, 0, 16, 36, 42, 44,
		58, 61, 66, 74, 84, 98, 104, 112, 120, 130, 136, 143, 147,
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

// LitebufParserInit initializes any static state used to implement LitebufParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLitebufParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LitebufParserInit() {
	staticData := &LitebufParserStaticData
	staticData.once.Do(litebufParserInit)
}

// NewLitebufParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLitebufParser(input antlr.TokenStream) *LitebufParser {
	LitebufParserInit()
	this := new(LitebufParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LitebufParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "LitebufParser.g4"

	return this
}

// LitebufParser tokens.
const (
	LitebufParserEOF        = antlr.TokenEOF
	LitebufParserNAMESPACE  = 1
	LitebufParserIMPORT     = 2
	LitebufParserSTRUCT     = 3
	LitebufParserUNION      = 4
	LitebufParserENUM       = 5
	LitebufParserROOT       = 6
	LitebufParserSERVICE    = 7
	LitebufParserCHANNEL    = 8
	LitebufParserLBRACE     = 9
	LitebufParserRBRACE     = 10
	LitebufParserCOLON      = 11
	LitebufParserEQUALS     = 12
	LitebufParserRARROW     = 13
	LitebufParserBOOLEAN    = 14
	LitebufParserINT8       = 15
	LitebufParserUINT8      = 16
	LitebufParserSTRING     = 17
	LitebufParserIDENTIFIER = 18
	LitebufParserORDER      = 19
	LitebufParserURL        = 20
	LitebufParserWS         = 21
)

// LitebufParser rules.
const (
	LitebufParserRULE_source     = 0
	LitebufParserRULE_namespace  = 1
	LitebufParserRULE_include    = 2
	LitebufParserRULE_def        = 3
	LitebufParserRULE_struct     = 4
	LitebufParserRULE_union      = 5
	LitebufParserRULE_field      = 6
	LitebufParserRULE_enum       = 7
	LitebufParserRULE_variant    = 8
	LitebufParserRULE_order      = 9
	LitebufParserRULE_root       = 10
	LitebufParserRULE_service    = 11
	LitebufParserRULE_fn         = 12
	LitebufParserRULE_channel    = 13
	LitebufParserRULE_type       = 14
	LitebufParserRULE_primitive  = 15
	LitebufParserRULE_identifier = 16
)

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Namespace() INamespaceContext
	EOF() antlr.TerminalNode
	Include() IIncludeContext
	AllDef() []IDefContext
	Def(i int) IDefContext
	AllRoot() []IRootContext
	Root(i int) IRootContext
	AllService() []IServiceContext
	Service(i int) IServiceContext
	AllChannel() []IChannelContext
	Channel(i int) IChannelContext

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	p := new(SourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_source
	return p
}

func InitEmptySourceContext(p *SourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_source
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	p := new(SourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) Namespace() INamespaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *SourceContext) EOF() antlr.TerminalNode {
	return s.GetToken(LitebufParserEOF, 0)
}

func (s *SourceContext) Include() IIncludeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *SourceContext) AllDef() []IDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefContext); ok {
			len++
		}
	}

	tst := make([]IDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefContext); ok {
			tst[i] = t.(IDefContext)
			i++
		}
	}

	return tst
}

func (s *SourceContext) Def(i int) IDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
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

	return t.(IDefContext)
}

func (s *SourceContext) AllRoot() []IRootContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRootContext); ok {
			len++
		}
	}

	tst := make([]IRootContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRootContext); ok {
			tst[i] = t.(IRootContext)
			i++
		}
	}

	return tst
}

func (s *SourceContext) Root(i int) IRootContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootContext); ok {
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

	return t.(IRootContext)
}

func (s *SourceContext) AllService() []IServiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceContext); ok {
			len++
		}
	}

	tst := make([]IServiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceContext); ok {
			tst[i] = t.(IServiceContext)
			i++
		}
	}

	return tst
}

func (s *SourceContext) Service(i int) IServiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceContext); ok {
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

	return t.(IServiceContext)
}

func (s *SourceContext) AllChannel() []IChannelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChannelContext); ok {
			len++
		}
	}

	tst := make([]IChannelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChannelContext); ok {
			tst[i] = t.(IChannelContext)
			i++
		}
	}

	return tst
}

func (s *SourceContext) Channel(i int) IChannelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChannelContext); ok {
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

	return t.(IChannelContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitSource(s)
	}
}

func (s *SourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LitebufParserRULE_source)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Namespace()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LitebufParserIMPORT {
		{
			p.SetState(35)
			p.Include()
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&504) != 0 {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LitebufParserSTRUCT, LitebufParserUNION, LitebufParserENUM:
			{
				p.SetState(38)
				p.Def()
			}

		case LitebufParserROOT:
			{
				p.SetState(39)
				p.Root()
			}

		case LitebufParserSERVICE:
			{
				p.SetState(40)
				p.Service()
			}

		case LitebufParserCHANNEL:
			{
				p.SetState(41)
				p.Channel()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(LitebufParserEOF)
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

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMESPACE() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	p := new(NamespaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_namespace
	return p
}

func InitEmptyNamespaceContext(p *NamespaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_namespace
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	p := new(NamespaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserNAMESPACE, 0)
}

func (s *NamespaceContext) Identifier() IIdentifierContext {
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

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (s *NamespaceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitNamespace(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LitebufParserRULE_namespace)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(LitebufParserNAMESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Identifier()
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

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	AllURL() []antlr.TerminalNode
	URL(i int) antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	p := new(IncludeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_include
	return p
}

func InitEmptyIncludeContext(p *IncludeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_include
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	p := new(IncludeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(LitebufParserIMPORT, 0)
}

func (s *IncludeContext) AllURL() []antlr.TerminalNode {
	return s.GetTokens(LitebufParserURL)
}

func (s *IncludeContext) URL(i int) antlr.TerminalNode {
	return s.GetToken(LitebufParserURL, i)
}

func (s *IncludeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserLBRACE, 0)
}

func (s *IncludeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserRBRACE, 0)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (s *IncludeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitInclude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LitebufParserRULE_include)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(LitebufParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LitebufParserURL:
		{
			p.SetState(53)
			p.Match(LitebufParserURL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LitebufParserLBRACE:
		{
			p.SetState(54)
			p.Match(LitebufParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LitebufParserURL {
			{
				p.SetState(55)
				p.Match(LitebufParserURL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(LitebufParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IDefContext is an interface to support dynamic dispatch.
type IDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct_() IStructContext
	Union() IUnionContext
	Enum() IEnumContext

	// IsDefContext differentiates from other interfaces.
	IsDefContext()
}

type DefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefContext() *DefContext {
	p := new(DefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_def
	return p
}

func InitEmptyDefContext(p *DefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_def
}

func (*DefContext) IsDefContext() {}

func NewDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefContext {
	p := new(DefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_def

	return p
}

func (s *DefContext) GetParser() antlr.Parser { return s.parser }

func (s *DefContext) Struct_() IStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructContext)
}

func (s *DefContext) Union() IUnionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionContext)
}

func (s *DefContext) Enum() IEnumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumContext)
}

func (s *DefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterDef(s)
	}
}

func (s *DefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitDef(s)
	}
}

func (s *DefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Def() (localctx IDefContext) {
	localctx = NewDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LitebufParserRULE_def)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LitebufParserSTRUCT:
		{
			p.SetState(63)
			p.Struct_()
		}

	case LitebufParserUNION:
		{
			p.SetState(64)
			p.Union()
		}

	case LitebufParserENUM:
		{
			p.SetState(65)
			p.Enum()
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

// IStructContext is an interface to support dynamic dispatch.
type IStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	Identifier() IIdentifierContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructContext() *StructContext {
	p := new(StructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_struct
	return p
}

func InitEmptyStructContext(p *StructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_struct
}

func (*StructContext) IsStructContext() {}

func NewStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructContext {
	p := new(StructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_struct

	return p
}

func (s *StructContext) GetParser() antlr.Parser { return s.parser }

func (s *StructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(LitebufParserSTRUCT, 0)
}

func (s *StructContext) Identifier() IIdentifierContext {
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

func (s *StructContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserLBRACE, 0)
}

func (s *StructContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserRBRACE, 0)
}

func (s *StructContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Struct_() (localctx IStructContext) {
	localctx = NewStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LitebufParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(LitebufParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Identifier()
	}
	{
		p.SetState(70)
		p.Match(LitebufParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LitebufParserIDENTIFIER {
		{
			p.SetState(71)
			p.Field()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
		p.Match(LitebufParserRBRACE)
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

// IUnionContext is an interface to support dynamic dispatch.
type IUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNION() antlr.TerminalNode
	Identifier() IIdentifierContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsUnionContext differentiates from other interfaces.
	IsUnionContext()
}

type UnionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionContext() *UnionContext {
	p := new(UnionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_union
	return p
}

func InitEmptyUnionContext(p *UnionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_union
}

func (*UnionContext) IsUnionContext() {}

func NewUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionContext {
	p := new(UnionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_union

	return p
}

func (s *UnionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionContext) UNION() antlr.TerminalNode {
	return s.GetToken(LitebufParserUNION, 0)
}

func (s *UnionContext) Identifier() IIdentifierContext {
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

func (s *UnionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserLBRACE, 0)
}

func (s *UnionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserRBRACE, 0)
}

func (s *UnionContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *UnionContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (s *UnionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitUnion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Union() (localctx IUnionContext) {
	localctx = NewUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LitebufParserRULE_union)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(LitebufParserUNION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Identifier()
	}
	{
		p.SetState(80)
		p.Match(LitebufParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LitebufParserIDENTIFIER {
		{
			p.SetState(81)
			p.Field()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(LitebufParserRBRACE)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	p := new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	p := new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Identifier() IIdentifierContext {
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

func (s *FieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(LitebufParserCOLON, 0)
}

func (s *FieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LitebufParserRULE_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Identifier()
	}
	{
		p.SetState(89)
		p.Match(LitebufParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Type_()
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

// IEnumContext is an interface to support dynamic dispatch.
type IEnumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	Identifier() IIdentifierContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllVariant() []IVariantContext
	Variant(i int) IVariantContext

	// IsEnumContext differentiates from other interfaces.
	IsEnumContext()
}

type EnumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumContext() *EnumContext {
	p := new(EnumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_enum
	return p
}

func InitEmptyEnumContext(p *EnumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_enum
}

func (*EnumContext) IsEnumContext() {}

func NewEnumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumContext {
	p := new(EnumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_enum

	return p
}

func (s *EnumContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumContext) ENUM() antlr.TerminalNode {
	return s.GetToken(LitebufParserENUM, 0)
}

func (s *EnumContext) Identifier() IIdentifierContext {
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

func (s *EnumContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserLBRACE, 0)
}

func (s *EnumContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserRBRACE, 0)
}

func (s *EnumContext) AllVariant() []IVariantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariantContext); ok {
			len++
		}
	}

	tst := make([]IVariantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariantContext); ok {
			tst[i] = t.(IVariantContext)
			i++
		}
	}

	return tst
}

func (s *EnumContext) Variant(i int) IVariantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariantContext); ok {
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

	return t.(IVariantContext)
}

func (s *EnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterEnum(s)
	}
}

func (s *EnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitEnum(s)
	}
}

func (s *EnumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitEnum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Enum() (localctx IEnumContext) {
	localctx = NewEnumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LitebufParserRULE_enum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(LitebufParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Identifier()
	}
	{
		p.SetState(94)
		p.Match(LitebufParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LitebufParserIDENTIFIER {
		{
			p.SetState(95)
			p.Variant()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Match(LitebufParserRBRACE)
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

// IVariantContext is an interface to support dynamic dispatch.
type IVariantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Order() IOrderContext

	// IsVariantContext differentiates from other interfaces.
	IsVariantContext()
}

type VariantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariantContext() *VariantContext {
	p := new(VariantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_variant
	return p
}

func InitEmptyVariantContext(p *VariantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_variant
}

func (*VariantContext) IsVariantContext() {}

func NewVariantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariantContext {
	p := new(VariantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_variant

	return p
}

func (s *VariantContext) GetParser() antlr.Parser { return s.parser }

func (s *VariantContext) Identifier() IIdentifierContext {
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

func (s *VariantContext) Order() IOrderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderContext)
}

func (s *VariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterVariant(s)
	}
}

func (s *VariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitVariant(s)
	}
}

func (s *VariantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitVariant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Variant() (localctx IVariantContext) {
	localctx = NewVariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LitebufParserRULE_variant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Identifier()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LitebufParserEQUALS {
		{
			p.SetState(103)
			p.Order()
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

// IOrderContext is an interface to support dynamic dispatch.
type IOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUALS() antlr.TerminalNode
	ORDER() antlr.TerminalNode

	// IsOrderContext differentiates from other interfaces.
	IsOrderContext()
}

type OrderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderContext() *OrderContext {
	p := new(OrderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_order
	return p
}

func InitEmptyOrderContext(p *OrderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_order
}

func (*OrderContext) IsOrderContext() {}

func NewOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderContext {
	p := new(OrderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_order

	return p
}

func (s *OrderContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(LitebufParserEQUALS, 0)
}

func (s *OrderContext) ORDER() antlr.TerminalNode {
	return s.GetToken(LitebufParserORDER, 0)
}

func (s *OrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterOrder(s)
	}
}

func (s *OrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitOrder(s)
	}
}

func (s *OrderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitOrder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Order() (localctx IOrderContext) {
	localctx = NewOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LitebufParserRULE_order)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(LitebufParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(LitebufParserORDER)
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

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ROOT() antlr.TerminalNode
	Identifier() IIdentifierContext
	Def() IDefContext

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	p := new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	p := new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) ROOT() antlr.TerminalNode {
	return s.GetToken(LitebufParserROOT, 0)
}

func (s *RootContext) Identifier() IIdentifierContext {
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

func (s *RootContext) Def() IDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LitebufParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(LitebufParserROOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LitebufParserIDENTIFIER:
		{
			p.SetState(110)
			p.Identifier()
		}

	case LitebufParserSTRUCT, LitebufParserUNION, LitebufParserENUM:
		{
			p.SetState(111)
			p.Def()
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

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	Identifier() IIdentifierContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFn() []IFnContext
	Fn(i int) IFnContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	p := new(ServiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_service
	return p
}

func InitEmptyServiceContext(p *ServiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_service
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	p := new(ServiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(LitebufParserSERVICE, 0)
}

func (s *ServiceContext) Identifier() IIdentifierContext {
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

func (s *ServiceContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserLBRACE, 0)
}

func (s *ServiceContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LitebufParserRBRACE, 0)
}

func (s *ServiceContext) AllFn() []IFnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFnContext); ok {
			len++
		}
	}

	tst := make([]IFnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFnContext); ok {
			tst[i] = t.(IFnContext)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) Fn(i int) IFnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnContext); ok {
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

	return t.(IFnContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitService(s)
	}
}

func (s *ServiceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitService(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LitebufParserRULE_service)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(LitebufParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Identifier()
	}
	{
		p.SetState(116)
		p.Match(LitebufParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LitebufParserIDENTIFIER {
		{
			p.SetState(117)
			p.Fn()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
		p.Match(LitebufParserRBRACE)
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

// IFnContext is an interface to support dynamic dispatch.
type IFnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	COLON() antlr.TerminalNode
	RARROW() antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllDef() []IDefContext
	Def(i int) IDefContext

	// IsFnContext differentiates from other interfaces.
	IsFnContext()
}

type FnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnContext() *FnContext {
	p := new(FnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_fn
	return p
}

func InitEmptyFnContext(p *FnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_fn
}

func (*FnContext) IsFnContext() {}

func NewFnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnContext {
	p := new(FnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_fn

	return p
}

func (s *FnContext) GetParser() antlr.Parser { return s.parser }

func (s *FnContext) AllIdentifier() []IIdentifierContext {
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

func (s *FnContext) Identifier(i int) IIdentifierContext {
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

func (s *FnContext) COLON() antlr.TerminalNode {
	return s.GetToken(LitebufParserCOLON, 0)
}

func (s *FnContext) RARROW() antlr.TerminalNode {
	return s.GetToken(LitebufParserRARROW, 0)
}

func (s *FnContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *FnContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *FnContext) AllDef() []IDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefContext); ok {
			len++
		}
	}

	tst := make([]IDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefContext); ok {
			tst[i] = t.(IDefContext)
			i++
		}
	}

	return tst
}

func (s *FnContext) Def(i int) IDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
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

	return t.(IDefContext)
}

func (s *FnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterFn(s)
	}
}

func (s *FnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitFn(s)
	}
}

func (s *FnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitFn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Fn() (localctx IFnContext) {
	localctx = NewFnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LitebufParserRULE_fn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Identifier()
	}
	{
		p.SetState(126)
		p.Match(LitebufParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(127)
			p.Identifier()
		}

	case 2:
		{
			p.SetState(128)
			p.Type_()
		}

	case 3:
		{
			p.SetState(129)
			p.Def()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(132)
		p.Match(LitebufParserRARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(133)
			p.Identifier()
		}

	case 2:
		{
			p.SetState(134)
			p.Type_()
		}

	case 3:
		{
			p.SetState(135)
			p.Def()
		}

	case antlr.ATNInvalidAltNumber:
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

// IChannelContext is an interface to support dynamic dispatch.
type IChannelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHANNEL() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	Type_() ITypeContext
	Def() IDefContext

	// IsChannelContext differentiates from other interfaces.
	IsChannelContext()
}

type ChannelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChannelContext() *ChannelContext {
	p := new(ChannelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_channel
	return p
}

func InitEmptyChannelContext(p *ChannelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_channel
}

func (*ChannelContext) IsChannelContext() {}

func NewChannelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChannelContext {
	p := new(ChannelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_channel

	return p
}

func (s *ChannelContext) GetParser() antlr.Parser { return s.parser }

func (s *ChannelContext) CHANNEL() antlr.TerminalNode {
	return s.GetToken(LitebufParserCHANNEL, 0)
}

func (s *ChannelContext) AllIdentifier() []IIdentifierContext {
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

func (s *ChannelContext) Identifier(i int) IIdentifierContext {
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

func (s *ChannelContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ChannelContext) Def() IDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefContext)
}

func (s *ChannelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChannelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChannelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterChannel(s)
	}
}

func (s *ChannelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitChannel(s)
	}
}

func (s *ChannelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitChannel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Channel() (localctx IChannelContext) {
	localctx = NewChannelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LitebufParserRULE_channel)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(LitebufParserCHANNEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Identifier()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(140)
			p.Identifier()
		}

	case 2:
		{
			p.SetState(141)
			p.Type_()
		}

	case 3:
		{
			p.SetState(142)
			p.Def()
		}

	case antlr.ATNInvalidAltNumber:
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Primitive() IPrimitiveContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	p := new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	p := new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Identifier() IIdentifierContext {
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

func (s *TypeContext) Primitive() IPrimitiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LitebufParserRULE_type)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LitebufParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Identifier()
		}

	case LitebufParserBOOLEAN, LitebufParserINT8, LitebufParserUINT8, LitebufParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Primitive()
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

// IPrimitiveContext is an interface to support dynamic dispatch.
type IPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode
	INT8() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveContext() *PrimitiveContext {
	p := new(PrimitiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_primitive
	return p
}

func InitEmptyPrimitiveContext(p *PrimitiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_primitive
}

func (*PrimitiveContext) IsPrimitiveContext() {}

func NewPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveContext {
	p := new(PrimitiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_primitive

	return p
}

func (s *PrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(LitebufParserBOOLEAN, 0)
}

func (s *PrimitiveContext) INT8() antlr.TerminalNode {
	return s.GetToken(LitebufParserINT8, 0)
}

func (s *PrimitiveContext) UINT8() antlr.TerminalNode {
	return s.GetToken(LitebufParserUINT8, 0)
}

func (s *PrimitiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(LitebufParserSTRING, 0)
}

func (s *PrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterPrimitive(s)
	}
}

func (s *PrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitPrimitive(s)
	}
}

func (s *PrimitiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitPrimitive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Primitive() (localctx IPrimitiveContext) {
	localctx = NewPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LitebufParserRULE_primitive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	p := new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LitebufParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	p := new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LitebufParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LitebufParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LitebufListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LitebufVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LitebufParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LitebufParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(LitebufParserIDENTIFIER)
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
