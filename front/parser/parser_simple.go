package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parse simply turns the file at the given location into a parse (concrete syntax) tree
func Parse(filePath string) IUnitContext {
	input, _ := antlr.NewFileStream(filePath)
	lexer := NewScriptBlockLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewScriptBlockParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	return p.Unit()
}
