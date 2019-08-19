package resolver

import (
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
	"github.com/sirupsen/logrus"
)

type ResolveExpressionVisitor struct {
	*ast.BaseExpressionVisitor

	environment *ResolveEnvironment
}

func NewResolveExpressionVisitor(environment *ResolveEnvironment) *ResolveExpressionVisitor {
	visitor := new(ResolveExpressionVisitor)
	visitor.environment = environment
	return visitor
}
func (visitor *ResolveExpressionVisitor) VisitIdentifier(identifier *ast.IdentifierExpression) {
	address, free := ResolveIdentifier(identifier.Name, visitor.environment)
	identifier.Address = address
	identifier.Free = free

	logrus.WithFields(logrus.Fields{
		"identifier":  identifier.Name,
		"addressType": address.Type,
		"address":     address.Data,
		"free":        free,
	}).Info("resolved identifier")
}
func (visitor *ResolveExpressionVisitor) VisitAdd(add *ast.AddExpression) {
	add.Left.Accept(visitor)
	add.Right.Accept(visitor)
}
func (visitor *ResolveExpressionVisitor) VisitSubtract(add *ast.SubtractExpression) {
	add.Left.Accept(visitor)
	add.Right.Accept(visitor)
}
func (visitor *ResolveExpressionVisitor) VisitMultiply(add *ast.MultiplyExpression) {
	add.Left.Accept(visitor)
	add.Right.Accept(visitor)
}
func (visitor *ResolveExpressionVisitor) VisitDivide(add *ast.DivideExpression) {
	add.Left.Accept(visitor)
	add.Right.Accept(visitor)
}
func (visitor *ResolveExpressionVisitor) VisitIntegerDivide(add *ast.IntegerDivideExpression) {
	add.Left.Accept(visitor)
	add.Right.Accept(visitor)
}
func (visitor *ResolveExpressionVisitor) VisitPower(add *ast.PowerExpression) {
	add.Left.Accept(visitor)
	add.Right.Accept(visitor)
}
func (visitor *ResolveExpressionVisitor) VisitFormatter(format *ast.FormatterExpression) {
	for _, argument := range format.Arguments {
		argument.Accept(visitor)
	}
}
func (visitor *ResolveExpressionVisitor) VisitCall(call *ast.CallExpression) {
	call.Identifier.Accept(visitor)
	for _, argument := range call.Arguments {
		argument.Accept(visitor)
	}
}

type ResolveStatementVisitor struct {
	*ast.BaseStatementVisitor

	environment *ResolveEnvironment
	depth       int
}

func NewResolveStatementVisitor(environment *ResolveEnvironment, depth int) *ResolveStatementVisitor {
	visitor := new(ResolveStatementVisitor)
	visitor.environment = environment
	visitor.depth = depth
	return visitor
}
func (visitor *ResolveStatementVisitor) VisitFunctionCall(call *ast.FunctionCall) {
	call.Callee.Accept(NewResolveExpressionVisitor(visitor.environment))
	if call.Trailing != nil {
		frame := &FunctionFrame{call.Trailing.Parameters, call.Trailing.Body}
		environment := visitor.environment
		ResolveFunctionFrame(frame, environment, visitor.depth)
	}
}

func (visitor *ResolveStatementVisitor) VisitNativeCall(call *ast.NativeCall) {
	for _, expression := range call.Arguments {
		expression.Accept(NewResolveExpressionVisitor(visitor.environment))
	}
}

type ResolveTopDefinitionVisitor struct {
	*ast.BaseTopVisitor

	environment *ResolveEnvironment
}

func NewResolveTopDefinitionVisitor(environment *ResolveEnvironment) *ResolveTopDefinitionVisitor {
	visitor := new(ResolveTopDefinitionVisitor)
	visitor.environment = environment
	return visitor
}
func (visitor *ResolveTopDefinitionVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	visitor.environment.linkedLocals = symbols.NewLinkedSymbolTable()

	frame := &FunctionFrame{make([]string, 0), definition.Body}
	environment := visitor.environment
	depth := -1
	ResolveFunctionFrame(frame, environment, depth)
}

func (visitor *ResolveTopDefinitionVisitor) VisitFunctionShortcutDefinition(shortcut *ast.FunctionShortcutDefinition) {
	for _, argument := range shortcut.FunctionCall.Arguments {
		argument.Accept(NewResolveExpressionVisitor(visitor.environment))
	}
}

func (visitor *ResolveTopDefinitionVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	visitor.environment.linkedLocals = symbols.NewLinkedSymbolTable()

	frame := &FunctionFrame{definition.Parameters, definition.Body}
	environment := visitor.environment
	// TODO why does this work?
	depth := -1
	ResolveFunctionFrame(frame, environment, depth)
}

func (visitor *ResolveTopDefinitionVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	visitor.environment.linkedLocals = symbols.NewLinkedSymbolTable()

	expression := definition.Value
	environment := visitor.environment
	ResolveExpressionFrame(expression, environment)
}

type ResolveUnitVisitor struct {
	unit        *ast.Unit
	location    *location.UnitLocation
	environment *ResolveEnvironment
}

func NewResolveUnitVisitor(unit *ast.Unit, location *location.UnitLocation, environment *ResolveEnvironment) *ResolveUnitVisitor {
	visitor := new(ResolveUnitVisitor)
	visitor.unit = unit
	visitor.location = location
	visitor.environment = environment
	return visitor
}
func (visitor *ResolveUnitVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(NewResolveTopDefinitionVisitor(visitor.environment))
	}
}
