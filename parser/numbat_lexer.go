// Code generated from Numbat.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NumbatLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NumbatLexerLexerStaticData struct {
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

func numbatlexerLexerInit() {
	staticData := &NumbatLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'#'", "'true'", "'false'", "'null'", "'='", "'('", "','", "')'",
		"'do'", "'end'", "'proc'", "':'", "'var'", "'return'", "'program'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NEWLINE",
		"NUMBER", "HEX", "STRING", "TYPE_NAME", "NON_TYPE_NAME", "WS", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "NEWLINE", "NUMBER",
		"HEX", "STRING", "TYPE_NAME", "NON_TYPE_NAME", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 206, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 4,
		15, 110, 8, 15, 11, 15, 12, 15, 111, 1, 16, 3, 16, 115, 8, 16, 1, 16, 4,
		16, 118, 8, 16, 11, 16, 12, 16, 119, 1, 16, 1, 16, 4, 16, 124, 8, 16, 11,
		16, 12, 16, 125, 3, 16, 128, 8, 16, 1, 16, 1, 16, 3, 16, 132, 8, 16, 1,
		16, 4, 16, 135, 8, 16, 11, 16, 12, 16, 136, 3, 16, 139, 8, 16, 1, 17, 3,
		17, 142, 8, 17, 1, 17, 1, 17, 1, 17, 4, 17, 147, 8, 17, 11, 17, 12, 17,
		148, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 155, 8, 18, 10, 18, 12, 18, 158,
		9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 4, 19, 164, 8, 19, 11, 19, 12, 19, 165,
		1, 20, 1, 20, 5, 20, 170, 8, 20, 10, 20, 12, 20, 173, 9, 20, 1, 21, 4,
		21, 176, 8, 21, 11, 21, 12, 21, 177, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 5, 22, 186, 8, 22, 10, 22, 12, 22, 189, 9, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 200, 8, 23, 10, 23,
		12, 23, 203, 9, 23, 1, 23, 1, 23, 1, 187, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 1, 0, 12, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 1, 0, 101, 101, 2,
		0, 43, 43, 45, 45, 1, 0, 120, 120, 2, 0, 48, 57, 65, 70, 2, 0, 34, 34,
		92, 92, 1, 0, 65, 90, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 97,
		122, 3, 0, 48, 57, 95, 95, 97, 122, 2, 0, 9, 9, 32, 32, 222, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 1, 49, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 56, 1, 0, 0, 0, 7, 62, 1, 0,
		0, 0, 9, 67, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13, 71, 1, 0, 0, 0, 15, 73,
		1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 78, 1, 0, 0, 0, 21, 82, 1, 0, 0, 0,
		23, 87, 1, 0, 0, 0, 25, 89, 1, 0, 0, 0, 27, 93, 1, 0, 0, 0, 29, 100, 1,
		0, 0, 0, 31, 109, 1, 0, 0, 0, 33, 114, 1, 0, 0, 0, 35, 141, 1, 0, 0, 0,
		37, 150, 1, 0, 0, 0, 39, 161, 1, 0, 0, 0, 41, 167, 1, 0, 0, 0, 43, 175,
		1, 0, 0, 0, 45, 181, 1, 0, 0, 0, 47, 195, 1, 0, 0, 0, 49, 50, 5, 35, 0,
		0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 116, 0, 0, 52, 53, 5, 114, 0, 0, 53, 54,
		5, 117, 0, 0, 54, 55, 5, 101, 0, 0, 55, 4, 1, 0, 0, 0, 56, 57, 5, 102,
		0, 0, 57, 58, 5, 97, 0, 0, 58, 59, 5, 108, 0, 0, 59, 60, 5, 115, 0, 0,
		60, 61, 5, 101, 0, 0, 61, 6, 1, 0, 0, 0, 62, 63, 5, 110, 0, 0, 63, 64,
		5, 117, 0, 0, 64, 65, 5, 108, 0, 0, 65, 66, 5, 108, 0, 0, 66, 8, 1, 0,
		0, 0, 67, 68, 5, 61, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70, 5, 40, 0, 0, 70,
		12, 1, 0, 0, 0, 71, 72, 5, 44, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 41,
		0, 0, 74, 16, 1, 0, 0, 0, 75, 76, 5, 100, 0, 0, 76, 77, 5, 111, 0, 0, 77,
		18, 1, 0, 0, 0, 78, 79, 5, 101, 0, 0, 79, 80, 5, 110, 0, 0, 80, 81, 5,
		100, 0, 0, 81, 20, 1, 0, 0, 0, 82, 83, 5, 112, 0, 0, 83, 84, 5, 114, 0,
		0, 84, 85, 5, 111, 0, 0, 85, 86, 5, 99, 0, 0, 86, 22, 1, 0, 0, 0, 87, 88,
		5, 58, 0, 0, 88, 24, 1, 0, 0, 0, 89, 90, 5, 118, 0, 0, 90, 91, 5, 97, 0,
		0, 91, 92, 5, 114, 0, 0, 92, 26, 1, 0, 0, 0, 93, 94, 5, 114, 0, 0, 94,
		95, 5, 101, 0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 117, 0, 0, 97, 98, 5,
		114, 0, 0, 98, 99, 5, 110, 0, 0, 99, 28, 1, 0, 0, 0, 100, 101, 5, 112,
		0, 0, 101, 102, 5, 114, 0, 0, 102, 103, 5, 111, 0, 0, 103, 104, 5, 103,
		0, 0, 104, 105, 5, 114, 0, 0, 105, 106, 5, 97, 0, 0, 106, 107, 5, 109,
		0, 0, 107, 30, 1, 0, 0, 0, 108, 110, 7, 0, 0, 0, 109, 108, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112,
		32, 1, 0, 0, 0, 113, 115, 5, 45, 0, 0, 114, 113, 1, 0, 0, 0, 114, 115,
		1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 118, 7, 1, 0, 0, 117, 116, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0,
		120, 127, 1, 0, 0, 0, 121, 123, 5, 46, 0, 0, 122, 124, 7, 1, 0, 0, 123,
		122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126,
		1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 121, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 138, 1, 0, 0, 0, 129, 131, 7, 2, 0, 0, 130, 132, 7, 3, 0, 0,
		131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133,
		135, 7, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 134,
		1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 129, 1, 0,
		0, 0, 138, 139, 1, 0, 0, 0, 139, 34, 1, 0, 0, 0, 140, 142, 5, 45, 0, 0,
		141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143,
		144, 5, 48, 0, 0, 144, 146, 7, 4, 0, 0, 145, 147, 7, 5, 0, 0, 146, 145,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 36, 1, 0, 0, 0, 150, 156, 5, 34, 0, 0, 151, 155, 8, 6, 0, 0,
		152, 153, 5, 92, 0, 0, 153, 155, 9, 0, 0, 0, 154, 151, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157,
		1, 0, 0, 0, 157, 159, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 160, 5, 34,
		0, 0, 160, 38, 1, 0, 0, 0, 161, 163, 7, 7, 0, 0, 162, 164, 7, 8, 0, 0,
		163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 40, 1, 0, 0, 0, 167, 171, 7, 9, 0, 0, 168, 170, 7,
		10, 0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0,
		0, 171, 172, 1, 0, 0, 0, 172, 42, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174,
		176, 7, 11, 0, 0, 175, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 175,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 6, 21,
		0, 0, 180, 44, 1, 0, 0, 0, 181, 182, 5, 47, 0, 0, 182, 183, 5, 42, 0, 0,
		183, 187, 1, 0, 0, 0, 184, 186, 9, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186,
		189, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 190,
		1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 42, 0, 0, 191, 192, 5, 47,
		0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 6, 22, 0, 0, 194, 46, 1, 0, 0, 0,
		195, 196, 5, 47, 0, 0, 196, 197, 5, 47, 0, 0, 197, 201, 1, 0, 0, 0, 198,
		200, 8, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199,
		1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0,
		0, 0, 204, 205, 6, 23, 0, 0, 205, 48, 1, 0, 0, 0, 18, 0, 111, 114, 119,
		125, 127, 131, 136, 138, 141, 148, 154, 156, 165, 171, 177, 187, 201, 1,
		6, 0, 0,
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

// NumbatLexerInit initializes any static state used to implement NumbatLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumbatLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumbatLexerInit() {
	staticData := &NumbatLexerLexerStaticData
	staticData.once.Do(numbatlexerLexerInit)
}

// NewNumbatLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumbatLexer(input antlr.CharStream) *NumbatLexer {
	NumbatLexerInit()
	l := new(NumbatLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NumbatLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Numbat.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumbatLexer tokens.
const (
	NumbatLexerT__0          = 1
	NumbatLexerT__1          = 2
	NumbatLexerT__2          = 3
	NumbatLexerT__3          = 4
	NumbatLexerT__4          = 5
	NumbatLexerT__5          = 6
	NumbatLexerT__6          = 7
	NumbatLexerT__7          = 8
	NumbatLexerT__8          = 9
	NumbatLexerT__9          = 10
	NumbatLexerT__10         = 11
	NumbatLexerT__11         = 12
	NumbatLexerT__12         = 13
	NumbatLexerT__13         = 14
	NumbatLexerT__14         = 15
	NumbatLexerNEWLINE       = 16
	NumbatLexerNUMBER        = 17
	NumbatLexerHEX           = 18
	NumbatLexerSTRING        = 19
	NumbatLexerTYPE_NAME     = 20
	NumbatLexerNON_TYPE_NAME = 21
	NumbatLexerWS            = 22
	NumbatLexerCOMMENT       = 23
	NumbatLexerLINE_COMMENT  = 24
)
