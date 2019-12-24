// Generated from ScriptBlockParser.g4 by ANTLR 4.7.

package parser // ScriptBlockParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 278,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 7, 2, 54, 10, 2,
	12, 2, 14, 2, 57, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 64, 10, 3,
	12, 3, 14, 3, 67, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4,
	76, 10, 4, 12, 4, 14, 4, 79, 11, 4, 5, 4, 81, 10, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 7, 5, 89, 10, 5, 12, 5, 14, 5, 92, 11, 5, 5, 5, 94, 10,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 102, 10, 6, 12, 6, 14, 6,
	105, 11, 6, 5, 6, 107, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 5, 8, 117, 10, 8, 3, 8, 5, 8, 120, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 126, 10, 9, 3, 9, 3, 9, 7, 9, 130, 10, 9, 12, 9, 14, 9, 133, 11,
	9, 3, 9, 3, 9, 3, 10, 5, 10, 138, 10, 10, 3, 10, 5, 10, 141, 10, 10, 3,
	10, 5, 10, 144, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 5, 11, 151,
	10, 11, 3, 11, 5, 11, 154, 10, 11, 3, 11, 5, 11, 157, 10, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 5, 13,
	170, 10, 13, 3, 13, 5, 13, 173, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	14, 5, 14, 180, 10, 14, 3, 14, 5, 14, 183, 10, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 5, 17, 210, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 224, 10, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 245, 10, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 265, 10, 21, 12, 21,
	14, 21, 268, 11, 21, 3, 22, 5, 22, 271, 10, 22, 3, 22, 3, 22, 3, 22, 5,
	22, 276, 10, 22, 3, 22, 2, 3, 40, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 2, 2, 298, 2, 49, 3, 2,
	2, 2, 4, 60, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2, 8, 84, 3, 2, 2, 2, 10, 97,
	3, 2, 2, 2, 12, 110, 3, 2, 2, 2, 14, 113, 3, 2, 2, 2, 16, 121, 3, 2, 2,
	2, 18, 137, 3, 2, 2, 2, 20, 150, 3, 2, 2, 2, 22, 163, 3, 2, 2, 2, 24, 169,
	3, 2, 2, 2, 26, 179, 3, 2, 2, 2, 28, 189, 3, 2, 2, 2, 30, 194, 3, 2, 2,
	2, 32, 209, 3, 2, 2, 2, 34, 223, 3, 2, 2, 2, 36, 225, 3, 2, 2, 2, 38, 229,
	3, 2, 2, 2, 40, 244, 3, 2, 2, 2, 42, 270, 3, 2, 2, 2, 44, 45, 5, 30, 16,
	2, 45, 46, 7, 8, 2, 2, 46, 48, 3, 2, 2, 2, 47, 44, 3, 2, 2, 2, 48, 51,
	3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 55, 3, 2, 2, 2,
	51, 49, 3, 2, 2, 2, 52, 54, 5, 32, 17, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3,
	2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57,
	55, 3, 2, 2, 2, 58, 59, 7, 2, 2, 3, 59, 3, 3, 2, 2, 2, 60, 61, 7, 5, 2,
	2, 61, 65, 7, 8, 2, 2, 62, 64, 7, 7, 2, 2, 63, 62, 3, 2, 2, 2, 64, 67,
	3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2,
	67, 65, 3, 2, 2, 2, 68, 69, 7, 6, 2, 2, 69, 70, 7, 8, 2, 2, 70, 5, 3, 2,
	2, 2, 71, 80, 7, 19, 2, 2, 72, 77, 7, 39, 2, 2, 73, 74, 7, 25, 2, 2, 74,
	76, 7, 39, 2, 2, 75, 73, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2,
	2, 2, 77, 78, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 72,
	3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 7, 20, 2, 2,
	83, 7, 3, 2, 2, 2, 84, 93, 7, 19, 2, 2, 85, 90, 5, 40, 21, 2, 86, 87, 7,
	25, 2, 2, 87, 89, 5, 40, 21, 2, 88, 86, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2,
	90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3,
	2, 2, 2, 93, 85, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95,
	96, 7, 20, 2, 2, 96, 9, 3, 2, 2, 2, 97, 106, 7, 23, 2, 2, 98, 103, 5, 40,
	21, 2, 99, 100, 7, 25, 2, 2, 100, 102, 5, 40, 21, 2, 101, 99, 3, 2, 2,
	2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 98, 3, 2, 2, 2, 106, 107, 3,
	2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 7, 24, 2, 2, 109, 11, 3, 2, 2,
	2, 110, 111, 7, 11, 2, 2, 111, 112, 5, 8, 5, 2, 112, 13, 3, 2, 2, 2, 113,
	119, 7, 39, 2, 2, 114, 116, 5, 8, 5, 2, 115, 117, 5, 16, 9, 2, 116, 115,
	3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 120, 5, 16,
	9, 2, 119, 114, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 15, 3, 2, 2, 2,
	121, 125, 7, 21, 2, 2, 122, 123, 5, 6, 4, 2, 123, 124, 7, 18, 2, 2, 124,
	126, 3, 2, 2, 2, 125, 122, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 131, 7, 8, 2, 2, 128, 130, 5, 34, 18, 2, 129, 128, 3,
	2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2,
	2, 132, 134, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 135, 7, 22, 2, 2, 135,
	17, 3, 2, 2, 2, 136, 138, 5, 4, 3, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3,
	2, 2, 2, 138, 140, 3, 2, 2, 2, 139, 141, 5, 22, 12, 2, 140, 139, 3, 2,
	2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142, 144, 7, 17, 2, 2,
	143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145,
	146, 7, 9, 2, 2, 146, 147, 7, 39, 2, 2, 147, 148, 5, 16, 9, 2, 148, 19,
	3, 2, 2, 2, 149, 151, 5, 4, 3, 2, 150, 149, 3, 2, 2, 2, 150, 151, 3, 2,
	2, 2, 151, 153, 3, 2, 2, 2, 152, 154, 5, 22, 12, 2, 153, 152, 3, 2, 2,
	2, 153, 154, 3, 2, 2, 2, 154, 156, 3, 2, 2, 2, 155, 157, 7, 17, 2, 2, 156,
	155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159,
	7, 9, 2, 2, 159, 160, 7, 39, 2, 2, 160, 161, 7, 26, 2, 2, 161, 162, 5,
	14, 8, 2, 162, 21, 3, 2, 2, 2, 163, 164, 7, 38, 2, 2, 164, 165, 7, 39,
	2, 2, 165, 166, 7, 28, 2, 2, 166, 167, 7, 39, 2, 2, 167, 23, 3, 2, 2, 2,
	168, 170, 5, 4, 3, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	172, 3, 2, 2, 2, 171, 173, 7, 17, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173,
	3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 7, 10, 2, 2, 175, 176, 7, 39,
	2, 2, 176, 177, 5, 16, 9, 2, 177, 25, 3, 2, 2, 2, 178, 180, 5, 4, 3, 2,
	179, 178, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181,
	183, 7, 17, 2, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184,
	3, 2, 2, 2, 184, 185, 7, 12, 2, 2, 185, 186, 7, 39, 2, 2, 186, 187, 7,
	26, 2, 2, 187, 188, 5, 40, 21, 2, 188, 27, 3, 2, 2, 2, 189, 190, 7, 37,
	2, 2, 190, 191, 7, 39, 2, 2, 191, 192, 7, 36, 2, 2, 192, 193, 5, 8, 5,
	2, 193, 29, 3, 2, 2, 2, 194, 195, 7, 16, 2, 2, 195, 196, 7, 42, 2, 2, 196,
	31, 3, 2, 2, 2, 197, 198, 5, 26, 14, 2, 198, 199, 7, 8, 2, 2, 199, 210,
	3, 2, 2, 2, 200, 201, 5, 18, 10, 2, 201, 202, 7, 8, 2, 2, 202, 210, 3,
	2, 2, 2, 203, 204, 5, 24, 13, 2, 204, 205, 7, 8, 2, 2, 205, 210, 3, 2,
	2, 2, 206, 207, 5, 20, 11, 2, 207, 208, 7, 8, 2, 2, 208, 210, 3, 2, 2,
	2, 209, 197, 3, 2, 2, 2, 209, 200, 3, 2, 2, 2, 209, 203, 3, 2, 2, 2, 209,
	206, 3, 2, 2, 2, 210, 33, 3, 2, 2, 2, 211, 212, 5, 14, 8, 2, 212, 213,
	7, 8, 2, 2, 213, 224, 3, 2, 2, 2, 214, 215, 5, 12, 7, 2, 215, 216, 7, 8,
	2, 2, 216, 224, 3, 2, 2, 2, 217, 218, 5, 36, 19, 2, 218, 219, 7, 8, 2,
	2, 219, 224, 3, 2, 2, 2, 220, 221, 5, 38, 20, 2, 221, 222, 7, 8, 2, 2,
	222, 224, 3, 2, 2, 2, 223, 211, 3, 2, 2, 2, 223, 214, 3, 2, 2, 2, 223,
	217, 3, 2, 2, 2, 223, 220, 3, 2, 2, 2, 224, 35, 3, 2, 2, 2, 225, 226, 7,
	15, 2, 2, 226, 227, 5, 10, 6, 2, 227, 228, 5, 16, 9, 2, 228, 37, 3, 2,
	2, 2, 229, 230, 7, 14, 2, 2, 230, 231, 5, 22, 12, 2, 231, 39, 3, 2, 2,
	2, 232, 233, 8, 21, 1, 2, 233, 245, 5, 42, 22, 2, 234, 245, 7, 42, 2, 2,
	235, 245, 7, 39, 2, 2, 236, 237, 7, 19, 2, 2, 237, 238, 5, 40, 21, 2, 238,
	239, 7, 20, 2, 2, 239, 245, 3, 2, 2, 2, 240, 245, 5, 28, 15, 2, 241, 242,
	7, 13, 2, 2, 242, 243, 7, 39, 2, 2, 243, 245, 5, 8, 5, 2, 244, 232, 3,
	2, 2, 2, 244, 234, 3, 2, 2, 2, 244, 235, 3, 2, 2, 2, 244, 236, 3, 2, 2,
	2, 244, 240, 3, 2, 2, 2, 244, 241, 3, 2, 2, 2, 245, 266, 3, 2, 2, 2, 246,
	247, 12, 9, 2, 2, 247, 248, 7, 29, 2, 2, 248, 265, 5, 40, 21, 9, 249, 250,
	12, 8, 2, 2, 250, 251, 7, 30, 2, 2, 251, 265, 5, 40, 21, 9, 252, 253, 12,
	7, 2, 2, 253, 254, 7, 31, 2, 2, 254, 265, 5, 40, 21, 8, 255, 256, 12, 6,
	2, 2, 256, 257, 7, 32, 2, 2, 257, 265, 5, 40, 21, 7, 258, 259, 12, 5, 2,
	2, 259, 260, 7, 33, 2, 2, 260, 265, 5, 40, 21, 6, 261, 262, 12, 4, 2, 2,
	262, 263, 7, 34, 2, 2, 263, 265, 5, 40, 21, 5, 264, 246, 3, 2, 2, 2, 264,
	249, 3, 2, 2, 2, 264, 252, 3, 2, 2, 2, 264, 255, 3, 2, 2, 2, 264, 258,
	3, 2, 2, 2, 264, 261, 3, 2, 2, 2, 265, 268, 3, 2, 2, 2, 266, 264, 3, 2,
	2, 2, 266, 267, 3, 2, 2, 2, 267, 41, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2,
	269, 271, 7, 40, 2, 2, 270, 269, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271,
	272, 3, 2, 2, 2, 272, 275, 7, 41, 2, 2, 273, 274, 7, 35, 2, 2, 274, 276,
	7, 41, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 43, 3, 2,
	2, 2, 32, 49, 55, 65, 77, 80, 90, 93, 103, 106, 116, 119, 125, 131, 137,
	140, 143, 150, 153, 156, 169, 172, 179, 182, 209, 223, 244, 264, 266, 270,
	275,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'/*'", "'*/'", "", "", "'func'", "'script'", "'command'",
	"'const'", "'run'", "'raise'", "'delay'", "'import'", "'internal'", "'->'",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'='", "';'", "':'", "'^'",
	"'*'", "'/'", "'//'", "'+'", "'-'", "'.'", "'>'", "'<'", "'#'",
}
var symbolicNames = []string{
	"", "WS", "COMMENT", "DOC_START", "DOC_END", "DOC_LINE", "NEWLINES", "FUNCTION",
	"TEMPLATE", "NATIVE", "CONSTANT", "RUN", "RAISE", "DELAY", "IMPORT", "INTERNAL",
	"ARROW", "OPAREN", "CPAREN", "OCURLY", "CCURLY", "OSQUARE", "CSQUARE",
	"COMMA", "EQUALS", "SEMICOLON", "COLON", "POWER", "MULTIPLY", "DIVIDE",
	"INTEGER_DIVIDE", "PLUS", "SUBTRACT", "DOT", "GREATER_THAN", "LESS_THAN",
	"POUND", "IDENTIFIER", "SIGN", "DIGITS", "STRING",
}

var ruleNames = []string{
	"unit", "documentation", "parameterList", "argumentList", "structureList",
	"nativeCall", "functionCall", "functionFrame", "functionDefinition", "functionDefinitionShortcut",
	"tag", "templateDefinition", "constantDefinition", "formatter", "importLine",
	"topDefinition", "statement", "delayStructure", "raise", "expression",
	"number",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ScriptBlockParser struct {
	*antlr.BaseParser
}

func NewScriptBlockParser(input antlr.TokenStream) *ScriptBlockParser {
	this := new(ScriptBlockParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ScriptBlockParser.g4"

	return this
}

// ScriptBlockParser tokens.
const (
	ScriptBlockParserEOF            = antlr.TokenEOF
	ScriptBlockParserWS             = 1
	ScriptBlockParserCOMMENT        = 2
	ScriptBlockParserDOC_START      = 3
	ScriptBlockParserDOC_END        = 4
	ScriptBlockParserDOC_LINE       = 5
	ScriptBlockParserNEWLINES       = 6
	ScriptBlockParserFUNCTION       = 7
	ScriptBlockParserTEMPLATE       = 8
	ScriptBlockParserNATIVE         = 9
	ScriptBlockParserCONSTANT       = 10
	ScriptBlockParserRUN            = 11
	ScriptBlockParserRAISE          = 12
	ScriptBlockParserDELAY          = 13
	ScriptBlockParserIMPORT         = 14
	ScriptBlockParserINTERNAL       = 15
	ScriptBlockParserARROW          = 16
	ScriptBlockParserOPAREN         = 17
	ScriptBlockParserCPAREN         = 18
	ScriptBlockParserOCURLY         = 19
	ScriptBlockParserCCURLY         = 20
	ScriptBlockParserOSQUARE        = 21
	ScriptBlockParserCSQUARE        = 22
	ScriptBlockParserCOMMA          = 23
	ScriptBlockParserEQUALS         = 24
	ScriptBlockParserSEMICOLON      = 25
	ScriptBlockParserCOLON          = 26
	ScriptBlockParserPOWER          = 27
	ScriptBlockParserMULTIPLY       = 28
	ScriptBlockParserDIVIDE         = 29
	ScriptBlockParserINTEGER_DIVIDE = 30
	ScriptBlockParserPLUS           = 31
	ScriptBlockParserSUBTRACT       = 32
	ScriptBlockParserDOT            = 33
	ScriptBlockParserGREATER_THAN   = 34
	ScriptBlockParserLESS_THAN      = 35
	ScriptBlockParserPOUND          = 36
	ScriptBlockParserIDENTIFIER     = 37
	ScriptBlockParserSIGN           = 38
	ScriptBlockParserDIGITS         = 39
	ScriptBlockParserSTRING         = 40
)

// ScriptBlockParser rules.
const (
	ScriptBlockParserRULE_unit                       = 0
	ScriptBlockParserRULE_documentation              = 1
	ScriptBlockParserRULE_parameterList              = 2
	ScriptBlockParserRULE_argumentList               = 3
	ScriptBlockParserRULE_structureList              = 4
	ScriptBlockParserRULE_nativeCall                 = 5
	ScriptBlockParserRULE_functionCall               = 6
	ScriptBlockParserRULE_functionFrame              = 7
	ScriptBlockParserRULE_functionDefinition         = 8
	ScriptBlockParserRULE_functionDefinitionShortcut = 9
	ScriptBlockParserRULE_tag                        = 10
	ScriptBlockParserRULE_templateDefinition         = 11
	ScriptBlockParserRULE_constantDefinition         = 12
	ScriptBlockParserRULE_formatter                  = 13
	ScriptBlockParserRULE_importLine                 = 14
	ScriptBlockParserRULE_topDefinition              = 15
	ScriptBlockParserRULE_statement                  = 16
	ScriptBlockParserRULE_delayStructure             = 17
	ScriptBlockParserRULE_raise                      = 18
	ScriptBlockParserRULE_expression                 = 19
	ScriptBlockParserRULE_number                     = 20
)

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserEOF, 0)
}

func (s *UnitContext) AllImportLine() []IImportLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportLineContext)(nil)).Elem())
	var tst = make([]IImportLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportLineContext)
		}
	}

	return tst
}

func (s *UnitContext) ImportLine(i int) IImportLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportLineContext)
}

func (s *UnitContext) AllNEWLINES() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserNEWLINES)
}

func (s *UnitContext) NEWLINES(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, i)
}

func (s *UnitContext) AllTopDefinition() []ITopDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopDefinitionContext)(nil)).Elem())
	var tst = make([]ITopDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopDefinitionContext)
		}
	}

	return tst
}

func (s *UnitContext) TopDefinition(i int) ITopDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopDefinitionContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *ScriptBlockParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ScriptBlockParserRULE_unit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScriptBlockParserIMPORT {
		{
			p.SetState(42)
			p.ImportLine()
		}
		{
			p.SetState(43)
			p.Match(ScriptBlockParserNEWLINES)
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScriptBlockParserDOC_START)|(1<<ScriptBlockParserFUNCTION)|(1<<ScriptBlockParserTEMPLATE)|(1<<ScriptBlockParserCONSTANT)|(1<<ScriptBlockParserINTERNAL))) != 0) || _la == ScriptBlockParserPOUND {
		{
			p.SetState(50)
			p.TopDefinition()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(ScriptBlockParserEOF)
	}

	return localctx
}

// IDocumentationContext is an interface to support dynamic dispatch.
type IDocumentationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentationContext differentiates from other interfaces.
	IsDocumentationContext()
}

type DocumentationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentationContext() *DocumentationContext {
	var p = new(DocumentationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_documentation
	return p
}

func (*DocumentationContext) IsDocumentationContext() {}

func NewDocumentationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentationContext {
	var p = new(DocumentationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_documentation

	return p
}

func (s *DocumentationContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentationContext) DOC_START() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDOC_START, 0)
}

func (s *DocumentationContext) AllNEWLINES() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserNEWLINES)
}

func (s *DocumentationContext) NEWLINES(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, i)
}

func (s *DocumentationContext) DOC_END() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDOC_END, 0)
}

func (s *DocumentationContext) AllDOC_LINE() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserDOC_LINE)
}

func (s *DocumentationContext) DOC_LINE(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDOC_LINE, i)
}

func (s *DocumentationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterDocumentation(s)
	}
}

func (s *DocumentationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitDocumentation(s)
	}
}

func (p *ScriptBlockParser) Documentation() (localctx IDocumentationContext) {
	localctx = NewDocumentationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ScriptBlockParserRULE_documentation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(ScriptBlockParserDOC_START)
	}
	{
		p.SetState(59)
		p.Match(ScriptBlockParserNEWLINES)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScriptBlockParserDOC_LINE {
		{
			p.SetState(60)
			p.Match(ScriptBlockParserDOC_LINE)
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(ScriptBlockParserDOC_END)
	}
	{
		p.SetState(67)
		p.Match(ScriptBlockParserNEWLINES)
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserOPAREN, 0)
}

func (s *ParameterListContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCPAREN, 0)
}

func (s *ParameterListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserIDENTIFIER)
}

func (s *ParameterListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, i)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *ScriptBlockParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ScriptBlockParserRULE_parameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(ScriptBlockParserOPAREN)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserIDENTIFIER {
		{
			p.SetState(70)
			p.Match(ScriptBlockParserIDENTIFIER)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScriptBlockParserCOMMA {
			{
				p.SetState(71)
				p.Match(ScriptBlockParserCOMMA)
			}
			{
				p.SetState(72)
				p.Match(ScriptBlockParserIDENTIFIER)
			}

			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(80)
		p.Match(ScriptBlockParserCPAREN)
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserOPAREN, 0)
}

func (s *ArgumentListContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCPAREN, 0)
}

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *ScriptBlockParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ScriptBlockParserRULE_argumentList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(ScriptBlockParserOPAREN)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(ScriptBlockParserRUN-11))|(1<<(ScriptBlockParserOPAREN-11))|(1<<(ScriptBlockParserLESS_THAN-11))|(1<<(ScriptBlockParserIDENTIFIER-11))|(1<<(ScriptBlockParserSIGN-11))|(1<<(ScriptBlockParserDIGITS-11))|(1<<(ScriptBlockParserSTRING-11)))) != 0 {
		{
			p.SetState(83)
			p.expression(0)
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScriptBlockParserCOMMA {
			{
				p.SetState(84)
				p.Match(ScriptBlockParserCOMMA)
			}
			{
				p.SetState(85)
				p.expression(0)
			}

			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(93)
		p.Match(ScriptBlockParserCPAREN)
	}

	return localctx
}

// IStructureListContext is an interface to support dynamic dispatch.
type IStructureListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructureListContext differentiates from other interfaces.
	IsStructureListContext()
}

type StructureListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructureListContext() *StructureListContext {
	var p = new(StructureListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_structureList
	return p
}

func (*StructureListContext) IsStructureListContext() {}

func NewStructureListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructureListContext {
	var p = new(StructureListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_structureList

	return p
}

func (s *StructureListContext) GetParser() antlr.Parser { return s.parser }

func (s *StructureListContext) OSQUARE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserOSQUARE, 0)
}

func (s *StructureListContext) CSQUARE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCSQUARE, 0)
}

func (s *StructureListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *StructureListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StructureListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserCOMMA)
}

func (s *StructureListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCOMMA, i)
}

func (s *StructureListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructureListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructureListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterStructureList(s)
	}
}

func (s *StructureListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitStructureList(s)
	}
}

func (p *ScriptBlockParser) StructureList() (localctx IStructureListContext) {
	localctx = NewStructureListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ScriptBlockParserRULE_structureList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(ScriptBlockParserOSQUARE)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(ScriptBlockParserRUN-11))|(1<<(ScriptBlockParserOPAREN-11))|(1<<(ScriptBlockParserLESS_THAN-11))|(1<<(ScriptBlockParserIDENTIFIER-11))|(1<<(ScriptBlockParserSIGN-11))|(1<<(ScriptBlockParserDIGITS-11))|(1<<(ScriptBlockParserSTRING-11)))) != 0 {
		{
			p.SetState(96)
			p.expression(0)
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScriptBlockParserCOMMA {
			{
				p.SetState(97)
				p.Match(ScriptBlockParserCOMMA)
			}
			{
				p.SetState(98)
				p.expression(0)
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(106)
		p.Match(ScriptBlockParserCSQUARE)
	}

	return localctx
}

// INativeCallContext is an interface to support dynamic dispatch.
type INativeCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNativeCallContext differentiates from other interfaces.
	IsNativeCallContext()
}

type NativeCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNativeCallContext() *NativeCallContext {
	var p = new(NativeCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_nativeCall
	return p
}

func (*NativeCallContext) IsNativeCallContext() {}

func NewNativeCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NativeCallContext {
	var p = new(NativeCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_nativeCall

	return p
}

func (s *NativeCallContext) GetParser() antlr.Parser { return s.parser }

func (s *NativeCallContext) NATIVE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNATIVE, 0)
}

func (s *NativeCallContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *NativeCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NativeCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterNativeCall(s)
	}
}

func (s *NativeCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitNativeCall(s)
	}
}

func (p *ScriptBlockParser) NativeCall() (localctx INativeCallContext) {
	localctx = NewNativeCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ScriptBlockParserRULE_nativeCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(ScriptBlockParserNATIVE)
	}
	{
		p.SetState(109)
		p.ArgumentList()
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) FunctionFrame() IFunctionFrameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionFrameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionFrameContext)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *ScriptBlockParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ScriptBlockParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScriptBlockParserOPAREN:
		{
			p.SetState(112)
			p.ArgumentList()
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScriptBlockParserOCURLY {
			{
				p.SetState(113)
				p.FunctionFrame()
			}

		}

	case ScriptBlockParserOCURLY:
		{
			p.SetState(116)
			p.FunctionFrame()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionFrameContext is an interface to support dynamic dispatch.
type IFunctionFrameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionFrameContext differentiates from other interfaces.
	IsFunctionFrameContext()
}

type FunctionFrameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionFrameContext() *FunctionFrameContext {
	var p = new(FunctionFrameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_functionFrame
	return p
}

func (*FunctionFrameContext) IsFunctionFrameContext() {}

func NewFunctionFrameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionFrameContext {
	var p = new(FunctionFrameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_functionFrame

	return p
}

func (s *FunctionFrameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionFrameContext) OCURLY() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserOCURLY, 0)
}

func (s *FunctionFrameContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *FunctionFrameContext) CCURLY() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCCURLY, 0)
}

func (s *FunctionFrameContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionFrameContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserARROW, 0)
}

func (s *FunctionFrameContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FunctionFrameContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunctionFrameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionFrameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionFrameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionFrame(s)
	}
}

func (s *FunctionFrameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionFrame(s)
	}
}

func (p *ScriptBlockParser) FunctionFrame() (localctx IFunctionFrameContext) {
	localctx = NewFunctionFrameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ScriptBlockParserRULE_functionFrame)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(ScriptBlockParserOCURLY)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserOPAREN {
		{
			p.SetState(120)
			p.ParameterList()
		}
		{
			p.SetState(121)
			p.Match(ScriptBlockParserARROW)
		}

	}
	{
		p.SetState(125)
		p.Match(ScriptBlockParserNEWLINES)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ScriptBlockParserNATIVE-9))|(1<<(ScriptBlockParserRAISE-9))|(1<<(ScriptBlockParserDELAY-9))|(1<<(ScriptBlockParserIDENTIFIER-9)))) != 0 {
		{
			p.SetState(126)
			p.Statement()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
		p.Match(ScriptBlockParserCCURLY)
	}

	return localctx
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserFUNCTION, 0)
}

func (s *FunctionDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *FunctionDefinitionContext) FunctionFrame() IFunctionFrameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionFrameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionFrameContext)
}

func (s *FunctionDefinitionContext) Documentation() IDocumentationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentationContext)
}

func (s *FunctionDefinitionContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *FunctionDefinitionContext) INTERNAL() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserINTERNAL, 0)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *ScriptBlockParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ScriptBlockParserRULE_functionDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserDOC_START {
		{
			p.SetState(134)
			p.Documentation()
		}

	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserPOUND {
		{
			p.SetState(137)
			p.Tag()
		}

	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserINTERNAL {
		{
			p.SetState(140)
			p.Match(ScriptBlockParserINTERNAL)
		}

	}
	{
		p.SetState(143)
		p.Match(ScriptBlockParserFUNCTION)
	}
	{
		p.SetState(144)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	{
		p.SetState(145)
		p.FunctionFrame()
	}

	return localctx
}

// IFunctionDefinitionShortcutContext is an interface to support dynamic dispatch.
type IFunctionDefinitionShortcutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionShortcutContext differentiates from other interfaces.
	IsFunctionDefinitionShortcutContext()
}

type FunctionDefinitionShortcutContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionShortcutContext() *FunctionDefinitionShortcutContext {
	var p = new(FunctionDefinitionShortcutContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_functionDefinitionShortcut
	return p
}

func (*FunctionDefinitionShortcutContext) IsFunctionDefinitionShortcutContext() {}

func NewFunctionDefinitionShortcutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionShortcutContext {
	var p = new(FunctionDefinitionShortcutContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_functionDefinitionShortcut

	return p
}

func (s *FunctionDefinitionShortcutContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionShortcutContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserFUNCTION, 0)
}

func (s *FunctionDefinitionShortcutContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *FunctionDefinitionShortcutContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserEQUALS, 0)
}

func (s *FunctionDefinitionShortcutContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionDefinitionShortcutContext) Documentation() IDocumentationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentationContext)
}

func (s *FunctionDefinitionShortcutContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *FunctionDefinitionShortcutContext) INTERNAL() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserINTERNAL, 0)
}

func (s *FunctionDefinitionShortcutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionShortcutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionShortcutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionDefinitionShortcut(s)
	}
}

func (s *FunctionDefinitionShortcutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionDefinitionShortcut(s)
	}
}

func (p *ScriptBlockParser) FunctionDefinitionShortcut() (localctx IFunctionDefinitionShortcutContext) {
	localctx = NewFunctionDefinitionShortcutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ScriptBlockParserRULE_functionDefinitionShortcut)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserDOC_START {
		{
			p.SetState(147)
			p.Documentation()
		}

	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserPOUND {
		{
			p.SetState(150)
			p.Tag()
		}

	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserINTERNAL {
		{
			p.SetState(153)
			p.Match(ScriptBlockParserINTERNAL)
		}

	}
	{
		p.SetState(156)
		p.Match(ScriptBlockParserFUNCTION)
	}
	{
		p.SetState(157)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	{
		p.SetState(158)
		p.Match(ScriptBlockParserEQUALS)
	}
	{
		p.SetState(159)
		p.FunctionCall()
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) POUND() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserPOUND, 0)
}

func (s *TagContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserIDENTIFIER)
}

func (s *TagContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, i)
}

func (s *TagContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCOLON, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *ScriptBlockParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ScriptBlockParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(ScriptBlockParserPOUND)
	}
	{
		p.SetState(162)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	{
		p.SetState(163)
		p.Match(ScriptBlockParserCOLON)
	}
	{
		p.SetState(164)
		p.Match(ScriptBlockParserIDENTIFIER)
	}

	return localctx
}

// ITemplateDefinitionContext is an interface to support dynamic dispatch.
type ITemplateDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateDefinitionContext differentiates from other interfaces.
	IsTemplateDefinitionContext()
}

type TemplateDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateDefinitionContext() *TemplateDefinitionContext {
	var p = new(TemplateDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_templateDefinition
	return p
}

func (*TemplateDefinitionContext) IsTemplateDefinitionContext() {}

func NewTemplateDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateDefinitionContext {
	var p = new(TemplateDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_templateDefinition

	return p
}

func (s *TemplateDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateDefinitionContext) TEMPLATE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserTEMPLATE, 0)
}

func (s *TemplateDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *TemplateDefinitionContext) FunctionFrame() IFunctionFrameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionFrameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionFrameContext)
}

func (s *TemplateDefinitionContext) Documentation() IDocumentationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentationContext)
}

func (s *TemplateDefinitionContext) INTERNAL() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserINTERNAL, 0)
}

func (s *TemplateDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterTemplateDefinition(s)
	}
}

func (s *TemplateDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitTemplateDefinition(s)
	}
}

func (p *ScriptBlockParser) TemplateDefinition() (localctx ITemplateDefinitionContext) {
	localctx = NewTemplateDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ScriptBlockParserRULE_templateDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserDOC_START {
		{
			p.SetState(166)
			p.Documentation()
		}

	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserINTERNAL {
		{
			p.SetState(169)
			p.Match(ScriptBlockParserINTERNAL)
		}

	}
	{
		p.SetState(172)
		p.Match(ScriptBlockParserTEMPLATE)
	}
	{
		p.SetState(173)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	{
		p.SetState(174)
		p.FunctionFrame()
	}

	return localctx
}

// IConstantDefinitionContext is an interface to support dynamic dispatch.
type IConstantDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantDefinitionContext differentiates from other interfaces.
	IsConstantDefinitionContext()
}

type ConstantDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDefinitionContext() *ConstantDefinitionContext {
	var p = new(ConstantDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_constantDefinition
	return p
}

func (*ConstantDefinitionContext) IsConstantDefinitionContext() {}

func NewConstantDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefinitionContext {
	var p = new(ConstantDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_constantDefinition

	return p
}

func (s *ConstantDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefinitionContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCONSTANT, 0)
}

func (s *ConstantDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *ConstantDefinitionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserEQUALS, 0)
}

func (s *ConstantDefinitionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstantDefinitionContext) Documentation() IDocumentationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentationContext)
}

func (s *ConstantDefinitionContext) INTERNAL() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserINTERNAL, 0)
}

func (s *ConstantDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterConstantDefinition(s)
	}
}

func (s *ConstantDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitConstantDefinition(s)
	}
}

func (p *ScriptBlockParser) ConstantDefinition() (localctx IConstantDefinitionContext) {
	localctx = NewConstantDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ScriptBlockParserRULE_constantDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserDOC_START {
		{
			p.SetState(176)
			p.Documentation()
		}

	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserINTERNAL {
		{
			p.SetState(179)
			p.Match(ScriptBlockParserINTERNAL)
		}

	}
	{
		p.SetState(182)
		p.Match(ScriptBlockParserCONSTANT)
	}
	{
		p.SetState(183)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	{
		p.SetState(184)
		p.Match(ScriptBlockParserEQUALS)
	}
	{
		p.SetState(185)
		p.expression(0)
	}

	return localctx
}

// IFormatterContext is an interface to support dynamic dispatch.
type IFormatterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormatterContext differentiates from other interfaces.
	IsFormatterContext()
}

type FormatterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatterContext() *FormatterContext {
	var p = new(FormatterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_formatter
	return p
}

func (*FormatterContext) IsFormatterContext() {}

func NewFormatterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatterContext {
	var p = new(FormatterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_formatter

	return p
}

func (s *FormatterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatterContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserLESS_THAN, 0)
}

func (s *FormatterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *FormatterContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserGREATER_THAN, 0)
}

func (s *FormatterContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FormatterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFormatter(s)
	}
}

func (s *FormatterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFormatter(s)
	}
}

func (p *ScriptBlockParser) Formatter() (localctx IFormatterContext) {
	localctx = NewFormatterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ScriptBlockParserRULE_formatter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(ScriptBlockParserLESS_THAN)
	}
	{
		p.SetState(188)
		p.Match(ScriptBlockParserIDENTIFIER)
	}
	{
		p.SetState(189)
		p.Match(ScriptBlockParserGREATER_THAN)
	}
	{
		p.SetState(190)
		p.ArgumentList()
	}

	return localctx
}

// IImportLineContext is an interface to support dynamic dispatch.
type IImportLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportLineContext differentiates from other interfaces.
	IsImportLineContext()
}

type ImportLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportLineContext() *ImportLineContext {
	var p = new(ImportLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_importLine
	return p
}

func (*ImportLineContext) IsImportLineContext() {}

func NewImportLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportLineContext {
	var p = new(ImportLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_importLine

	return p
}

func (s *ImportLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportLineContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIMPORT, 0)
}

func (s *ImportLineContext) STRING() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserSTRING, 0)
}

func (s *ImportLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterImportLine(s)
	}
}

func (s *ImportLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitImportLine(s)
	}
}

func (p *ScriptBlockParser) ImportLine() (localctx IImportLineContext) {
	localctx = NewImportLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ScriptBlockParserRULE_importLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(ScriptBlockParserIMPORT)
	}
	{
		p.SetState(193)
		p.Match(ScriptBlockParserSTRING)
	}

	return localctx
}

// ITopDefinitionContext is an interface to support dynamic dispatch.
type ITopDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopDefinitionContext differentiates from other interfaces.
	IsTopDefinitionContext()
}

type TopDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDefinitionContext() *TopDefinitionContext {
	var p = new(TopDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_topDefinition
	return p
}

func (*TopDefinitionContext) IsTopDefinitionContext() {}

func NewTopDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDefinitionContext {
	var p = new(TopDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_topDefinition

	return p
}

func (s *TopDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDefinitionContext) CopyFrom(ctx *TopDefinitionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TopDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionDefinitionTopContext struct {
	*TopDefinitionContext
}

func NewFunctionDefinitionTopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefinitionTopContext {
	var p = new(FunctionDefinitionTopContext)

	p.TopDefinitionContext = NewEmptyTopDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopDefinitionContext))

	return p
}

func (s *FunctionDefinitionTopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionTopContext) FunctionDefinition() IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *FunctionDefinitionTopContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *FunctionDefinitionTopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionDefinitionTop(s)
	}
}

func (s *FunctionDefinitionTopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionDefinitionTop(s)
	}
}

type FunctionShortcutTopContext struct {
	*TopDefinitionContext
}

func NewFunctionShortcutTopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionShortcutTopContext {
	var p = new(FunctionShortcutTopContext)

	p.TopDefinitionContext = NewEmptyTopDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopDefinitionContext))

	return p
}

func (s *FunctionShortcutTopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionShortcutTopContext) FunctionDefinitionShortcut() IFunctionDefinitionShortcutContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionShortcutContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionShortcutContext)
}

func (s *FunctionShortcutTopContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *FunctionShortcutTopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionShortcutTop(s)
	}
}

func (s *FunctionShortcutTopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionShortcutTop(s)
	}
}

type ConstantDefinitionTopContext struct {
	*TopDefinitionContext
}

func NewConstantDefinitionTopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantDefinitionTopContext {
	var p = new(ConstantDefinitionTopContext)

	p.TopDefinitionContext = NewEmptyTopDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopDefinitionContext))

	return p
}

func (s *ConstantDefinitionTopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefinitionTopContext) ConstantDefinition() IConstantDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantDefinitionContext)
}

func (s *ConstantDefinitionTopContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *ConstantDefinitionTopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterConstantDefinitionTop(s)
	}
}

func (s *ConstantDefinitionTopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitConstantDefinitionTop(s)
	}
}

type TemplateDefinitionTopContext struct {
	*TopDefinitionContext
}

func NewTemplateDefinitionTopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TemplateDefinitionTopContext {
	var p = new(TemplateDefinitionTopContext)

	p.TopDefinitionContext = NewEmptyTopDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopDefinitionContext))

	return p
}

func (s *TemplateDefinitionTopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateDefinitionTopContext) TemplateDefinition() ITemplateDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateDefinitionContext)
}

func (s *TemplateDefinitionTopContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *TemplateDefinitionTopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterTemplateDefinitionTop(s)
	}
}

func (s *TemplateDefinitionTopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitTemplateDefinitionTop(s)
	}
}

func (p *ScriptBlockParser) TopDefinition() (localctx ITopDefinitionContext) {
	localctx = NewTopDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ScriptBlockParserRULE_topDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstantDefinitionTopContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.ConstantDefinition()
		}
		{
			p.SetState(196)
			p.Match(ScriptBlockParserNEWLINES)
		}

	case 2:
		localctx = NewFunctionDefinitionTopContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.FunctionDefinition()
		}
		{
			p.SetState(199)
			p.Match(ScriptBlockParserNEWLINES)
		}

	case 3:
		localctx = NewTemplateDefinitionTopContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(201)
			p.TemplateDefinition()
		}
		{
			p.SetState(202)
			p.Match(ScriptBlockParserNEWLINES)
		}

	case 4:
		localctx = NewFunctionShortcutTopContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.FunctionDefinitionShortcut()
		}
		{
			p.SetState(205)
			p.Match(ScriptBlockParserNEWLINES)
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RaiseStatementContext struct {
	*StatementContext
}

func NewRaiseStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RaiseStatementContext {
	var p = new(RaiseStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *RaiseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RaiseStatementContext) Raise() IRaiseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRaiseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRaiseContext)
}

func (s *RaiseStatementContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *RaiseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterRaiseStatement(s)
	}
}

func (s *RaiseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitRaiseStatement(s)
	}
}

type FunctionCallStatementContext struct {
	*StatementContext
}

func NewFunctionCallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FunctionCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallStatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallStatementContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *FunctionCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFunctionCallStatement(s)
	}
}

type DelayStructureStatementContext struct {
	*StatementContext
}

func NewDelayStructureStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DelayStructureStatementContext {
	var p = new(DelayStructureStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *DelayStructureStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelayStructureStatementContext) DelayStructure() IDelayStructureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelayStructureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelayStructureContext)
}

func (s *DelayStructureStatementContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *DelayStructureStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterDelayStructureStatement(s)
	}
}

func (s *DelayStructureStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitDelayStructureStatement(s)
	}
}

type NativeCallStatementContext struct {
	*StatementContext
}

func NewNativeCallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NativeCallStatementContext {
	var p = new(NativeCallStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NativeCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeCallStatementContext) NativeCall() INativeCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INativeCallContext)
}

func (s *NativeCallStatementContext) NEWLINES() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserNEWLINES, 0)
}

func (s *NativeCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterNativeCallStatement(s)
	}
}

func (s *NativeCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitNativeCallStatement(s)
	}
}

func (p *ScriptBlockParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ScriptBlockParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScriptBlockParserIDENTIFIER:
		localctx = NewFunctionCallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.FunctionCall()
		}
		{
			p.SetState(210)
			p.Match(ScriptBlockParserNEWLINES)
		}

	case ScriptBlockParserNATIVE:
		localctx = NewNativeCallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.NativeCall()
		}
		{
			p.SetState(213)
			p.Match(ScriptBlockParserNEWLINES)
		}

	case ScriptBlockParserDELAY:
		localctx = NewDelayStructureStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(215)
			p.DelayStructure()
		}
		{
			p.SetState(216)
			p.Match(ScriptBlockParserNEWLINES)
		}

	case ScriptBlockParserRAISE:
		localctx = NewRaiseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(218)
			p.Raise()
		}
		{
			p.SetState(219)
			p.Match(ScriptBlockParserNEWLINES)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDelayStructureContext is an interface to support dynamic dispatch.
type IDelayStructureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelayStructureContext differentiates from other interfaces.
	IsDelayStructureContext()
}

type DelayStructureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelayStructureContext() *DelayStructureContext {
	var p = new(DelayStructureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_delayStructure
	return p
}

func (*DelayStructureContext) IsDelayStructureContext() {}

func NewDelayStructureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelayStructureContext {
	var p = new(DelayStructureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_delayStructure

	return p
}

func (s *DelayStructureContext) GetParser() antlr.Parser { return s.parser }

func (s *DelayStructureContext) DELAY() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDELAY, 0)
}

func (s *DelayStructureContext) StructureList() IStructureListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructureListContext)
}

func (s *DelayStructureContext) FunctionFrame() IFunctionFrameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionFrameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionFrameContext)
}

func (s *DelayStructureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelayStructureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelayStructureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterDelayStructure(s)
	}
}

func (s *DelayStructureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitDelayStructure(s)
	}
}

func (p *ScriptBlockParser) DelayStructure() (localctx IDelayStructureContext) {
	localctx = NewDelayStructureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ScriptBlockParserRULE_delayStructure)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(ScriptBlockParserDELAY)
	}
	{
		p.SetState(224)
		p.StructureList()
	}
	{
		p.SetState(225)
		p.FunctionFrame()
	}

	return localctx
}

// IRaiseContext is an interface to support dynamic dispatch.
type IRaiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRaiseContext differentiates from other interfaces.
	IsRaiseContext()
}

type RaiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRaiseContext() *RaiseContext {
	var p = new(RaiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_raise
	return p
}

func (*RaiseContext) IsRaiseContext() {}

func NewRaiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RaiseContext {
	var p = new(RaiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_raise

	return p
}

func (s *RaiseContext) GetParser() antlr.Parser { return s.parser }

func (s *RaiseContext) RAISE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserRAISE, 0)
}

func (s *RaiseContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *RaiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RaiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RaiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterRaise(s)
	}
}

func (s *RaiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitRaise(s)
	}
}

func (p *ScriptBlockParser) Raise() (localctx IRaiseContext) {
	localctx = NewRaiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ScriptBlockParserRULE_raise)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(ScriptBlockParserRAISE)
	}
	{
		p.SetState(228)
		p.Tag()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringExprContext struct {
	*ExpressionContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserSTRING, 0)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

type DivideExprContext struct {
	*ExpressionContext
}

func NewDivideExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideExprContext {
	var p = new(DivideExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DivideExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DivideExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DivideExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDIVIDE, 0)
}

func (s *DivideExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterDivideExpr(s)
	}
}

func (s *DivideExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitDivideExpr(s)
	}
}

type IntegerDivideExprContext struct {
	*ExpressionContext
}

func NewIntegerDivideExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerDivideExprContext {
	var p = new(IntegerDivideExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntegerDivideExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerDivideExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IntegerDivideExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IntegerDivideExprContext) INTEGER_DIVIDE() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserINTEGER_DIVIDE, 0)
}

func (s *IntegerDivideExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterIntegerDivideExpr(s)
	}
}

func (s *IntegerDivideExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitIntegerDivideExpr(s)
	}
}

type SubtractExprContext struct {
	*ExpressionContext
}

func NewSubtractExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractExprContext {
	var p = new(SubtractExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SubtractExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SubtractExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubtractExprContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserSUBTRACT, 0)
}

func (s *SubtractExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterSubtractExpr(s)
	}
}

func (s *SubtractExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitSubtractExpr(s)
	}
}

type PowerExprContext struct {
	*ExpressionContext
}

func NewPowerExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExprContext {
	var p = new(PowerExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowerExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExprContext) POWER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserPOWER, 0)
}

func (s *PowerExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterPowerExpr(s)
	}
}

func (s *PowerExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitPowerExpr(s)
	}
}

type AddExprContext struct {
	*ExpressionContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserPLUS, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

type NumberExprContext struct {
	*ExpressionContext
}

func NewNumberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExprContext {
	var p = new(NumberExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *NumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterNumberExpr(s)
	}
}

func (s *NumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitNumberExpr(s)
	}
}

type MultiplyExprContext struct {
	*ExpressionContext
}

func NewMultiplyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyExprContext {
	var p = new(MultiplyExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultiplyExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplyExprContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserMULTIPLY, 0)
}

func (s *MultiplyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterMultiplyExpr(s)
	}
}

func (s *MultiplyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitMultiplyExpr(s)
	}
}

type CallExprContext struct {
	*ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) RUN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserRUN, 0)
}

func (s *CallExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *CallExprContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

type FormatterExprContext struct {
	*ExpressionContext
}

func NewFormatterExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FormatterExprContext {
	var p = new(FormatterExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FormatterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatterExprContext) Formatter() IFormatterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormatterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormatterContext)
}

func (s *FormatterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterFormatterExpr(s)
	}
}

func (s *FormatterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitFormatterExpr(s)
	}
}

type IdentifierExprContext struct {
	*ExpressionContext
}

func NewIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExprContext {
	var p = new(IdentifierExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserIDENTIFIER, 0)
}

func (s *IdentifierExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitIdentifierExpr(s)
	}
}

type ParenthExprContext struct {
	*ExpressionContext
}

func NewParenthExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthExprContext {
	var p = new(ParenthExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthExprContext) OPAREN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserOPAREN, 0)
}

func (s *ParenthExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthExprContext) CPAREN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserCPAREN, 0)
}

func (s *ParenthExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterParenthExpr(s)
	}
}

func (s *ParenthExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitParenthExpr(s)
	}
}

func (p *ScriptBlockParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ScriptBlockParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, ScriptBlockParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScriptBlockParserSIGN, ScriptBlockParserDIGITS:
		localctx = NewNumberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(231)
			p.Number()
		}

	case ScriptBlockParserSTRING:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(ScriptBlockParserSTRING)
		}

	case ScriptBlockParserIDENTIFIER:
		localctx = NewIdentifierExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(233)
			p.Match(ScriptBlockParserIDENTIFIER)
		}

	case ScriptBlockParserOPAREN:
		localctx = NewParenthExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(ScriptBlockParserOPAREN)
		}
		{
			p.SetState(235)
			p.expression(0)
		}
		{
			p.SetState(236)
			p.Match(ScriptBlockParserCPAREN)
		}

	case ScriptBlockParserLESS_THAN:
		localctx = NewFormatterExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Formatter()
		}

	case ScriptBlockParserRUN:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)
			p.Match(ScriptBlockParserRUN)
		}
		{
			p.SetState(240)
			p.Match(ScriptBlockParserIDENTIFIER)
		}
		{
			p.SetState(241)
			p.ArgumentList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ScriptBlockParserRULE_expression)
				p.SetState(244)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(245)
					p.Match(ScriptBlockParserPOWER)
				}
				{
					p.SetState(246)
					p.expression(7)
				}

			case 2:
				localctx = NewMultiplyExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ScriptBlockParserRULE_expression)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(248)
					p.Match(ScriptBlockParserMULTIPLY)
				}
				{
					p.SetState(249)
					p.expression(7)
				}

			case 3:
				localctx = NewDivideExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ScriptBlockParserRULE_expression)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(251)
					p.Match(ScriptBlockParserDIVIDE)
				}
				{
					p.SetState(252)
					p.expression(6)
				}

			case 4:
				localctx = NewIntegerDivideExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ScriptBlockParserRULE_expression)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(254)
					p.Match(ScriptBlockParserINTEGER_DIVIDE)
				}
				{
					p.SetState(255)
					p.expression(5)
				}

			case 5:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ScriptBlockParserRULE_expression)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(257)
					p.Match(ScriptBlockParserPLUS)
				}
				{
					p.SetState(258)
					p.expression(4)
				}

			case 6:
				localctx = NewSubtractExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ScriptBlockParserRULE_expression)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(260)
					p.Match(ScriptBlockParserSUBTRACT)
				}
				{
					p.SetState(261)
					p.expression(3)
				}

			}

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScriptBlockParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScriptBlockParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDIGITS() []antlr.TerminalNode {
	return s.GetTokens(ScriptBlockParserDIGITS)
}

func (s *NumberContext) DIGITS(i int) antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDIGITS, i)
}

func (s *NumberContext) SIGN() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserSIGN, 0)
}

func (s *NumberContext) DOT() antlr.TerminalNode {
	return s.GetToken(ScriptBlockParserDOT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScriptBlockParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *ScriptBlockParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ScriptBlockParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScriptBlockParserSIGN {
		{
			p.SetState(267)
			p.Match(ScriptBlockParserSIGN)
		}

	}
	{
		p.SetState(270)
		p.Match(ScriptBlockParserDIGITS)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(271)
			p.Match(ScriptBlockParserDOT)
		}
		{
			p.SetState(272)
			p.Match(ScriptBlockParserDIGITS)
		}

	}

	return localctx
}

func (p *ScriptBlockParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ScriptBlockParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
