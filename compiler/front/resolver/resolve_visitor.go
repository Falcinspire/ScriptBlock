package resolver

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/front/symbols"
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
	address := ResolveIdentifier(identifier.Name, identifier.Metadata, visitor.environment)
	identifier.Address = address

	logrus.WithFields(logrus.Fields{
		"identifier":  identifier.Name,
		"addressType": address.Type,
		"address":     address.Data,
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
		ResolveFunctionFrame(frame, environment, visitor.depth+1)
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
	depth := 0
	ResolveFunctionFrame(frame, environment, depth)
}

func (visitor *ResolveTopDefinitionVisitor) VisitFunctionShortcutDefinition(shortcut *ast.FunctionShortcutDefinition) {
	expressionVisitor := NewResolveExpressionVisitor(visitor.environment)
	shortcut.FunctionCall.Callee.Accept(expressionVisitor)
	for _, argument := range shortcut.FunctionCall.Arguments {
		argument.Accept(expressionVisitor)
	}
}

func (visitor *ResolveTopDefinitionVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	visitor.environment.linkedLocals = symbols.NewLinkedSymbolTable()

	frame := &FunctionFrame{definition.Parameters, definition.Body}
	environment := visitor.environment
	depth := 0
	ResolveFunctionFrame(frame, environment, depth)
}

func (visitor *ResolveTopDefinitionVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	visitor.environment.linkedLocals = symbols.NewLinkedSymbolTable()

	expression := definition.Value
	environment := visitor.environment
	ResolveExpressionFrame(expression, environment)
}

type ResolveUnitVisitor struct {
	environment *ResolveEnvironment
}

func NewResolveUnitVisitor(environment *ResolveEnvironment) *ResolveUnitVisitor {
	return &ResolveUnitVisitor{environment}
}
func (visitor *ResolveUnitVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(NewResolveTopDefinitionVisitor(visitor.environment))
	}
}
