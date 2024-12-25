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
		"", "'true'", "'false'", "'null'", "'['", "']'", "'='", "'('", "')'",
		"','", "'do'", "'end'", "'proc'", "':'", "'var'", "'return'", "'program'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"UNIT", "NEWLINE", "NUMBER", "HEX", "STRING", "TYPE_NAME", "NON_TYPE_NAME",
		"WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"prog", "expr_bool", "expr_num", "expr_hex", "expr_str", "expr_null",
		"expr_constant", "expr_var", "expr_call", "expr_seq", "expr_all", "param_expr",
		"param_type", "param", "type_super_atomic_seq", "type_super_atomic",
		"type_out", "type_in", "type", "proc_body", "proc_name", "proc_type",
		"proc_def", "proc", "call_arg", "call_args", "call_secondary", "call_primary",
		"call", "call_stmt", "var_expr", "var_type", "var_name", "var_stmt",
		"assignment_expr", "assignment_var", "assignment", "return_expr", "return",
		"return_stmt", "program", "statement", "object",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 285, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 5, 0, 90, 8, 0, 10, 0, 12, 0, 93, 9, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 3, 2, 99, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 112, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 122, 8, 9, 10, 9, 12, 9, 125, 9, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 133, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 143, 8, 13, 1, 14, 1, 14, 3,
		14, 147, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 153, 8, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 163, 8, 17, 10, 17,
		12, 17, 166, 9, 17, 1, 17, 1, 17, 3, 17, 170, 8, 17, 1, 18, 3, 18, 173,
		8, 18, 1, 18, 1, 18, 1, 18, 3, 18, 178, 8, 18, 3, 18, 180, 8, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 5, 19, 186, 8, 19, 10, 19, 12, 19, 189, 9, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 3, 22, 200,
		8, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 207, 8, 24, 1, 25, 1,
		25, 1, 25, 5, 25, 212, 8, 25, 10, 25, 12, 25, 215, 9, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 224, 8, 28, 1, 28, 1, 28, 3,
		28, 228, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 244, 8, 33, 1, 33, 3,
		33, 247, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 253, 8, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 3, 36, 260, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 38, 1, 38, 3, 38, 268, 8, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 41, 3, 41, 279, 8, 41, 1, 42, 1, 42, 3, 42, 283, 8,
		42, 1, 42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 2, 1, 0, 1, 2, 1, 0, 22,
		23, 280, 0, 91, 1, 0, 0, 0, 2, 94, 1, 0, 0, 0, 4, 96, 1, 0, 0, 0, 6, 100,
		1, 0, 0, 0, 8, 102, 1, 0, 0, 0, 10, 104, 1, 0, 0, 0, 12, 111, 1, 0, 0,
		0, 14, 113, 1, 0, 0, 0, 16, 115, 1, 0, 0, 0, 18, 117, 1, 0, 0, 0, 20, 132,
		1, 0, 0, 0, 22, 134, 1, 0, 0, 0, 24, 137, 1, 0, 0, 0, 26, 139, 1, 0, 0,
		0, 28, 144, 1, 0, 0, 0, 30, 150, 1, 0, 0, 0, 32, 154, 1, 0, 0, 0, 34, 169,
		1, 0, 0, 0, 36, 179, 1, 0, 0, 0, 38, 181, 1, 0, 0, 0, 40, 192, 1, 0, 0,
		0, 42, 194, 1, 0, 0, 0, 44, 196, 1, 0, 0, 0, 46, 201, 1, 0, 0, 0, 48, 206,
		1, 0, 0, 0, 50, 208, 1, 0, 0, 0, 52, 216, 1, 0, 0, 0, 54, 218, 1, 0, 0,
		0, 56, 220, 1, 0, 0, 0, 58, 231, 1, 0, 0, 0, 60, 233, 1, 0, 0, 0, 62, 236,
		1, 0, 0, 0, 64, 238, 1, 0, 0, 0, 66, 240, 1, 0, 0, 0, 68, 248, 1, 0, 0,
		0, 70, 254, 1, 0, 0, 0, 72, 256, 1, 0, 0, 0, 74, 263, 1, 0, 0, 0, 76, 265,
		1, 0, 0, 0, 78, 269, 1, 0, 0, 0, 80, 271, 1, 0, 0, 0, 82, 278, 1, 0, 0,
		0, 84, 282, 1, 0, 0, 0, 86, 90, 5, 24, 0, 0, 87, 90, 5, 18, 0, 0, 88, 90,
		3, 84, 42, 0, 89, 86, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0,
		0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 1, 1,
		0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 7, 0, 0, 0, 95, 3, 1, 0, 0, 0, 96,
		98, 5, 19, 0, 0, 97, 99, 5, 17, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 99, 5, 1, 0, 0, 0, 100, 101, 5, 20, 0, 0, 101, 7, 1, 0, 0, 0, 102,
		103, 5, 21, 0, 0, 103, 9, 1, 0, 0, 0, 104, 105, 5, 3, 0, 0, 105, 11, 1,
		0, 0, 0, 106, 112, 3, 2, 1, 0, 107, 112, 3, 4, 2, 0, 108, 112, 3, 6, 3,
		0, 109, 112, 3, 8, 4, 0, 110, 112, 3, 10, 5, 0, 111, 106, 1, 0, 0, 0, 111,
		107, 1, 0, 0, 0, 111, 108, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 110,
		1, 0, 0, 0, 112, 13, 1, 0, 0, 0, 113, 114, 5, 23, 0, 0, 114, 15, 1, 0,
		0, 0, 115, 116, 3, 56, 28, 0, 116, 17, 1, 0, 0, 0, 117, 123, 5, 4, 0, 0,
		118, 122, 3, 14, 7, 0, 119, 122, 3, 12, 6, 0, 120, 122, 3, 16, 8, 0, 121,
		118, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 125,
		1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0,
		0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 5, 0, 0, 127, 19, 1, 0, 0, 0,
		128, 133, 3, 14, 7, 0, 129, 133, 3, 12, 6, 0, 130, 133, 3, 16, 8, 0, 131,
		133, 3, 18, 9, 0, 132, 128, 1, 0, 0, 0, 132, 129, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 21, 1, 0, 0, 0, 134, 135, 5, 6,
		0, 0, 135, 136, 3, 12, 6, 0, 136, 23, 1, 0, 0, 0, 137, 138, 3, 36, 18,
		0, 138, 25, 1, 0, 0, 0, 139, 140, 5, 23, 0, 0, 140, 142, 3, 24, 12, 0,
		141, 143, 3, 22, 11, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143,
		27, 1, 0, 0, 0, 144, 146, 5, 4, 0, 0, 145, 147, 5, 19, 0, 0, 146, 145,
		1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 5, 5,
		0, 0, 149, 29, 1, 0, 0, 0, 150, 152, 5, 22, 0, 0, 151, 153, 3, 28, 14,
		0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 31, 1, 0, 0, 0, 154,
		155, 3, 30, 15, 0, 155, 33, 1, 0, 0, 0, 156, 157, 5, 7, 0, 0, 157, 170,
		5, 8, 0, 0, 158, 159, 5, 7, 0, 0, 159, 164, 3, 26, 13, 0, 160, 161, 5,
		9, 0, 0, 161, 163, 3, 26, 13, 0, 162, 160, 1, 0, 0, 0, 163, 166, 1, 0,
		0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 167, 1, 0, 0, 0,
		166, 164, 1, 0, 0, 0, 167, 168, 5, 8, 0, 0, 168, 170, 1, 0, 0, 0, 169,
		156, 1, 0, 0, 0, 169, 158, 1, 0, 0, 0, 170, 35, 1, 0, 0, 0, 171, 173, 3,
		34, 17, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 180, 3, 32, 16, 0, 175, 177, 3, 34, 17, 0, 176, 178, 3, 32,
		16, 0, 177, 176, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0,
		179, 172, 1, 0, 0, 0, 179, 175, 1, 0, 0, 0, 180, 37, 1, 0, 0, 0, 181, 187,
		5, 10, 0, 0, 182, 186, 5, 24, 0, 0, 183, 186, 5, 18, 0, 0, 184, 186, 3,
		82, 41, 0, 185, 182, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0,
		0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 190, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 11, 0, 0, 191,
		39, 1, 0, 0, 0, 192, 193, 5, 23, 0, 0, 193, 41, 1, 0, 0, 0, 194, 195, 3,
		36, 18, 0, 195, 43, 1, 0, 0, 0, 196, 197, 5, 12, 0, 0, 197, 199, 3, 40,
		20, 0, 198, 200, 3, 42, 21, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0,
		0, 200, 45, 1, 0, 0, 0, 201, 202, 3, 44, 22, 0, 202, 203, 3, 38, 19, 0,
		203, 47, 1, 0, 0, 0, 204, 207, 3, 14, 7, 0, 205, 207, 3, 12, 6, 0, 206,
		204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 49, 1, 0, 0, 0, 208, 213, 3,
		20, 10, 0, 209, 210, 5, 9, 0, 0, 210, 212, 3, 20, 10, 0, 211, 209, 1, 0,
		0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0,
		214, 51, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 23, 0, 0, 217,
		53, 1, 0, 0, 0, 218, 219, 7, 1, 0, 0, 219, 55, 1, 0, 0, 0, 220, 223, 3,
		54, 27, 0, 221, 222, 5, 13, 0, 0, 222, 224, 3, 52, 26, 0, 223, 221, 1,
		0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 227, 5, 7, 0,
		0, 226, 228, 3, 50, 25, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0,
		228, 229, 1, 0, 0, 0, 229, 230, 5, 8, 0, 0, 230, 57, 1, 0, 0, 0, 231, 232,
		3, 56, 28, 0, 232, 59, 1, 0, 0, 0, 233, 234, 5, 6, 0, 0, 234, 235, 3, 20,
		10, 0, 235, 61, 1, 0, 0, 0, 236, 237, 3, 36, 18, 0, 237, 63, 1, 0, 0, 0,
		238, 239, 5, 23, 0, 0, 239, 65, 1, 0, 0, 0, 240, 241, 5, 14, 0, 0, 241,
		243, 3, 64, 32, 0, 242, 244, 3, 62, 31, 0, 243, 242, 1, 0, 0, 0, 243, 244,
		1, 0, 0, 0, 244, 246, 1, 0, 0, 0, 245, 247, 3, 60, 30, 0, 246, 245, 1,
		0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 67, 1, 0, 0, 0, 248, 249, 5, 6, 0,
		0, 249, 252, 3, 20, 10, 0, 250, 251, 5, 9, 0, 0, 251, 253, 3, 20, 10, 0,
		252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 69, 1, 0, 0, 0, 254, 255,
		5, 23, 0, 0, 255, 71, 1, 0, 0, 0, 256, 259, 3, 70, 35, 0, 257, 258, 5,
		9, 0, 0, 258, 260, 3, 70, 35, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0,
		0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 3, 68, 34, 0, 262, 73, 1, 0, 0, 0,
		263, 264, 3, 20, 10, 0, 264, 75, 1, 0, 0, 0, 265, 267, 5, 15, 0, 0, 266,
		268, 3, 74, 37, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 77,
		1, 0, 0, 0, 269, 270, 3, 76, 38, 0, 270, 79, 1, 0, 0, 0, 271, 272, 5, 16,
		0, 0, 272, 273, 3, 38, 19, 0, 273, 81, 1, 0, 0, 0, 274, 279, 3, 58, 29,
		0, 275, 279, 3, 66, 33, 0, 276, 279, 3, 78, 39, 0, 277, 279, 3, 72, 36,
		0, 278, 274, 1, 0, 0, 0, 278, 275, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278,
		277, 1, 0, 0, 0, 279, 83, 1, 0, 0, 0, 280, 283, 3, 80, 40, 0, 281, 283,
		3, 46, 23, 0, 282, 280, 1, 0, 0, 0, 282, 281, 1, 0, 0, 0, 283, 85, 1, 0,
		0, 0, 29, 89, 91, 98, 111, 121, 123, 132, 142, 146, 152, 164, 169, 172,
		177, 179, 185, 187, 199, 206, 213, 223, 227, 243, 246, 252, 259, 267, 278,
		282,
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
	NumbatParserT__14         = 15
	NumbatParserT__15         = 16
	NumbatParserUNIT          = 17
	NumbatParserNEWLINE       = 18
	NumbatParserNUMBER        = 19
	NumbatParserHEX           = 20
	NumbatParserSTRING        = 21
	NumbatParserTYPE_NAME     = 22
	NumbatParserNON_TYPE_NAME = 23
	NumbatParserWS            = 24
	NumbatParserCOMMENT       = 25
	NumbatParserLINE_COMMENT  = 26
)

// NumbatParser rules.
const (
	NumbatParserRULE_prog                  = 0
	NumbatParserRULE_expr_bool             = 1
	NumbatParserRULE_expr_num              = 2
	NumbatParserRULE_expr_hex              = 3
	NumbatParserRULE_expr_str              = 4
	NumbatParserRULE_expr_null             = 5
	NumbatParserRULE_expr_constant         = 6
	NumbatParserRULE_expr_var              = 7
	NumbatParserRULE_expr_call             = 8
	NumbatParserRULE_expr_seq              = 9
	NumbatParserRULE_expr_all              = 10
	NumbatParserRULE_param_expr            = 11
	NumbatParserRULE_param_type            = 12
	NumbatParserRULE_param                 = 13
	NumbatParserRULE_type_super_atomic_seq = 14
	NumbatParserRULE_type_super_atomic     = 15
	NumbatParserRULE_type_out              = 16
	NumbatParserRULE_type_in               = 17
	NumbatParserRULE_type                  = 18
	NumbatParserRULE_proc_body             = 19
	NumbatParserRULE_proc_name             = 20
	NumbatParserRULE_proc_type             = 21
	NumbatParserRULE_proc_def              = 22
	NumbatParserRULE_proc                  = 23
	NumbatParserRULE_call_arg              = 24
	NumbatParserRULE_call_args             = 25
	NumbatParserRULE_call_secondary        = 26
	NumbatParserRULE_call_primary          = 27
	NumbatParserRULE_call                  = 28
	NumbatParserRULE_call_stmt             = 29
	NumbatParserRULE_var_expr              = 30
	NumbatParserRULE_var_type              = 31
	NumbatParserRULE_var_name              = 32
	NumbatParserRULE_var_stmt              = 33
	NumbatParserRULE_assignment_expr       = 34
	NumbatParserRULE_assignment_var        = 35
	NumbatParserRULE_assignment            = 36
	NumbatParserRULE_return_expr           = 37
	NumbatParserRULE_return                = 38
	NumbatParserRULE_return_stmt           = 39
	NumbatParserRULE_program               = 40
	NumbatParserRULE_statement             = 41
	NumbatParserRULE_object                = 42
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17108992) != 0 {
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NumbatParserWS:
			{
				p.SetState(86)
				p.Match(NumbatParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserNEWLINE:
			{
				p.SetState(87)
				p.Match(NumbatParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserT__11, NumbatParserT__15:
			{
				p.SetState(88)
				p.Object()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(93)
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
		p.SetState(94)
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
	UNIT() antlr.TerminalNode

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

func (s *Expr_numContext) UNIT() antlr.TerminalNode {
	return s.GetToken(NumbatParserUNIT, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(NumbatParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserUNIT {
		{
			p.SetState(97)
			p.Match(NumbatParserUNIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.SetState(100)
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
		p.SetState(102)
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
		p.SetState(104)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserT__0, NumbatParserT__1:
		{
			p.SetState(106)
			p.Expr_bool()
		}

	case NumbatParserNUMBER:
		{
			p.SetState(107)
			p.Expr_num()
		}

	case NumbatParserHEX:
		{
			p.SetState(108)
			p.Expr_hex()
		}

	case NumbatParserSTRING:
		{
			p.SetState(109)
			p.Expr_str()
		}

	case NumbatParserT__2:
		{
			p.SetState(110)
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
		p.SetState(113)
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
		p.SetState(115)
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

// IExpr_seqContext is an interface to support dynamic dispatch.
type IExpr_seqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr_var() []IExpr_varContext
	Expr_var(i int) IExpr_varContext
	AllExpr_constant() []IExpr_constantContext
	Expr_constant(i int) IExpr_constantContext
	AllExpr_call() []IExpr_callContext
	Expr_call(i int) IExpr_callContext

	// IsExpr_seqContext differentiates from other interfaces.
	IsExpr_seqContext()
}

type Expr_seqContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_seqContext() *Expr_seqContext {
	var p = new(Expr_seqContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_seq
	return p
}

func InitEmptyExpr_seqContext(p *Expr_seqContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_expr_seq
}

func (*Expr_seqContext) IsExpr_seqContext() {}

func NewExpr_seqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_seqContext {
	var p = new(Expr_seqContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_expr_seq

	return p
}

func (s *Expr_seqContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_seqContext) AllExpr_var() []IExpr_varContext {
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

func (s *Expr_seqContext) Expr_var(i int) IExpr_varContext {
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

func (s *Expr_seqContext) AllExpr_constant() []IExpr_constantContext {
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

func (s *Expr_seqContext) Expr_constant(i int) IExpr_constantContext {
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

func (s *Expr_seqContext) AllExpr_call() []IExpr_callContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpr_callContext); ok {
			len++
		}
	}

	tst := make([]IExpr_callContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpr_callContext); ok {
			tst[i] = t.(IExpr_callContext)
			i++
		}
	}

	return tst
}

func (s *Expr_seqContext) Expr_call(i int) IExpr_callContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_callContext); ok {
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

	return t.(IExpr_callContext)
}

func (s *Expr_seqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_seqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_seqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterExpr_seq(s)
	}
}

func (s *Expr_seqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitExpr_seq(s)
	}
}

func (p *NumbatParser) Expr_seq() (localctx IExpr_seqContext) {
	localctx = NewExpr_seqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NumbatParserRULE_expr_seq)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(NumbatParserT__3)
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
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16252942) != 0 {
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(118)
				p.Expr_var()
			}

		case 2:
			{
				p.SetState(119)
				p.Expr_constant()
			}

		case 3:
			{
				p.SetState(120)
				p.Expr_call()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(NumbatParserT__4)
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

// IExpr_allContext is an interface to support dynamic dispatch.
type IExpr_allContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_var() IExpr_varContext
	Expr_constant() IExpr_constantContext
	Expr_call() IExpr_callContext
	Expr_seq() IExpr_seqContext

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

func (s *Expr_allContext) Expr_seq() IExpr_seqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_seqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_seqContext)
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
	p.EnterRule(localctx, 20, NumbatParserRULE_expr_all)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(128)
			p.Expr_var()
		}

	case 2:
		{
			p.SetState(129)
			p.Expr_constant()
		}

	case 3:
		{
			p.SetState(130)
			p.Expr_call()
		}

	case 4:
		{
			p.SetState(131)
			p.Expr_seq()
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
	p.EnterRule(localctx, 22, NumbatParserRULE_param_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(NumbatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
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
	p.EnterRule(localctx, 24, NumbatParserRULE_param_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 26, NumbatParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(NumbatParserNON_TYPE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Param_type()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__5 {
		{
			p.SetState(141)
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

// IType_super_atomic_seqContext is an interface to support dynamic dispatch.
type IType_super_atomic_seqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsType_super_atomic_seqContext differentiates from other interfaces.
	IsType_super_atomic_seqContext()
}

type Type_super_atomic_seqContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_super_atomic_seqContext() *Type_super_atomic_seqContext {
	var p = new(Type_super_atomic_seqContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_super_atomic_seq
	return p
}

func InitEmptyType_super_atomic_seqContext(p *Type_super_atomic_seqContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_super_atomic_seq
}

func (*Type_super_atomic_seqContext) IsType_super_atomic_seqContext() {}

func NewType_super_atomic_seqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_super_atomic_seqContext {
	var p = new(Type_super_atomic_seqContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_type_super_atomic_seq

	return p
}

func (s *Type_super_atomic_seqContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_super_atomic_seqContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NumbatParserNUMBER, 0)
}

func (s *Type_super_atomic_seqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_super_atomic_seqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_super_atomic_seqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterType_super_atomic_seq(s)
	}
}

func (s *Type_super_atomic_seqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitType_super_atomic_seq(s)
	}
}

func (p *NumbatParser) Type_super_atomic_seq() (localctx IType_super_atomic_seqContext) {
	localctx = NewType_super_atomic_seqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NumbatParserRULE_type_super_atomic_seq)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(NumbatParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserNUMBER {
		{
			p.SetState(145)
			p.Match(NumbatParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(148)
		p.Match(NumbatParserT__4)
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

// IType_super_atomicContext is an interface to support dynamic dispatch.
type IType_super_atomicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_NAME() antlr.TerminalNode
	Type_super_atomic_seq() IType_super_atomic_seqContext

	// IsType_super_atomicContext differentiates from other interfaces.
	IsType_super_atomicContext()
}

type Type_super_atomicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_super_atomicContext() *Type_super_atomicContext {
	var p = new(Type_super_atomicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_super_atomic
	return p
}

func InitEmptyType_super_atomicContext(p *Type_super_atomicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_type_super_atomic
}

func (*Type_super_atomicContext) IsType_super_atomicContext() {}

func NewType_super_atomicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_super_atomicContext {
	var p = new(Type_super_atomicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_type_super_atomic

	return p
}

func (s *Type_super_atomicContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_super_atomicContext) TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserTYPE_NAME, 0)
}

func (s *Type_super_atomicContext) Type_super_atomic_seq() IType_super_atomic_seqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_super_atomic_seqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_super_atomic_seqContext)
}

func (s *Type_super_atomicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_super_atomicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_super_atomicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterType_super_atomic(s)
	}
}

func (s *Type_super_atomicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitType_super_atomic(s)
	}
}

func (p *NumbatParser) Type_super_atomic() (localctx IType_super_atomicContext) {
	localctx = NewType_super_atomicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NumbatParserRULE_type_super_atomic)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(NumbatParserTYPE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__3 {
		{
			p.SetState(151)
			p.Type_super_atomic_seq()
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
	Type_super_atomic() IType_super_atomicContext

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

func (s *Type_outContext) Type_super_atomic() IType_super_atomicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_super_atomicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_super_atomicContext)
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
	p.EnterRule(localctx, 32, NumbatParserRULE_type_out)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Type_super_atomic()
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
	p.EnterRule(localctx, 34, NumbatParserRULE_type_in)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(156)
			p.Match(NumbatParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(NumbatParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(158)
			p.Match(NumbatParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Param()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == NumbatParserT__8 {
			{
				p.SetState(160)
				p.Match(NumbatParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(161)
				p.Param()
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(167)
			p.Match(NumbatParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
	p.EnterRule(localctx, 36, NumbatParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == NumbatParserT__6 {
			{
				p.SetState(171)
				p.Type_in()
			}

		}
		{
			p.SetState(174)
			p.Type_out()
		}

	case 2:
		{
			p.SetState(175)
			p.Type_in()
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(176)
				p.Type_out()
			}

		} else if p.HasError() { // JIM
			goto errorExit
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
	p.EnterRule(localctx, 38, NumbatParserRULE_proc_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(NumbatParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29671424) != 0 {
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case NumbatParserWS:
			{
				p.SetState(182)
				p.Match(NumbatParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserNEWLINE:
			{
				p.SetState(183)
				p.Match(NumbatParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case NumbatParserT__13, NumbatParserT__14, NumbatParserTYPE_NAME, NumbatParserNON_TYPE_NAME:
			{
				p.SetState(184)
				p.Statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(NumbatParserT__10)
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
	p.EnterRule(localctx, 40, NumbatParserRULE_proc_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
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
	p.EnterRule(localctx, 42, NumbatParserRULE_proc_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
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
	p.EnterRule(localctx, 44, NumbatParserRULE_proc_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(NumbatParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Proc_name()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__6 || _la == NumbatParserTYPE_NAME {
		{
			p.SetState(198)
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
	p.EnterRule(localctx, 46, NumbatParserRULE_proc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Proc_def()
	}
	{
		p.SetState(202)
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

// ICall_argContext is an interface to support dynamic dispatch.
type ICall_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_var() IExpr_varContext
	Expr_constant() IExpr_constantContext

	// IsCall_argContext differentiates from other interfaces.
	IsCall_argContext()
}

type Call_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_argContext() *Call_argContext {
	var p = new(Call_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_arg
	return p
}

func InitEmptyCall_argContext(p *Call_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_arg
}

func (*Call_argContext) IsCall_argContext() {}

func NewCall_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_argContext {
	var p = new(Call_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call_arg

	return p
}

func (s *Call_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_argContext) Expr_var() IExpr_varContext {
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

func (s *Call_argContext) Expr_constant() IExpr_constantContext {
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

func (s *Call_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_argContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall_arg(s)
	}
}

func (s *Call_argContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall_arg(s)
	}
}

func (p *NumbatParser) Call_arg() (localctx ICall_argContext) {
	localctx = NewCall_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NumbatParserRULE_call_arg)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserNON_TYPE_NAME:
		{
			p.SetState(204)
			p.Expr_var()
		}

	case NumbatParserT__0, NumbatParserT__1, NumbatParserT__2, NumbatParserNUMBER, NumbatParserHEX, NumbatParserSTRING:
		{
			p.SetState(205)
			p.Expr_constant()
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

// ICall_argsContext is an interface to support dynamic dispatch.
type ICall_argsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr_all() []IExpr_allContext
	Expr_all(i int) IExpr_allContext

	// IsCall_argsContext differentiates from other interfaces.
	IsCall_argsContext()
}

type Call_argsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_argsContext() *Call_argsContext {
	var p = new(Call_argsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_args
	return p
}

func InitEmptyCall_argsContext(p *Call_argsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_call_args
}

func (*Call_argsContext) IsCall_argsContext() {}

func NewCall_argsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_argsContext {
	var p = new(Call_argsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_call_args

	return p
}

func (s *Call_argsContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_argsContext) AllExpr_all() []IExpr_allContext {
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

func (s *Call_argsContext) Expr_all(i int) IExpr_allContext {
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

func (s *Call_argsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_argsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_argsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterCall_args(s)
	}
}

func (s *Call_argsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitCall_args(s)
	}
}

func (p *NumbatParser) Call_args() (localctx ICall_argsContext) {
	localctx = NewCall_argsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NumbatParserRULE_call_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Expr_all()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == NumbatParserT__8 {
		{
			p.SetState(209)
			p.Match(NumbatParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Expr_all()
		}

		p.SetState(215)
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
	p.EnterRule(localctx, 52, NumbatParserRULE_call_secondary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
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
	p.EnterRule(localctx, 54, NumbatParserRULE_call_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
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
	Call_args() ICall_argsContext

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

func (s *CallContext) Call_args() ICall_argsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_argsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_argsContext)
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
	p.EnterRule(localctx, 56, NumbatParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Call_primary()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__12 {
		{
			p.SetState(221)
			p.Match(NumbatParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Call_secondary()
		}

	}
	{
		p.SetState(225)
		p.Match(NumbatParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16252958) != 0 {
		{
			p.SetState(226)
			p.Call_args()
		}

	}
	{
		p.SetState(229)
		p.Match(NumbatParserT__7)
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
	p.EnterRule(localctx, 58, NumbatParserRULE_call_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
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

// IVar_exprContext is an interface to support dynamic dispatch.
type IVar_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_all() IExpr_allContext

	// IsVar_exprContext differentiates from other interfaces.
	IsVar_exprContext()
}

type Var_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_exprContext() *Var_exprContext {
	var p = new(Var_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_expr
	return p
}

func InitEmptyVar_exprContext(p *Var_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_expr
}

func (*Var_exprContext) IsVar_exprContext() {}

func NewVar_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_exprContext {
	var p = new(Var_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_var_expr

	return p
}

func (s *Var_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_exprContext) Expr_all() IExpr_allContext {
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

func (s *Var_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterVar_expr(s)
	}
}

func (s *Var_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitVar_expr(s)
	}
}

func (p *NumbatParser) Var_expr() (localctx IVar_exprContext) {
	localctx = NewVar_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NumbatParserRULE_var_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(NumbatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
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

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) Type_() ITypeContext {
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

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *NumbatParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NumbatParserRULE_var_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
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

// IVar_nameContext is an interface to support dynamic dispatch.
type IVar_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_TYPE_NAME() antlr.TerminalNode

	// IsVar_nameContext differentiates from other interfaces.
	IsVar_nameContext()
}

type Var_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_nameContext() *Var_nameContext {
	var p = new(Var_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_name
	return p
}

func InitEmptyVar_nameContext(p *Var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_name
}

func (*Var_nameContext) IsVar_nameContext() {}

func NewVar_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_nameContext {
	var p = new(Var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_var_name

	return p
}

func (s *Var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_nameContext) NON_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(NumbatParserNON_TYPE_NAME, 0)
}

func (s *Var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterVar_name(s)
	}
}

func (s *Var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitVar_name(s)
	}
}

func (p *NumbatParser) Var_name() (localctx IVar_nameContext) {
	localctx = NewVar_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NumbatParserRULE_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
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

// IVar_stmtContext is an interface to support dynamic dispatch.
type IVar_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_name() IVar_nameContext
	Var_type() IVar_typeContext
	Var_expr() IVar_exprContext

	// IsVar_stmtContext differentiates from other interfaces.
	IsVar_stmtContext()
}

type Var_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_stmtContext() *Var_stmtContext {
	var p = new(Var_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_stmt
	return p
}

func InitEmptyVar_stmtContext(p *Var_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = NumbatParserRULE_var_stmt
}

func (*Var_stmtContext) IsVar_stmtContext() {}

func NewVar_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_stmtContext {
	var p = new(Var_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumbatParserRULE_var_stmt

	return p
}

func (s *Var_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_stmtContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Var_stmtContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Var_stmtContext) Var_expr() IVar_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_exprContext)
}

func (s *Var_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.EnterVar_stmt(s)
	}
}

func (s *Var_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumbatListener); ok {
		listenerT.ExitVar_stmt(s)
	}
}

func (p *NumbatParser) Var_stmt() (localctx IVar_stmtContext) {
	localctx = NewVar_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, NumbatParserRULE_var_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(NumbatParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Var_name()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(242)
			p.Var_type()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__5 {
		{
			p.SetState(245)
			p.Var_expr()
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
	p.EnterRule(localctx, 68, NumbatParserRULE_assignment_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(NumbatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Expr_all()
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__8 {
		{
			p.SetState(250)
			p.Match(NumbatParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
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
	p.EnterRule(localctx, 70, NumbatParserRULE_assignment_var)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
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
	p.EnterRule(localctx, 72, NumbatParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Assignment_var()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == NumbatParserT__8 {
		{
			p.SetState(257)
			p.Match(NumbatParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Assignment_var()
		}

	}
	{
		p.SetState(261)
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
	p.EnterRule(localctx, 74, NumbatParserRULE_return_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
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
	p.EnterRule(localctx, 76, NumbatParserRULE_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(NumbatParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(266)
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
	p.EnterRule(localctx, 78, NumbatParserRULE_return_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
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
	p.EnterRule(localctx, 80, NumbatParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(NumbatParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
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
	Var_stmt() IVar_stmtContext
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

func (s *StatementContext) Var_stmt() IVar_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_stmtContext)
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
	p.EnterRule(localctx, 82, NumbatParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(274)
			p.Call_stmt()
		}

	case 2:
		{
			p.SetState(275)
			p.Var_stmt()
		}

	case 3:
		{
			p.SetState(276)
			p.Return_stmt()
		}

	case 4:
		{
			p.SetState(277)
			p.Assignment()
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
	p.EnterRule(localctx, 84, NumbatParserRULE_object)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case NumbatParserT__15:
		{
			p.SetState(280)
			p.Program()
		}

	case NumbatParserT__11:
		{
			p.SetState(281)
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
