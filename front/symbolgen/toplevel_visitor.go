package symbolgen

import (
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
)

type TopDefinitionVisitor struct {
	*ast.BaseTopVisitor

	unitAddress *location.UnitLocation

	ExportedSymbols symbols.SymbolTable
	InternalSymbols symbols.SymbolTable
}

func NewTopDefinitionVisitor(address *location.UnitLocation, exported, internal symbols.SymbolTable) *TopDefinitionVisitor {
	return &TopDefinitionVisitor{nil, address, exported, internal}
}
func (visitor *TopDefinitionVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	address := symbols.NewUnitAddressBox(visitor.unitAddress.Module, visitor.unitAddress.Unit, definition.Name)
	var symboltable symbols.SymbolTable
	if definition.Internal {
		symboltable = visitor.InternalSymbols
	} else {
		symboltable = visitor.ExportedSymbols
	}
	symboltable[definition.Name] = address
}
func (visitor *TopDefinitionVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	address := symbols.NewUnitAddressBox(visitor.unitAddress.Module, visitor.unitAddress.Unit, definition.Name)
	var symboltable symbols.SymbolTable
	if definition.Internal {
		symboltable = visitor.InternalSymbols
	} else {
		symboltable = visitor.ExportedSymbols
	}
	symboltable[definition.Name] = address
}
func (visitor *TopDefinitionVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	address := symbols.NewUnitAddressBox(visitor.unitAddress.Module, visitor.unitAddress.Unit, definition.Name)
	var symboltable symbols.SymbolTable
	if definition.Internal {
		symboltable = visitor.InternalSymbols
	} else {
		symboltable = visitor.ExportedSymbols
	}
	symboltable[definition.Name] = address
}

func (visitor *TopDefinitionVisitor) VisitFunctionShortcutDefinition(shortcut *ast.FunctionShortcutDefinition) {
	address := symbols.NewUnitAddressBox(visitor.unitAddress.Module, visitor.unitAddress.Unit, shortcut.Name)
	var symboltable symbols.SymbolTable
	if shortcut.Internal {
		symboltable = visitor.InternalSymbols
	} else {
		symboltable = visitor.ExportedSymbols
	}
	symboltable[shortcut.Name] = address
}

type UnitSymbolVisitor struct {
	unitAddress *location.UnitLocation

	Exported symbols.SymbolTable
	Internal symbols.SymbolTable
}

func NewUnitSymbolVisitor(unitAddress *location.UnitLocation) *UnitSymbolVisitor {
	visitor := new(UnitSymbolVisitor)
	visitor.unitAddress = unitAddress
	visitor.Exported = symbols.NewSymbolTable()
	visitor.Internal = symbols.NewSymbolTable()
	return visitor
}

func (visitor *UnitSymbolVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definitionVisitor := NewTopDefinitionVisitor(visitor.unitAddress, visitor.Exported, visitor.Internal)
		definition.Accept(definitionVisitor)
	}
}
