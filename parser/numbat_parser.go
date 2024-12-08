// Code generated from Numbat.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Numbat

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

type NumbatParser struct {
	*antlr.BaseParser
}

var NumbatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func numbatParserInit() {
	staticData := &NumbatParserStaticData
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
		"prog", "expr_bool", "expr_num", "expr_hex", "expr_str", "expr_null",
		"expr_constant", "expr_var", "expr_call", "expr_all", "param_expr",
		"param_type", "param", "type_out", "type_in", "type", "proc_body", "proc_name",
		"proc_type", "proc_def", "proc", "call_expr", "call_secondary", "call_primary",
		"call", "call_stmt", "let_expr", "let_var_type", "let_var_name", "let",
		"assignment_expr", "assignment_var", "assignment", "return_expr", "return",
		"return_stmt", "program", "statement", "object",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 245, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 5, 0, 82, 8, 0, 10, 0, 12,
		0, 85, 9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 102, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 3, 9, 111, 8, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 3, 12, 120, 8, 12, 1, 12, 3, 12, 123, 8, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 5, 14, 131, 8, 14, 10, 14, 12, 14, 134, 9, 14,
		1, 14, 1, 14, 1, 15, 3, 15, 139, 8, 15, 1, 15, 1, 15, 1, 15, 3, 15, 144,
		8, 15, 3, 15, 146, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 152, 8, 16,
		10, 16, 12, 16, 155, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 3, 19, 166, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 4, 21, 174, 8, 21, 11, 21, 12, 21, 175, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 3, 24, 185, 8, 24, 1, 24, 3, 24, 188, 8, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 3, 29, 204, 8, 29, 1, 29, 3, 29, 207, 8, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 3, 30, 213, 8, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 3, 32, 220, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 3, 34,
		228, 8, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 3, 37, 239, 8, 37, 1, 38, 1, 38, 3, 38, 243, 8, 38, 1, 38, 0, 0, 39,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 0, 2, 1, 0, 1, 2, 1, 0, 19, 20, 237, 0, 83, 1, 0, 0, 0, 2, 86,
		1, 0, 0, 0, 4, 88, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 92, 1, 0, 0, 0, 10,
		94, 1, 0, 0, 0, 12, 101, 1, 0, 0, 0, 14, 103, 1, 0, 0, 0, 16, 105, 1, 0,
		0, 0, 18, 110, 1, 0, 0, 0, 20, 112, 1, 0, 0, 0, 22, 115, 1, 0, 0, 0, 24,
		117, 1, 0, 0, 0, 26, 124, 1, 0, 0, 0, 28, 126, 1, 0, 0, 0, 30, 145, 1,
		0, 0, 0, 32, 147, 1, 0, 0, 0, 34, 158, 1, 0, 0, 0, 36, 160, 1, 0, 0, 0,
		38, 162, 1, 0, 0, 0, 40, 167, 1, 0, 0, 0, 42, 170, 1, 0, 0, 0, 44, 177,
		1, 0, 0, 0, 46, 179, 1, 0, 0, 0, 48, 181, 1, 0, 0, 0, 50, 191, 1, 0, 0,
		0, 52, 193, 1, 0, 0, 0, 54, 196, 1, 0, 0, 0, 56, 198, 1, 0, 0, 0, 58, 200,
		1, 0, 0, 0, 60, 208, 1, 0, 0, 0, 62, 214, 1, 0, 0, 0, 64, 216, 1, 0, 0,
		0, 66, 223, 1, 0, 0, 0, 68, 225, 1, 0, 0, 0, 70, 229, 1, 0, 0, 0, 72, 231,
		1, 0, 0, 0, 74, 238, 1, 0, 0, 0, 76, 242, 1, 0, 0, 0, 78, 82, 5, 21, 0,
		0, 79, 82, 5, 15, 0, 0, 80, 82, 3, 76, 38, 0, 81, 78, 1, 0, 0, 0, 81, 79,
		1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		83, 84, 1, 0, 0, 0, 84, 1, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 7, 0,
		0, 0, 87, 3, 1, 0, 0, 0, 88, 89, 5, 16, 0, 0, 89, 5, 1, 0, 0, 0, 90, 91,
		5, 17, 0, 0, 91, 7, 1, 0, 0, 0, 92, 93, 5, 18, 0, 0, 93, 9, 1, 0, 0, 0,
		94, 95, 5, 3, 0, 0, 95, 11, 1, 0, 0, 0, 96, 102, 3, 2, 1, 0, 97, 102, 3,
		4, 2, 0, 98, 102, 3, 6, 3, 0, 99, 102, 3, 8, 4, 0, 100, 102, 3, 10, 5,
		0, 101, 96, 1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101,
		99, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102, 13, 1, 0, 0, 0, 103, 104, 5,
		20, 0, 0, 104, 15, 1, 0, 0, 0, 105, 106, 3, 48, 24, 0, 106, 17, 1, 0, 0,
		0, 107, 111, 3, 14, 7, 0, 108, 111, 3, 12, 6, 0, 109, 111, 3, 16, 8, 0,
		110, 107, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111,
		19, 1, 0, 0, 0, 112, 113, 5, 4, 0, 0, 113, 114, 3, 12, 6, 0, 114, 21, 1,
		0, 0, 0, 115, 116, 3, 30, 15, 0, 116, 23, 1, 0, 0, 0, 117, 119, 5, 20,
		0, 0, 118, 120, 3, 22, 11, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0,
		0, 120, 122, 1, 0, 0, 0, 121, 123, 3, 20, 10, 0, 122, 121, 1, 0, 0, 0,
		122, 123, 1, 0, 0, 0, 123, 25, 1, 0, 0, 0, 124, 125, 5, 19, 0, 0, 125,
		27, 1, 0, 0, 0, 126, 127, 5, 5, 0, 0, 127, 132, 3, 24, 12, 0, 128, 129,
		5, 6, 0, 0, 129, 131, 3, 24, 12, 0, 130, 128, 1, 0, 0, 0, 131, 134, 1,
		0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0,
		0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 7, 0, 0, 136, 29, 1, 0, 0, 0, 137,
		139, 3, 28, 14, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 146, 3, 26, 13, 0, 141, 143, 3, 28, 14, 0, 142, 144, 3,
		26, 13, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0,
		0, 0, 145, 138, 1, 0, 0, 0, 145, 141, 1, 0, 0, 0, 146, 31, 1, 0, 0, 0,
		147, 153, 5, 8, 0, 0, 148, 152, 5, 21, 0, 0, 149, 152, 5, 15, 0, 0, 150,
		152, 3, 74, 37, 0, 151, 148, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 150,
		1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0,
		0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 5, 9, 0, 0,
		157, 33, 1, 0, 0, 0, 158, 159, 5, 20, 0, 0, 159, 35, 1, 0, 0, 0, 160, 161,
		3, 30, 15, 0, 161, 37, 1, 0, 0, 0, 162, 163, 5, 10, 0, 0, 163, 165, 3,
		34, 17, 0, 164, 166, 3, 36, 18, 0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0,
		0, 0, 166, 39, 1, 0, 0, 0, 167, 168, 3, 38, 19, 0, 168, 169, 3, 32, 16,
		0, 169, 41, 1, 0, 0, 0, 170, 173, 5, 11, 0, 0, 171, 174, 3, 14, 7, 0, 172,
		174, 3, 12, 6, 0, 173, 171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 43, 1, 0,
		0, 0, 177, 178, 5, 20, 0, 0, 178, 45, 1, 0, 0, 0, 179, 180, 7, 1, 0, 0,
		180, 47, 1, 0, 0, 0, 181, 182, 5, 5, 0, 0, 182, 184, 3, 46, 23, 0, 183,
		185, 3, 44, 22, 0, 184, 183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187,
		1, 0, 0, 0, 186, 188, 3, 42, 21, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1,
		0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 5, 7, 0, 0, 190, 49, 1, 0, 0,
		0, 191, 192, 3, 48, 24, 0, 192, 51, 1, 0, 0, 0, 193, 194, 5, 4, 0, 0, 194,
		195, 3, 18, 9, 0, 195, 53, 1, 0, 0, 0, 196, 197, 3, 30, 15, 0, 197, 55,
		1, 0, 0, 0, 198, 199, 5, 20, 0, 0, 199, 57, 1, 0, 0, 0, 200, 201, 5, 12,
		0, 0, 201, 203, 3, 56, 28, 0, 202, 204, 3, 54, 27, 0, 203, 202, 1, 0, 0,
		0, 203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 207, 3, 52, 26, 0,
		206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 59, 1, 0, 0, 0, 208, 209,
		5, 4, 0, 0, 209, 212, 3, 18, 9, 0, 210, 211, 5, 6, 0, 0, 211, 213, 3, 18,
		9, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 61, 1, 0, 0, 0,
		214, 215, 5, 20, 0, 0, 215, 63, 1, 0, 0, 0, 216, 219, 3, 62, 31, 0, 217,
		218, 5, 6, 0, 0, 218, 220, 3, 62, 31, 0, 219, 217, 1, 0, 0, 0, 219, 220,
		1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 222, 3, 60, 30, 0, 222, 65, 1, 0,
		0, 0, 223, 224, 3, 18, 9, 0, 224, 67, 1, 0, 0, 0, 225, 227, 5, 13, 0, 0,
		226, 228, 3, 66, 33, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		69, 1, 0, 0, 0, 229, 230, 3, 68, 34, 0, 230, 71, 1, 0, 0, 0, 231, 232,
		5, 14, 0, 0, 232, 233, 3, 32, 16, 0, 233, 73, 1, 0, 0, 0, 234, 239, 3,
		50, 25, 0, 235, 239, 3, 58, 29, 0, 236, 239, 3, 70, 35, 0, 237, 239, 3,
		64, 32, 0, 238, 234, 1, 0, 0, 0, 238, 235, 1, 0, 0, 0, 238, 236, 1, 0,
		0, 0, 238, 237, 1, 0, 0, 0, 239, 75, 1, 0, 0, 0, 240, 243, 3, 72, 36, 0,
		241, 243, 3, 40, 20, 0, 242, 240, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243,
		77, 1, 0, 0, 0, 24, 81, 83, 101, 110, 119, 122, 132, 138, 143, 145, 151,
		153, 165, 173, 175, 184, 187, 203, 206, 212, 219, 227, 238, 242,
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

// NumbatParserInit initializes any static state used to implement NumbatParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNumbatParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumbatParserInit() {
	staticData := &NumbatParserStaticData
	staticData.once.Do(numbatParserInit)
}

// NewNumbatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNumbatParser(input antlr.TokenStream) *NumbatParser {
	NumbatParserInit()
	this := new(NumbatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &NumbatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Numbat.g4"

	return this
}

// NumbatParser tokens.
const (
	NumbatParserEOF           = antlr.TokenEOF
	NumbatParserT__0          = 1
	NumbatParserT__1          = 2
	NumbatParserT__2          = 3
	NumbatParserT__3          = 4
	NumbatParserT__4          = 5
	NumbatParserT__5          = 6
	NumbatParserT__6          = 7
	NumbatParserT__7          = 8
	NumbatParserT__8          = 9
	NumbatParserT__9          = 10
	NumbatParserT__10         = 11
	NumbatParserT__11         = 12
	NumbatParserT__12         = 13
	NumbatParserT__13         = 14
	NumbatParserNEWLINE       = 15
	NumbatParserNUMBER        = 16
	NumbatParserHEX           = 17
	NumbatParserSTRING        = 18
	NumbatParserTYPE_NAME     = 19
	NumbatParserNON_TYPE_NAME = 20
	NumbatParserWS            = 21
	NumbatParserCOMMENT       = 22
	NumbatParserLINE_COMMENT  = 23
)

// NumbatParser rules.
const (
	NumbatParserRULE_prog            = 0
	NumbatParserRULE_expr_bool       = 1
	NumbatParserRULE_expr_num        = 2
	NumbatParserRULE_expr_hex        = 3
	NumbatParserRULE_expr_str        = 4
	NumbatParserRULE_expr_null       = 5
	NumbatParserRULE_expr_constant   = 6
	NumbatParserRULE_expr_var        = 7
	NumbatParserRULE_expr_call       = 8
	NumbatParserRULE_expr_all        = 9
	NumbatParserRULE_param_expr      = 10
	NumbatParserRULE_param_type      = 11
	NumbatParserRULE_param           = 12
	NumbatParserRULE_type_out        = 13
	NumbatParserRULE_type_in         = 14
	NumbatParserRULE_type            = 15
	NumbatParserRULE_proc_body       = 16
	NumbatParserRULE_proc_name       = 17
	NumbatParserRULE_proc_type       = 18
	NumbatParserRULE_proc_def        = 19
	NumbatParserRULE_proc            = 20
	NumbatParserRULE_call_expr       = 21
	NumbatParserRULE_call_secondary  = 22
	NumbatParserRULE_call_primary    = 23
	NumbatParserRULE_call            = 24
	NumbatParserRULE_call_stmt       = 25
	NumbatParserRULE_let_expr        = 26
	NumbatParserRULE_let_var_type    = 27
	NumbatParserRULE_let_var_name    = 28
	NumbatParserRULE_let             = 29
	NumbatParserRULE_assignment_expr = 30
	NumbatParserRULE_assignment_var  = 31
	NumbatParserRULE_assignment      = 32
	NumbatParserRULE_return_expr     = 33
	NumbatParserRULE_return          = 34
	NumbatParserRULE_return_stmt     = 35
	NumbatParserRULE_program         = 36
	NumbatParserRULE_statement       = 37
	NumbatParserRULE_object          = 38
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllObject() []IObjectContext
	Object(i int) IObjectContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(NumbatParserWS)
}

func (s *ProgContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(NumbatParserWS, i)
}

func (s *ProgContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumbatParserNEWLINE)
}

func (s *ProgContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumbatParserNEWLINE, i)
}

func (s *ProgContext) AllObject() []IObjectContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectContext); ok {
			len++
		}
	}

	tst := make([]IObjectContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectContext); ok {
			tst[i] = t.(IObjectContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Object(i int) IObjectContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
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

	return t.(IObjectContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *NumbatParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NumbatParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147328) != 0 {
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NumbatParserWS:
			{
				p.SetState(78)
				p.Match(NumbatParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserNEWLINE:
			{
				p.SetState(79)
				p.Match(NumbatParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserT__9, NumbatParserT__13:
			{
				p.SetState(80)
				p.Object()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(85)
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

// IExpr_boolContext is an interface to support dynamic dispatch.
type IExpr_boolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpr_boolContext differentiates from other interfaces.
	IsExpr_boolContext()
}

type Expr_boolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_boolContext() *Expr_boolContext {
	var p = new(Expr_boolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_bool
	return p
}

func InitEmptyExpr_boolContext(p *Expr_boolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_bool
}

func (*Expr_boolContext) IsExpr_boolContext() {}

func NewExpr_boolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_boolContext {
	var p = new(Expr_boolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_bool

	return p
}

func (s *Expr_boolContext) GetParser() antlr.Parser { return s.parser }
func (s *Expr_boolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_boolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_boolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_bool(s)
	}
}

func (s *Expr_boolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_bool(s)
	}
}

func (p *NumbatParser) Expr_bool() (localctx IExpr_boolContext) {
	localctx = NewExpr_boolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NumbatParserRULE_expr_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NumbatParserT__0 || _la == NumbatParserT__1) {
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

// IExpr_numContext is an interface to support dynamic dispatch.
type IExpr_numContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsExpr_numContext differentiates from other interfaces.
	IsExpr_numContext()
}

type Expr_numContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_numContext() *Expr_numContext {
	var p = new(Expr_numContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_num
	return p
}

func InitEmptyExpr_numContext(p *Expr_numContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_num
}

func (*Expr_numContext) IsExpr_numContext() {}

func NewExpr_numContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_numContext {
	var p = new(Expr_numContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_num

	return p
}

func (s *Expr_numContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_numContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NumbatParserNUMBER, 0)
}

func (s *Expr_numContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_numContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_numContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_num(s)
	}
}

func (s *Expr_numContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_num(s)
	}
}

func (p *NumbatParser) Expr_num() (localctx IExpr_numContext) {
	localctx = NewExpr_numContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NumbatParserRULE_expr_num)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(NumbatParserNUMBER)
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

// IExpr_hexContext is an interface to support dynamic dispatch.
type IExpr_hexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HEX() antlr.TerminalNode

	// IsExpr_hexContext differentiates from other interfaces.
	IsExpr_hexContext()
}

type Expr_hexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_hexContext() *Expr_hexContext {
	var p = new(Expr_hexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_hex
	return p
}

func InitEmptyExpr_hexContext(p *Expr_hexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_hex
}

func (*Expr_hexContext) IsExpr_hexContext() {}

func NewExpr_hexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_hexContext {
	var p = new(Expr_hexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_hex

	return p
}

func (s *Expr_hexContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_hexContext) HEX() antlr.TerminalNode {
	return s.GetToken(NumbatParserHEX, 0)
}

func (s *Expr_hexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_hexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_hexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_hex(s)
	}
}

func (s *Expr_hexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_hex(s)
	}
}

func (p *NumbatParser) Expr_hex() (localctx IExpr_hexContext) {
	localctx = NewExpr_hexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NumbatParserRULE_expr_hex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(NumbatParserHEX)
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

// IExpr_strContext is an interface to support dynamic dispatch.
type IExpr_strContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsExpr_strContext differentiates from other interfaces.
	IsExpr_strContext()
}

type Expr_strContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_strContext() *Expr_strContext {
	var p = new(Expr_strContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_str
	return p
}

func InitEmptyExpr_strContext(p *Expr_strContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_str
}

func (*Expr_strContext) IsExpr_strContext() {}

func NewExpr_strContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_strContext {
	var p = new(Expr_strContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_str

	return p
}

func (s *Expr_strContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_strContext) STRING() antlr.TerminalNode {
	return s.GetToken(NumbatParserSTRING, 0)
}

func (s *Expr_strContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_strContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_strContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_str(s)
	}
}

func (s *Expr_strContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_str(s)
	}
}

func (p *NumbatParser) Expr_str() (localctx IExpr_strContext) {
	localctx = NewExpr_strContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NumbatParserRULE_expr_str)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(NumbatParserSTRING)
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

// IExpr_nullContext is an interface to support dynamic dispatch.
type IExpr_nullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpr_nullContext differentiates from other interfaces.
	IsExpr_nullContext()
}

type Expr_nullContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_nullContext() *Expr_nullContext {
	var p = new(Expr_nullContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_null
	return p
}

func InitEmptyExpr_nullContext(p *Expr_nullContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_null
}

func (*Expr_nullContext) IsExpr_nullContext() {}

func NewExpr_nullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_nullContext {
	var p = new(Expr_nullContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_null

	return p
}

func (s *Expr_nullContext) GetParser() antlr.Parser { return s.parser }
func (s *Expr_nullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_nullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_nullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_null(s)
	}
}

func (s *Expr_nullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_null(s)
	}
}

func (p *NumbatParser) Expr_null() (localctx IExpr_nullContext) {
	localctx = NewExpr_nullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NumbatParserRULE_expr_null)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(NumbatParserT__2)
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

// IExpr_constantContext is an interface to support dynamic dispatch.
type IExpr_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_bool() IExpr_boolContext
	Expr_num() IExpr_numContext
	Expr_hex() IExpr_hexContext
	Expr_str() IExpr_strContext
	Expr_null() IExpr_nullContext

	// IsExpr_constantContext differentiates from other interfaces.
	IsExpr_constantContext()
}

type Expr_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_constantContext() *Expr_constantContext {
	var p = new(Expr_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_constant
	return p
}

func InitEmptyExpr_constantContext(p *Expr_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_constant
}

func (*Expr_constantContext) IsExpr_constantContext() {}

func NewExpr_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_constantContext {
	var p = new(Expr_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_constant

	return p
}

func (s *Expr_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_constantContext) Expr_bool() IExpr_boolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_boolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_boolContext)
}

func (s *Expr_constantContext) Expr_num() IExpr_numContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_numContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_numContext)
}

func (s *Expr_constantContext) Expr_hex() IExpr_hexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_hexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_hexContext)
}

func (s *Expr_constantContext) Expr_str() IExpr_strContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_strContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_strContext)
}

func (s *Expr_constantContext) Expr_null() IExpr_nullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_nullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_nullContext)
}

func (s *Expr_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_constant(s)
	}
}

func (s *Expr_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_constant(s)
	}
}

func (p *NumbatParser) Expr_constant() (localctx IExpr_constantContext) {
	localctx = NewExpr_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NumbatParserRULE_expr_constant)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserT__0, NumbatParserT__1:
		{
			p.SetState(96)
			p.Expr_bool()
		}

	case NumbatParserNUMBER:
		{
			p.SetState(97)
			p.Expr_num()
		}

	case NumbatParserHEX:
		{
			p.SetState(98)
			p.Expr_hex()
		}

	case NumbatParserSTRING:
		{
			p.SetState(99)
			p.Expr_str()
		}

	case NumbatParserT__2:
		{
			p.SetState(100)
			p.Expr_null()
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

// IExpr_varContext is an interface to support dynamic dispatch.
type IExpr_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode

	// IsExpr_varContext differentiates from other interfaces.
	IsExpr_varContext()
}

type Expr_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_varContext() *Expr_varContext {
	var p = new(Expr_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_var
	return p
}

func InitEmptyExpr_varContext(p *Expr_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_var
}

func (*Expr_varContext) IsExpr_varContext() {}

func NewExpr_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_varContext {
	var p = new(Expr_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_var

	return p
}

func (s *Expr_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_varContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Expr_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_var(s)
	}
}

func (s *Expr_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_var(s)
	}
}

func (p *NumbatParser) Expr_var() (localctx IExpr_varContext) {
	localctx = NewExpr_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NumbatParserRULE_expr_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(NumbatParserNON_TYPE_NAME)
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

// IExpr_callContext is an interface to support dynamic dispatch.
type IExpr_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call() ICallContext

	// IsExpr_callContext differentiates from other interfaces.
	IsExpr_callContext()
}

type Expr_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_callContext() *Expr_callContext {
	var p = new(Expr_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_call
	return p
}

func InitEmptyExpr_callContext(p *Expr_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_call
}

func (*Expr_callContext) IsExpr_callContext() {}

func NewExpr_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_callContext {
	var p = new(Expr_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_call

	return p
}

func (s *Expr_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_callContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *Expr_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_call(s)
	}
}

func (s *Expr_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_call(s)
	}
}

func (p *NumbatParser) Expr_call() (localctx IExpr_callContext) {
	localctx = NewExpr_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NumbatParserRULE_expr_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Call()
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

// IExpr_allContext is an interface to support dynamic dispatch.
type IExpr_allContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_var() IExpr_varContext
	Expr_constant() IExpr_constantContext
	Expr_call() IExpr_callContext

	// IsExpr_allContext differentiates from other interfaces.
	IsExpr_allContext()
}

type Expr_allContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_allContext() *Expr_allContext {
	var p = new(Expr_allContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_all
	return p
}

func InitEmptyExpr_allContext(p *Expr_allContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_all
}

func (*Expr_allContext) IsExpr_allContext() {}

func NewExpr_allContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_allContext {
	var p = new(Expr_allContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_all

	return p
}

func (s *Expr_allContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_allContext) Expr_var() IExpr_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_varContext)
}

func (s *Expr_allContext) Expr_constant() IExpr_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_constantContext)
}

func (s *Expr_allContext) Expr_call() IExpr_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_callContext)
}

func (s *Expr_allContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_allContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_allContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_all(s)
	}
}

func (s *Expr_allContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_all(s)
	}
}

func (p *NumbatParser) Expr_all() (localctx IExpr_allContext) {
	localctx = NewExpr_allContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NumbatParserRULE_expr_all)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserNON_TYPE_NAME:
		{
			p.SetState(107)
			p.Expr_var()
		}

	case NumbatParserT__0, NumbatParserT__1, NumbatParserT__2, NumbatParserNUMBER, NumbatParserHEX, NumbatParserSTRING:
		{
			p.SetState(108)
			p.Expr_constant()
		}

	case NumbatParserT__4:
		{
			p.SetState(109)
			p.Expr_call()
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

// IParam_exprContext is an interface to support dynamic dispatch.
type IParam_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_constant() IExpr_constantContext

	// IsParam_exprContext differentiates from other interfaces.
	IsParam_exprContext()
}

type Param_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_exprContext() *Param_exprContext {
	var p = new(Param_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_param_expr
	return p
}

func InitEmptyParam_exprContext(p *Param_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_param_expr
}

func (*Param_exprContext) IsParam_exprContext() {}

func NewParam_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_exprContext {
	var p = new(Param_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_param_expr

	return p
}

func (s *Param_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_exprContext) Expr_constant() IExpr_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_constantContext)
}

func (s *Param_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterParam_expr(s)
	}
}

func (s *Param_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitParam_expr(s)
	}
}

func (p *NumbatParser) Param_expr() (localctx IParam_exprContext) {
	localctx = NewParam_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NumbatParserRULE_param_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(NumbatParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Expr_constant()
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

// IParam_typeContext is an interface to support dynamic dispatch.
type IParam_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsParam_typeContext differentiates from other interfaces.
	IsParam_typeContext()
}

type Param_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_typeContext() *Param_typeContext {
	var p = new(Param_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_param_type
	return p
}

func InitEmptyParam_typeContext(p *Param_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_param_type
}

func (*Param_typeContext) IsParam_typeContext() {}

func NewParam_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_typeContext {
	var p = new(Param_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_param_type

	return p
}

func (s *Param_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_typeContext) Type_() ITypeContext {
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

func (s *Param_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterParam_type(s)
	}
}

func (s *Param_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitParam_type(s)
	}
}

func (p *NumbatParser) Param_type() (localctx IParam_typeContext) {
	localctx = NewParam_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NumbatParserRULE_param_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
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

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode
	Param_type() IParam_typeContext
	Param_expr() IParam_exprContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *ParamContext) Param_type() IParam_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_typeContext)
}

func (s *ParamContext) Param_expr() IParam_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_exprContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *NumbatParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NumbatParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(NumbatParserNON_TYPE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__4 || _la == NumbatParserTYPE_NAME {
		{
			p.SetState(118)
			p.Param_type()
		}

	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__3 {
		{
			p.SetState(121)
			p.Param_expr()
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

// IType_outContext is an interface to support dynamic dispatch.
type IType_outContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_NAME() antlr.TerminalNode

	// IsType_outContext differentiates from other interfaces.
	IsType_outContext()
}

type Type_outContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_outContext() *Type_outContext {
	var p = new(Type_outContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_out
	return p
}

func InitEmptyType_outContext(p *Type_outContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_out
}

func (*Type_outContext) IsType_outContext() {}

func NewType_outContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_outContext {
	var p = new(Type_outContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_type_out

	return p
}

func (s *Type_outContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_outContext) TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserTYPE_NAME, 0)
}

func (s *Type_outContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_outContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_outContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterType_out(s)
	}
}

func (s *Type_outContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitType_out(s)
	}
}

func (p *NumbatParser) Type_out() (localctx IType_outContext) {
	localctx = NewType_outContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NumbatParserRULE_type_out)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(NumbatParserTYPE_NAME)
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

// IType_inContext is an interface to support dynamic dispatch.
type IType_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsType_inContext differentiates from other interfaces.
	IsType_inContext()
}

type Type_inContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_inContext() *Type_inContext {
	var p = new(Type_inContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_in
	return p
}

func InitEmptyType_inContext(p *Type_inContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_in
}

func (*Type_inContext) IsType_inContext() {}

func NewType_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_inContext {
	var p = new(Type_inContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_type_in

	return p
}

func (s *Type_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_inContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *Type_inContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
}

func (s *Type_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterType_in(s)
	}
}

func (s *Type_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitType_in(s)
	}
}

func (p *NumbatParser) Type_in() (localctx IType_inContext) {
	localctx = NewType_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NumbatParserRULE_type_in)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(NumbatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Param()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NumbatParserT__5 {
		{
			p.SetState(128)
			p.Match(NumbatParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Param()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
		p.Match(NumbatParserT__6)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_out() IType_outContext
	Type_in() IType_inContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Type_out() IType_outContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_outContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_outContext)
}

func (s *TypeContext) Type_in() IType_inContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_inContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_inContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *NumbatParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NumbatParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NumbatParserT__4 {
			{
				p.SetState(137)
				p.Type_in()
			}

		}
		{
			p.SetState(140)
			p.Type_out()
		}

	case 2:
		{
			p.SetState(141)
			p.Type_in()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NumbatParserTYPE_NAME {
			{
				p.SetState(142)
				p.Type_out()
			}

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

// IProc_bodyContext is an interface to support dynamic dispatch.
type IProc_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProc_bodyContext differentiates from other interfaces.
	IsProc_bodyContext()
}

type Proc_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProc_bodyContext() *Proc_bodyContext {
	var p = new(Proc_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_body
	return p
}

func InitEmptyProc_bodyContext(p *Proc_bodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_body
}

func (*Proc_bodyContext) IsProc_bodyContext() {}

func NewProc_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Proc_bodyContext {
	var p = new(Proc_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_proc_body

	return p
}

func (s *Proc_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Proc_bodyContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(NumbatParserWS)
}

func (s *Proc_bodyContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(NumbatParserWS, i)
}

func (s *Proc_bodyContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NumbatParserNEWLINE)
}

func (s *Proc_bodyContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NumbatParserNEWLINE, i)
}

func (s *Proc_bodyContext) AllStatement() []IStatementContext {
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

func (s *Proc_bodyContext) Statement(i int) IStatementContext {
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

func (s *Proc_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Proc_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Proc_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProc_body(s)
	}
}

func (s *Proc_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProc_body(s)
	}
}

func (p *NumbatParser) Proc_body() (localctx IProc_bodyContext) {
	localctx = NewProc_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NumbatParserRULE_proc_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(NumbatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3190816) != 0 {
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NumbatParserWS:
			{
				p.SetState(148)
				p.Match(NumbatParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserNEWLINE:
			{
				p.SetState(149)
				p.Match(NumbatParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserT__4, NumbatParserT__11, NumbatParserT__12, NumbatParserNON_TYPE_NAME:
			{
				p.SetState(150)
				p.Statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(NumbatParserT__8)
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

// IProc_nameContext is an interface to support dynamic dispatch.
type IProc_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode

	// IsProc_nameContext differentiates from other interfaces.
	IsProc_nameContext()
}

type Proc_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProc_nameContext() *Proc_nameContext {
	var p = new(Proc_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_name
	return p
}

func InitEmptyProc_nameContext(p *Proc_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_name
}

func (*Proc_nameContext) IsProc_nameContext() {}

func NewProc_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Proc_nameContext {
	var p = new(Proc_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_proc_name

	return p
}

func (s *Proc_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Proc_nameContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Proc_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Proc_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Proc_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProc_name(s)
	}
}

func (s *Proc_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProc_name(s)
	}
}

func (p *NumbatParser) Proc_name() (localctx IProc_nameContext) {
	localctx = NewProc_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NumbatParserRULE_proc_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(NumbatParserNON_TYPE_NAME)
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

// IProc_typeContext is an interface to support dynamic dispatch.
type IProc_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsProc_typeContext differentiates from other interfaces.
	IsProc_typeContext()
}

type Proc_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProc_typeContext() *Proc_typeContext {
	var p = new(Proc_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_type
	return p
}

func InitEmptyProc_typeContext(p *Proc_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_type
}

func (*Proc_typeContext) IsProc_typeContext() {}

func NewProc_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Proc_typeContext {
	var p = new(Proc_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_proc_type

	return p
}

func (s *Proc_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Proc_typeContext) Type_() ITypeContext {
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

func (s *Proc_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Proc_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Proc_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProc_type(s)
	}
}

func (s *Proc_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProc_type(s)
	}
}

func (p *NumbatParser) Proc_type() (localctx IProc_typeContext) {
	localctx = NewProc_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NumbatParserRULE_proc_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
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

// IProc_defContext is an interface to support dynamic dispatch.
type IProc_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Proc_name() IProc_nameContext
	Proc_type() IProc_typeContext

	// IsProc_defContext differentiates from other interfaces.
	IsProc_defContext()
}

type Proc_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProc_defContext() *Proc_defContext {
	var p = new(Proc_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_def
	return p
}

func InitEmptyProc_defContext(p *Proc_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc_def
}

func (*Proc_defContext) IsProc_defContext() {}

func NewProc_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Proc_defContext {
	var p = new(Proc_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_proc_def

	return p
}

func (s *Proc_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Proc_defContext) Proc_name() IProc_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProc_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProc_nameContext)
}

func (s *Proc_defContext) Proc_type() IProc_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProc_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProc_typeContext)
}

func (s *Proc_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Proc_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Proc_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProc_def(s)
	}
}

func (s *Proc_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProc_def(s)
	}
}

func (p *NumbatParser) Proc_def() (localctx IProc_defContext) {
	localctx = NewProc_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NumbatParserRULE_proc_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(NumbatParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Proc_name()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__4 || _la == NumbatParserTYPE_NAME {
		{
			p.SetState(164)
			p.Proc_type()
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

// IProcContext is an interface to support dynamic dispatch.
type IProcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Proc_def() IProc_defContext
	Proc_body() IProc_bodyContext

	// IsProcContext differentiates from other interfaces.
	IsProcContext()
}

type ProcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcContext() *ProcContext {
	var p = new(ProcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc
	return p
}

func InitEmptyProcContext(p *ProcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_proc
}

func (*ProcContext) IsProcContext() {}

func NewProcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcContext {
	var p = new(ProcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_proc

	return p
}

func (s *ProcContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcContext) Proc_def() IProc_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProc_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProc_defContext)
}

func (s *ProcContext) Proc_body() IProc_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProc_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProc_bodyContext)
}

func (s *ProcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProc(s)
	}
}

func (s *ProcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProc(s)
	}
}

func (p *NumbatParser) Proc() (localctx IProcContext) {
	localctx = NewProcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NumbatParserRULE_proc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Proc_def()
	}
	{
		p.SetState(168)
		p.Proc_body()
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

// ICall_exprContext is an interface to support dynamic dispatch.
type ICall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr_var() []IExpr_varContext
	Expr_var(i int) IExpr_varContext
	AllExpr_constant() []IExpr_constantContext
	Expr_constant(i int) IExpr_constantContext

	// IsCall_exprContext differentiates from other interfaces.
	IsCall_exprContext()
}

type Call_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_exprContext() *Call_exprContext {
	var p = new(Call_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_expr
	return p
}

func InitEmptyCall_exprContext(p *Call_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_expr
}

func (*Call_exprContext) IsCall_exprContext() {}

func NewCall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_exprContext {
	var p = new(Call_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call_expr

	return p
}

func (s *Call_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_exprContext) AllExpr_var() []IExpr_varContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpr_varContext); ok {
			len++
		}
	}

	tst := make([]IExpr_varContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpr_varContext); ok {
			tst[i] = t.(IExpr_varContext)
			i++
		}
	}

	return tst
}

func (s *Call_exprContext) Expr_var(i int) IExpr_varContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_varContext); ok {
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

	return t.(IExpr_varContext)
}

func (s *Call_exprContext) AllExpr_constant() []IExpr_constantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpr_constantContext); ok {
			len++
		}
	}

	tst := make([]IExpr_constantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpr_constantContext); ok {
			tst[i] = t.(IExpr_constantContext)
			i++
		}
	}

	return tst
}

func (s *Call_exprContext) Expr_constant(i int) IExpr_constantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_constantContext); ok {
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

	return t.(IExpr_constantContext)
}

func (s *Call_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall_expr(s)
	}
}

func (s *Call_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall_expr(s)
	}
}

func (p *NumbatParser) Call_expr() (localctx ICall_exprContext) {
	localctx = NewCall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NumbatParserRULE_call_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(NumbatParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1507342) != 0) {
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NumbatParserNON_TYPE_NAME:
			{
				p.SetState(171)
				p.Expr_var()
			}

		case NumbatParserT__0, NumbatParserT__1, NumbatParserT__2, NumbatParserNUMBER, NumbatParserHEX, NumbatParserSTRING:
			{
				p.SetState(172)
				p.Expr_constant()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(175)
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

// ICall_secondaryContext is an interface to support dynamic dispatch.
type ICall_secondaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode

	// IsCall_secondaryContext differentiates from other interfaces.
	IsCall_secondaryContext()
}

type Call_secondaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_secondaryContext() *Call_secondaryContext {
	var p = new(Call_secondaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_secondary
	return p
}

func InitEmptyCall_secondaryContext(p *Call_secondaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_secondary
}

func (*Call_secondaryContext) IsCall_secondaryContext() {}

func NewCall_secondaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_secondaryContext {
	var p = new(Call_secondaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call_secondary

	return p
}

func (s *Call_secondaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_secondaryContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Call_secondaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_secondaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_secondaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall_secondary(s)
	}
}

func (s *Call_secondaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall_secondary(s)
	}
}

func (p *NumbatParser) Call_secondary() (localctx ICall_secondaryContext) {
	localctx = NewCall_secondaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, NumbatParserRULE_call_secondary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(NumbatParserNON_TYPE_NAME)
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

// ICall_primaryContext is an interface to support dynamic dispatch.
type ICall_primaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_NAME() antlr.TerminalNode
	NON_TYPE_NAME() antlr.TerminalNode

	// IsCall_primaryContext differentiates from other interfaces.
	IsCall_primaryContext()
}

type Call_primaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_primaryContext() *Call_primaryContext {
	var p = new(Call_primaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_primary
	return p
}

func InitEmptyCall_primaryContext(p *Call_primaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_primary
}

func (*Call_primaryContext) IsCall_primaryContext() {}

func NewCall_primaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_primaryContext {
	var p = new(Call_primaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call_primary

	return p
}

func (s *Call_primaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_primaryContext) TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserTYPE_NAME, 0)
}

func (s *Call_primaryContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Call_primaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_primaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_primaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall_primary(s)
	}
}

func (s *Call_primaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall_primary(s)
	}
}

func (p *NumbatParser) Call_primary() (localctx ICall_primaryContext) {
	localctx = NewCall_primaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NumbatParserRULE_call_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NumbatParserTYPE_NAME || _la == NumbatParserNON_TYPE_NAME) {
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

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call_primary() ICall_primaryContext
	Call_secondary() ICall_secondaryContext
	Call_expr() ICall_exprContext

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) Call_primary() ICall_primaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_primaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_primaryContext)
}

func (s *CallContext) Call_secondary() ICall_secondaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_secondaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_secondaryContext)
}

func (s *CallContext) Call_expr() ICall_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *NumbatParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NumbatParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(NumbatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Call_primary()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserNON_TYPE_NAME {
		{
			p.SetState(183)
			p.Call_secondary()
		}

	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__10 {
		{
			p.SetState(186)
			p.Call_expr()
		}

	}
	{
		p.SetState(189)
		p.Match(NumbatParserT__6)
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

// ICall_stmtContext is an interface to support dynamic dispatch.
type ICall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call() ICallContext

	// IsCall_stmtContext differentiates from other interfaces.
	IsCall_stmtContext()
}

type Call_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_stmtContext() *Call_stmtContext {
	var p = new(Call_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_stmt
	return p
}

func InitEmptyCall_stmtContext(p *Call_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_stmt
}

func (*Call_stmtContext) IsCall_stmtContext() {}

func NewCall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_stmtContext {
	var p = new(Call_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call_stmt

	return p
}

func (s *Call_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_stmtContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *Call_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall_stmt(s)
	}
}

func (s *Call_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall_stmt(s)
	}
}

func (p *NumbatParser) Call_stmt() (localctx ICall_stmtContext) {
	localctx = NewCall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NumbatParserRULE_call_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Call()
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

// ILet_exprContext is an interface to support dynamic dispatch.
type ILet_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_all() IExpr_allContext

	// IsLet_exprContext differentiates from other interfaces.
	IsLet_exprContext()
}

type Let_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_exprContext() *Let_exprContext {
	var p = new(Let_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let_expr
	return p
}

func InitEmptyLet_exprContext(p *Let_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let_expr
}

func (*Let_exprContext) IsLet_exprContext() {}

func NewLet_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_exprContext {
	var p = new(Let_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_let_expr

	return p
}

func (s *Let_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_exprContext) Expr_all() IExpr_allContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_allContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_allContext)
}

func (s *Let_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterLet_expr(s)
	}
}

func (s *Let_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitLet_expr(s)
	}
}

func (p *NumbatParser) Let_expr() (localctx ILet_exprContext) {
	localctx = NewLet_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NumbatParserRULE_let_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(NumbatParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Expr_all()
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

// ILet_var_typeContext is an interface to support dynamic dispatch.
type ILet_var_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsLet_var_typeContext differentiates from other interfaces.
	IsLet_var_typeContext()
}

type Let_var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_var_typeContext() *Let_var_typeContext {
	var p = new(Let_var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let_var_type
	return p
}

func InitEmptyLet_var_typeContext(p *Let_var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let_var_type
}

func (*Let_var_typeContext) IsLet_var_typeContext() {}

func NewLet_var_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_var_typeContext {
	var p = new(Let_var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_let_var_type

	return p
}

func (s *Let_var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_var_typeContext) Type_() ITypeContext {
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

func (s *Let_var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterLet_var_type(s)
	}
}

func (s *Let_var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitLet_var_type(s)
	}
}

func (p *NumbatParser) Let_var_type() (localctx ILet_var_typeContext) {
	localctx = NewLet_var_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NumbatParserRULE_let_var_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
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

// ILet_var_nameContext is an interface to support dynamic dispatch.
type ILet_var_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode

	// IsLet_var_nameContext differentiates from other interfaces.
	IsLet_var_nameContext()
}

type Let_var_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_var_nameContext() *Let_var_nameContext {
	var p = new(Let_var_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let_var_name
	return p
}

func InitEmptyLet_var_nameContext(p *Let_var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let_var_name
}

func (*Let_var_nameContext) IsLet_var_nameContext() {}

func NewLet_var_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_var_nameContext {
	var p = new(Let_var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_let_var_name

	return p
}

func (s *Let_var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_var_nameContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Let_var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterLet_var_name(s)
	}
}

func (s *Let_var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitLet_var_name(s)
	}
}

func (p *NumbatParser) Let_var_name() (localctx ILet_var_nameContext) {
	localctx = NewLet_var_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NumbatParserRULE_let_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(NumbatParserNON_TYPE_NAME)
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

// ILetContext is an interface to support dynamic dispatch.
type ILetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Let_var_name() ILet_var_nameContext
	Let_var_type() ILet_var_typeContext
	Let_expr() ILet_exprContext

	// IsLetContext differentiates from other interfaces.
	IsLetContext()
}

type LetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetContext() *LetContext {
	var p = new(LetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let
	return p
}

func InitEmptyLetContext(p *LetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_let
}

func (*LetContext) IsLetContext() {}

func NewLetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetContext {
	var p = new(LetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_let

	return p
}

func (s *LetContext) GetParser() antlr.Parser { return s.parser }

func (s *LetContext) Let_var_name() ILet_var_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILet_var_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILet_var_nameContext)
}

func (s *LetContext) Let_var_type() ILet_var_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILet_var_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILet_var_typeContext)
}

func (s *LetContext) Let_expr() ILet_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILet_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILet_exprContext)
}

func (s *LetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterLet(s)
	}
}

func (s *LetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitLet(s)
	}
}

func (p *NumbatParser) Let() (localctx ILetContext) {
	localctx = NewLetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NumbatParserRULE_let)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(NumbatParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Let_var_name()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(202)
			p.Let_var_type()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__3 {
		{
			p.SetState(205)
			p.Let_expr()
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

// IAssignment_exprContext is an interface to support dynamic dispatch.
type IAssignment_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr_all() []IExpr_allContext
	Expr_all(i int) IExpr_allContext

	// IsAssignment_exprContext differentiates from other interfaces.
	IsAssignment_exprContext()
}

type Assignment_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_exprContext() *Assignment_exprContext {
	var p = new(Assignment_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_assignment_expr
	return p
}

func InitEmptyAssignment_exprContext(p *Assignment_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_assignment_expr
}

func (*Assignment_exprContext) IsAssignment_exprContext() {}

func NewAssignment_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_exprContext {
	var p = new(Assignment_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_assignment_expr

	return p
}

func (s *Assignment_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_exprContext) AllExpr_all() []IExpr_allContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpr_allContext); ok {
			len++
		}
	}

	tst := make([]IExpr_allContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpr_allContext); ok {
			tst[i] = t.(IExpr_allContext)
			i++
		}
	}

	return tst
}

func (s *Assignment_exprContext) Expr_all(i int) IExpr_allContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_allContext); ok {
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

	return t.(IExpr_allContext)
}

func (s *Assignment_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterAssignment_expr(s)
	}
}

func (s *Assignment_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitAssignment_expr(s)
	}
}

func (p *NumbatParser) Assignment_expr() (localctx IAssignment_exprContext) {
	localctx = NewAssignment_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NumbatParserRULE_assignment_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(NumbatParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Expr_all()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__5 {
		{
			p.SetState(210)
			p.Match(NumbatParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Expr_all()
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

// IAssignment_varContext is an interface to support dynamic dispatch.
type IAssignment_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode

	// IsAssignment_varContext differentiates from other interfaces.
	IsAssignment_varContext()
}

type Assignment_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_varContext() *Assignment_varContext {
	var p = new(Assignment_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_assignment_var
	return p
}

func InitEmptyAssignment_varContext(p *Assignment_varContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_assignment_var
}

func (*Assignment_varContext) IsAssignment_varContext() {}

func NewAssignment_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_varContext {
	var p = new(Assignment_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_assignment_var

	return p
}

func (s *Assignment_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_varContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Assignment_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterAssignment_var(s)
	}
}

func (s *Assignment_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitAssignment_var(s)
	}
}

func (p *NumbatParser) Assignment_var() (localctx IAssignment_varContext) {
	localctx = NewAssignment_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NumbatParserRULE_assignment_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(NumbatParserNON_TYPE_NAME)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignment_var() []IAssignment_varContext
	Assignment_var(i int) IAssignment_varContext
	Assignment_expr() IAssignment_exprContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllAssignment_var() []IAssignment_varContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignment_varContext); ok {
			len++
		}
	}

	tst := make([]IAssignment_varContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignment_varContext); ok {
			tst[i] = t.(IAssignment_varContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) Assignment_var(i int) IAssignment_varContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_varContext); ok {
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

	return t.(IAssignment_varContext)
}

func (s *AssignmentContext) Assignment_expr() IAssignment_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_exprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *NumbatParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NumbatParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Assignment_var()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__5 {
		{
			p.SetState(217)
			p.Match(NumbatParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Assignment_var()
		}

	}
	{
		p.SetState(221)
		p.Assignment_expr()
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

// IReturn_exprContext is an interface to support dynamic dispatch.
type IReturn_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_all() IExpr_allContext

	// IsReturn_exprContext differentiates from other interfaces.
	IsReturn_exprContext()
}

type Return_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_exprContext() *Return_exprContext {
	var p = new(Return_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_return_expr
	return p
}

func InitEmptyReturn_exprContext(p *Return_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_return_expr
}

func (*Return_exprContext) IsReturn_exprContext() {}

func NewReturn_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_exprContext {
	var p = new(Return_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_return_expr

	return p
}

func (s *Return_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_exprContext) Expr_all() IExpr_allContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_allContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_allContext)
}

func (s *Return_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterReturn_expr(s)
	}
}

func (s *Return_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitReturn_expr(s)
	}
}

func (p *NumbatParser) Return_expr() (localctx IReturn_exprContext) {
	localctx = NewReturn_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, NumbatParserRULE_return_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Expr_all()
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

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Return_expr() IReturn_exprContext

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_return
	return p
}

func InitEmptyReturnContext(p *ReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_return
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) Return_expr() IReturn_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_exprContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *NumbatParser) Return_() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, NumbatParserRULE_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(NumbatParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(226)
			p.Return_expr()
		}

	} else if p.HasError() { // JIM
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

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Return_() IReturnContext

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_return_stmt
	return p
}

func InitEmptyReturn_stmtContext(p *Return_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_return_stmt
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterReturn_stmt(s)
	}
}

func (s *Return_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitReturn_stmt(s)
	}
}

func (p *NumbatParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NumbatParserRULE_return_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Return_()
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

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Proc_body() IProc_bodyContext

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
	p.RuleIndex = NumbatParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Proc_body() IProc_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProc_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProc_bodyContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *NumbatParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, NumbatParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(NumbatParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Proc_body()
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
	Call_stmt() ICall_stmtContext
	Let() ILetContext
	Return_stmt() IReturn_stmtContext
	Assignment() IAssignmentContext

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
	p.RuleIndex = NumbatParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Call_stmt() ICall_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_stmtContext)
}

func (s *StatementContext) Let() ILetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetContext)
}

func (s *StatementContext) Return_stmt() IReturn_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *NumbatParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, NumbatParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserT__4:
		{
			p.SetState(234)
			p.Call_stmt()
		}

	case NumbatParserT__11:
		{
			p.SetState(235)
			p.Let()
		}

	case NumbatParserT__12:
		{
			p.SetState(236)
			p.Return_stmt()
		}

	case NumbatParserNON_TYPE_NAME:
		{
			p.SetState(237)
			p.Assignment()
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

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Program() IProgramContext
	Proc() IProcContext

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_object
	return p
}

func InitEmptyObjectContext(p *ObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_object
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) Program() IProgramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *ObjectContext) Proc() IProcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *NumbatParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, NumbatParserRULE_object)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserT__13:
		{
			p.SetState(240)
			p.Program()
		}

	case NumbatParserT__9:
		{
			p.SetState(241)
			p.Proc()
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