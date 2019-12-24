package astbook

import (
	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/ast/location"
)

type AstBook map[string]AstUnitBook

func NewAstBook() AstBook {
	return make(AstBook)
}

type AstUnitBook map[string]*ast.Unit

func NewAstUnitBook() AstUnitBook {
	return make(AstUnitBook)
}

func InsertAst(location *location.UnitLocation, tree *ast.Unit, book AstBook) {
	_, exists := book[location.Module]
	if !exists {
		book[location.Module] = NewAstUnitBook()
	}

	book[location.Module][location.Unit] = tree
}

func LookupAst(location *location.UnitLocation, book AstBook) *ast.Unit {
	return book[location.Module][location.Unit]
}
