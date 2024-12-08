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
		"", "'true'", "'false'", "'null'", "'='", "'('", "','", "')'", "'do'",
		"'end'", "'proc'", "':'", "'let'", "'return'", "'program'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NEWLINE",
		"NUMBER", "HEX", "STRING", "TYPE_NAME", "NON_TYPE_NAME", "WS", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "NEWLINE", "NUMBER", "HEX",
		"STRING", "TYPE_NAME", "NON_TYPE_NAME", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 202, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 4, 14, 106, 8, 14, 11, 14, 12,
		14, 107, 1, 15, 3, 15, 111, 8, 15, 1, 15, 4, 15, 114, 8, 15, 11, 15, 12,
		15, 115, 1, 15, 1, 15, 4, 15, 120, 8, 15, 11, 15, 12, 15, 121, 3, 15, 124,
		8, 15, 1, 15, 1, 15, 3, 15, 128, 8, 15, 1, 15, 4, 15, 131, 8, 15, 11, 15,
		12, 15, 132, 3, 15, 135, 8, 15, 1, 16, 3, 16, 138, 8, 16, 1, 16, 1, 16,
		1, 16, 4, 16, 143, 8, 16, 11, 16, 12, 16, 144, 1, 17, 1, 17, 1, 17, 1,
		17, 5, 17, 151, 8, 17, 10, 17, 12, 17, 154, 9, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 4, 18, 160, 8, 18, 11, 18, 12, 18, 161, 1, 19, 1, 19, 5, 19, 166,
		8, 19, 10, 19, 12, 19, 169, 9, 19, 1, 20, 4, 20, 172, 8, 20, 11, 20, 12,
		20, 173, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 182, 8, 21, 10,
		21, 12, 21, 185, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 5, 22, 196, 8, 22, 10, 22, 12, 22, 199, 9, 22, 1, 22, 1,
		22, 1, 183, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 1, 0, 12, 2, 0, 10, 10, 13,
		13, 1, 0, 48, 57, 1, 0, 101, 101, 2, 0, 43, 43, 45, 45, 1, 0, 120, 120,
		2, 0, 48, 57, 65, 70, 2, 0, 34, 34, 92, 92, 1, 0, 65, 90, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 1, 0, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122,
		2, 0, 9, 9, 32, 32, 218, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 52, 1, 0, 0, 0, 5, 58, 1, 0,
		0, 0, 7, 63, 1, 0, 0, 0, 9, 65, 1, 0, 0, 0, 11, 67, 1, 0, 0, 0, 13, 69,
		1, 0, 0, 0, 15, 71, 1, 0, 0, 0, 17, 74, 1, 0, 0, 0, 19, 78, 1, 0, 0, 0,
		21, 83, 1, 0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 89, 1, 0, 0, 0, 27, 96, 1,
		0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 110, 1, 0, 0, 0, 33, 137, 1, 0, 0, 0,
		35, 146, 1, 0, 0, 0, 37, 157, 1, 0, 0, 0, 39, 163, 1, 0, 0, 0, 41, 171,
		1, 0, 0, 0, 43, 177, 1, 0, 0, 0, 45, 191, 1, 0, 0, 0, 47, 48, 5, 116, 0,
		0, 48, 49, 5, 114, 0, 0, 49, 50, 5, 117, 0, 0, 50, 51, 5, 101, 0, 0, 51,
		2, 1, 0, 0, 0, 52, 53, 5, 102, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 108,
		0, 0, 55, 56, 5, 115, 0, 0, 56, 57, 5, 101, 0, 0, 57, 4, 1, 0, 0, 0, 58,
		59, 5, 110, 0, 0, 59, 60, 5, 117, 0, 0, 60, 61, 5, 108, 0, 0, 61, 62, 5,
		108, 0, 0, 62, 6, 1, 0, 0, 0, 63, 64, 5, 61, 0, 0, 64, 8, 1, 0, 0, 0, 65,
		66, 5, 40, 0, 0, 66, 10, 1, 0, 0, 0, 67, 68, 5, 44, 0, 0, 68, 12, 1, 0,
		0, 0, 69, 70, 5, 41, 0, 0, 70, 14, 1, 0, 0, 0, 71, 72, 5, 100, 0, 0, 72,
		73, 5, 111, 0, 0, 73, 16, 1, 0, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5,
		110, 0, 0, 76, 77, 5, 100, 0, 0, 77, 18, 1, 0, 0, 0, 78, 79, 5, 112, 0,
		0, 79, 80, 5, 114, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 99, 0, 0, 82,
		20, 1, 0, 0, 0, 83, 84, 5, 58, 0, 0, 84, 22, 1, 0, 0, 0, 85, 86, 5, 108,
		0, 0, 86, 87, 5, 101, 0, 0, 87, 88, 5, 116, 0, 0, 88, 24, 1, 0, 0, 0, 89,
		90, 5, 114, 0, 0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5,
		117, 0, 0, 93, 94, 5, 114, 0, 0, 94, 95, 5, 110, 0, 0, 95, 26, 1, 0, 0,
		0, 96, 97, 5, 112, 0, 0, 97, 98, 5, 114, 0, 0, 98, 99, 5, 111, 0, 0, 99,
		100, 5, 103, 0, 0, 100, 101, 5, 114, 0, 0, 101, 102, 5, 97, 0, 0, 102,
		103, 5, 109, 0, 0, 103, 28, 1, 0, 0, 0, 104, 106, 7, 0, 0, 0, 105, 104,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 30, 1, 0, 0, 0, 109, 111, 5, 45, 0, 0, 110, 109, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 114, 7, 1, 0, 0, 113,
		112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 123, 1, 0, 0, 0, 117, 119, 5, 46, 0, 0, 118, 120, 7, 1,
		0, 0, 119, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 117, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 134, 1, 0, 0, 0, 125, 127, 7, 2, 0, 0, 126, 128,
		7, 3, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130, 1, 0,
		0, 0, 129, 131, 7, 1, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134,
		125, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 32, 1, 0, 0, 0, 136, 138, 5,
		45, 0, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0,
		0, 139, 140, 5, 48, 0, 0, 140, 142, 7, 4, 0, 0, 141, 143, 7, 5, 0, 0, 142,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145,
		1, 0, 0, 0, 145, 34, 1, 0, 0, 0, 146, 152, 5, 34, 0, 0, 147, 151, 8, 6,
		0, 0, 148, 149, 5, 92, 0, 0, 149, 151, 9, 0, 0, 0, 150, 147, 1, 0, 0, 0,
		150, 148, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 155, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156,
		5, 34, 0, 0, 156, 36, 1, 0, 0, 0, 157, 159, 7, 7, 0, 0, 158, 160, 7, 8,
		0, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 38, 1, 0, 0, 0, 163, 167, 7, 9, 0, 0, 164, 166,
		7, 10, 0, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0,
		0, 0, 167, 168, 1, 0, 0, 0, 168, 40, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0,
		170, 172, 7, 11, 0, 0, 171, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176,
		6, 20, 0, 0, 176, 42, 1, 0, 0, 0, 177, 178, 5, 47, 0, 0, 178, 179, 5, 42,
		0, 0, 179, 183, 1, 0, 0, 0, 180, 182, 9, 0, 0, 0, 181, 180, 1, 0, 0, 0,
		182, 185, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184,
		186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 42, 0, 0, 187, 188,
		5, 47, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 6, 21, 0, 0, 190, 44, 1, 0,
		0, 0, 191, 192, 5, 47, 0, 0, 192, 193, 5, 47, 0, 0, 193, 197, 1, 0, 0,
		0, 194, 196, 8, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197,
		195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 197,
		1, 0, 0, 0, 200, 201, 6, 22, 0, 0, 201, 46, 1, 0, 0, 0, 18, 0, 107, 110,
		115, 121, 123, 127, 132, 134, 137, 144, 150, 152, 161, 167, 173, 183, 197,
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
	NumbatLexerNEWLINE       = 15
	NumbatLexerNUMBER        = 16
	NumbatLexerHEX           = 17
	NumbatLexerSTRING        = 18
	NumbatLexerTYPE_NAME     = 19
	NumbatLexerNON_TYPE_NAME = 20
	NumbatLexerWS            = 21
	NumbatLexerCOMMENT       = 22
	NumbatLexerLINE_COMMENT  = 23
)
