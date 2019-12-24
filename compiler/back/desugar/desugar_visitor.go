package desugar

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/falcinspire/scriptblock/compiler/ast/symbol"
	"github.com/sirupsen/logrus"
)

type desugarExpressionVisitor struct {
	*ast.BaseExpressionVisitor

	depth         int
	freeVariables *freeVariableSet
}

func newDesugarExpressionVisitor(depth int, freeVariables *freeVariableSet) *desugarExpressionVisitor {
	return &desugarExpressionVisitor{nil, depth, freeVariables}
}

func (visitor *desugarExpressionVisitor) VisitIdentifier(identifier *ast.IdentifierExpression) {
	if identifier.Free {
		AddToFreeSet(identifier.Address.Data.(*symbol.ParameterAddress), visitor.freeVariables)
		identifier.Address = symbol.NewCaptureAddressBox(identifier.Address.Data.(*symbol.ParameterAddress).Name)
		identifier.Free = false
	}
}

type desugarStatementVisitor struct {
	*ast.BaseStatementVisitor

	injector *functionInjector
	rewriter *StatementRewriter

	depth         int
	freeVariables *freeVariableSet
}

func newDesugarStatementVisitor(injector *functionInjector, rewriter *StatementRewriter, depth int, freeVariables *freeVariableSet) *desugarStatementVisitor {
	return &desugarStatementVisitor{nil, injector, rewriter, depth, freeVariables}
}

func (visitor *desugarStatementVisitor) VisitFunctionCall(call *ast.FunctionCall) {
	for _, argument := range call.Arguments {
		argument.Accept(newDesugarExpressionVisitor(visitor.depth, visitor.freeVariables))
	}

	if call.Trailing != nil {
		trailing := call.Trailing
		trailingToExpression := desugarTrailing(&functionFrame{trailing.Parameters, trailing.Body}, visitor.injector, visitor.depth, visitor.freeVariables)
		call.Arguments = append(call.Arguments, trailingToExpression)
		call.Trailing = nil
	}
}
func (visitor *desugarStatementVisitor) VisitDelay(call *ast.DelayStatement) {
	logrus.WithFields(logrus.Fields{
		"line_local": visitor.rewriter.index,
	}).Info("rewriting delay call")

	trailingToExpression := desugarTrailing(&functionFrame{[]string{}, call.Body}, visitor.injector, visitor.depth, visitor.freeVariables)
	call.FunctionCall = ast.NewFunctionCall(trailingToExpression, []ast.Expression{}, nil, nil)
}

func (visitor *desugarStatementVisitor) VisitNativeCall(call *ast.NativeCall) {
	for _, expression := range call.Arguments {
		expression.Accept(newDesugarExpressionVisitor(visitor.depth, visitor.freeVariables))
	}
}

type desugarTopDefinitionVisitor struct {
	*ast.BaseTopVisitor

	injector *functionInjector
}

func newDesugarTopDefinitionVisitor(injector *functionInjector) *desugarTopDefinitionVisitor {
	visitor := new(desugarTopDefinitionVisitor)
	visitor.injector = injector
	return visitor
}
func (visitor *desugarTopDefinitionVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	logrus.WithFields(logrus.Fields{
		"function": definition.Name,
	}).Info("desugaring")

	frame := &functionFrame{make([]string, 0), definition.Body}
	injector := visitor.injector
	depth := -1
	desugarFunctionFrame(frame, injector, depth)
}

func (visitor *desugarTopDefinitionVisitor) VisitFunctionShortcutDefinition(shortcut *ast.FunctionShortcutDefinition) {
	logrus.WithFields(logrus.Fields{
		"function-shortcut": shortcut.Name,
	}).Info("desugaring")

	visitor.injector.replaceFunction(shortcut.Name, []ast.Statement{shortcut.FunctionCall}, shortcut.Internal, shortcut.Tag, shortcut.Documentation)
}

func (visitor *desugarTopDefinitionVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	logrus.WithFields(logrus.Fields{
		"template": definition.Name,
	}).Info("desugaring")

	frame := &functionFrame{definition.Parameters, definition.Body}
	injector := visitor.injector
	depth := -1
	desugarFunctionFrame(frame, injector, depth)
}

func (visitor *desugarTopDefinitionVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	logrus.WithFields(logrus.Fields{
		"constant": definition.Name,
	}).Info("desugaring")

	expression := definition.Value
	desugarExpressionFrame(expression)
}

type desugarUnitVisitor struct {
	unit     *ast.Unit
	location *location.UnitLocation
}

func newDesugarUnitVisitor(unit *ast.Unit, location *location.UnitLocation) *desugarUnitVisitor {
	visitor := new(desugarUnitVisitor)
	visitor.unit = unit
	visitor.location = location
	return visitor
}
func (visitor *desugarUnitVisitor) VisitUnit(unit *ast.Unit) {
	for i, definition := range unit.Definitions {
		definition.Accept(newDesugarTopDefinitionVisitor(newFunctionInjector(visitor.unit, visitor.location, i)))
	}
}
