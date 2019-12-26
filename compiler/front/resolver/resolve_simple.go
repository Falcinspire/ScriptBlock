package resolver

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/falcinspire/scriptblock/compiler/ast/symbol"
	"github.com/falcinspire/scriptblock/compiler/front/imports"
	"github.com/falcinspire/scriptblock/compiler/front/symbols"
)

func ResolvePass(unit *ast.Unit, unitLocation *location.UnitLocation, symbollibrary *symbols.SymbolLibrary, importbook imports.ImportBook) {
	environment := CreateResolveEnvironment(unitLocation, symbollibrary, importbook)
	unit.Accept(NewResolveUnitVisitor(environment))
}

func MakeLocalTableFromParameters(parameters []string, depth int) symbols.LocalSymbolTable {
	localTable := symbols.NewLocalSymbolTable()
	for _, parameter := range parameters {
		localTable[parameter] = symbol.NewParameterAddressBox(depth, parameter)
	}
	return localTable
}

func ResolveFunctionFrame(frame *FunctionFrame, environment *ResolveEnvironment, depth int) {
	environment.linkedLocals.PushTable(MakeLocalTableFromParameters(frame.Parameters, depth))

	for _, statement := range frame.Body {
		statement.Accept(NewResolveStatementVisitor(environment, depth))
	}

	environment.linkedLocals.PopTable()
}

func ResolveExpressionFrame(expression ast.Expression, environment *ResolveEnvironment) {
	environment.linkedLocals.PushTable(symbols.NewLocalSymbolTable())

	expressionVisitor := NewResolveExpressionVisitor(environment)
	expression.Accept(expressionVisitor)

	environment.linkedLocals.PopTable()
}

func ResolveIdentifier(name string, metadata *ast.Metadata, environment *ResolveEnvironment) (address *symbol.AddressBox) {
	tryLocal, exists := FindLocal(name, environment)
	if exists {
		return tryLocal
	}

	tryClosed, exists := FindClosed(name, environment)
	if exists {
		return tryClosed
	}

	tryUnit, exists := FindUnit(name, environment)
	if exists {
		return tryUnit
	}

	tryImport, exists := FindImported(name, environment)
	if exists {
		return tryImport
	}

	panic(NewResolveError(name, metadata))
}
