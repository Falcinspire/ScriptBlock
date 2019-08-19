package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/falcinspire/scriptblock/front/location"
)

// Parse simply turns the file at the given location into a parse (concrete syntax) tree
func Parse(location *location.UnitLocation) IUnitContext {
	filePath := fmt.Sprintf("%s/%s.sb", location.Module, location.Unit)
	input, _ := antlr.NewFileStream(filePath)
	lexer := NewScriptBlockLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewScriptBlockParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	return p.Unit()
}
