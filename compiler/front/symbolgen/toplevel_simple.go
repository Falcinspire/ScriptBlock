package symbolgen

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/falcinspire/scriptblock/compiler/front/symbols"
)

func SymbolsPass(unit *ast.Unit, address *location.UnitLocation, symbollibrary *symbols.SymbolLibrary) {
	visitor := NewUnitSymbolVisitor(address)
	unit.Accept(visitor)
	symbols.InsertSymbolTable(address.Module, address.Unit, visitor.Exported, symbollibrary.Exported)
	symbols.InsertSymbolTable(address.Module, address.Unit, visitor.Internal, symbollibrary.Internal)
}
