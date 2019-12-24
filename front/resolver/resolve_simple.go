package resolver

import (
	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/ast/symbol"
	"github.com/falcinspire/scriptblock/front/imports"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
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
	newDepth := depth + 1
	environment.linkedLocals.PushTable(MakeLocalTableFromParameters(frame.Parameters, newDepth))

	for _, statement := range frame.Body {
		statement.Accept(NewResolveStatementVisitor(environment, newDepth))
	}

	environment.linkedLocals.PopTable()
}

func ResolveExpressionFrame(expression ast.Expression, environment *ResolveEnvironment) {
	environment.linkedLocals.PushTable(symbols.NewLocalSymbolTable())

	expressionVisitor := NewResolveExpressionVisitor(environment)
	expression.Accept(expressionVisitor)

	environment.linkedLocals.PopTable()
}

func ResolveIdentifier(name string, metadata *ast.Metadata, environment *ResolveEnvironment) (address *symbol.AddressBox, free bool) {
	tryLocal, exists := FindLocal(name, environment)
	if exists {
		return tryLocal, false
	}

	tryClosed, exists := FindClosed(name, environment)
	if exists {
		return tryClosed, true
	}

	tryUnit, exists := FindUnit(name, environment)
	if exists {
		return tryUnit, false
	}

	tryImport, exists := FindImported(name, environment)
	if exists {
		return tryImport, false
	}

	panic(NewResolveError(name, metadata))
}
