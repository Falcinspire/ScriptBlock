package desugar

import (
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
	"github.com/sirupsen/logrus"
)

type DesugarExpressionVisitor struct {
	*ast.BaseExpressionVisitor

	depth         int
	freeVariables *freeVariableSet
}

func NewDesugarExpressionVisitor(depth int, freeVariables *freeVariableSet) *DesugarExpressionVisitor {
	return &DesugarExpressionVisitor{nil, depth, freeVariables}
}

func (visitor *DesugarExpressionVisitor) VisitIdentifier(identifier *ast.IdentifierExpression) {
	// capturing identifier
	// request address to the caller
	// make the identifier point to enclosed
	if identifier.Free {
		AddToFreeSet(identifier.Address.Data.(*symbols.ParameterAddress), visitor.freeVariables)
		identifier.Address = symbols.NewCaptureAddressBox(identifier.Address.Data.(*symbols.ParameterAddress).Name)
		identifier.Free = false
	}
}

type DesugarStatementVisitor struct {
	*ast.BaseStatementVisitor

	injector *FunctionInjector
	rewriter *StatementRewriter

	depth         int
	freeVariables *freeVariableSet
}

func NewDesugarStatementVisitor(injector *FunctionInjector, rewriter *StatementRewriter, depth int, freeVariables *freeVariableSet) *DesugarStatementVisitor {
	return &DesugarStatementVisitor{nil, injector, rewriter, depth, freeVariables}
}

func (visitor *DesugarStatementVisitor) VisitFunctionCall(call *ast.FunctionCall) {
	for _, argument := range call.Arguments {
		argument.Accept(NewDesugarExpressionVisitor(visitor.depth, visitor.freeVariables))
	}

	if call.Trailing != nil {
		trailing := call.Trailing
		trailingToExpression := DesugarTrailing(&FunctionFrame{trailing.Parameters, trailing.Body}, visitor.injector, visitor.depth, visitor.freeVariables)
		call.Arguments = append(call.Arguments, trailingToExpression)
		call.Trailing = nil // mark this to be removed
	}
}
func (visitor *DesugarStatementVisitor) VisitDelay(call *ast.DelayStatement) {
	logrus.WithFields(logrus.Fields{
		"line_local": visitor.rewriter.index,
	}).Info("rewriting delay call")

	trailingToExpression := DesugarTrailing(&FunctionFrame{[]string{}, call.Body}, visitor.injector, visitor.depth, visitor.freeVariables)
	call.FunctionCall = ast.NewFunctionCall(trailingToExpression, []ast.Expression{}, nil)
}

func (visitor *DesugarStatementVisitor) VisitNativeCall(call *ast.NativeCall) {
	for _, expression := range call.Arguments {
		expression.Accept(NewDesugarExpressionVisitor(visitor.depth, visitor.freeVariables))
	}
}

type DesugarTopDefinitionVisitor struct {
	*ast.BaseTopVisitor

	injector *FunctionInjector
}

func NewDesugarTopDefinitionVisitor(injector *FunctionInjector) *DesugarTopDefinitionVisitor {
	visitor := new(DesugarTopDefinitionVisitor)
	visitor.injector = injector
	return visitor
}
func (visitor *DesugarTopDefinitionVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	logrus.WithFields(logrus.Fields{
		"function": definition.Name,
	}).Info("desugaring")

	frame := &FunctionFrame{make([]string, 0), definition.Body}
	injector := visitor.injector
	depth := -1
	DesugarFunctionFrame(frame, injector, depth)
}

func (visitor *DesugarTopDefinitionVisitor) VisitFunctionShortcutDefinition(shortcut *ast.FunctionShortcutDefinition) {
	logrus.WithFields(logrus.Fields{
		"function-shortcut": shortcut.Name,
	}).Info("desugaring")

	visitor.injector.ReplaceFunction(shortcut.Name, []ast.Statement{shortcut.FunctionCall}, shortcut.Internal, shortcut.Tag, shortcut.Documentation)
}

func (visitor *DesugarTopDefinitionVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	logrus.WithFields(logrus.Fields{
		"template": definition.Name,
	}).Info("desugaring")

	frame := &FunctionFrame{definition.Parameters, definition.Body}
	injector := visitor.injector
	depth := -1
	DesugarFunctionFrame(frame, injector, depth)
}

func (visitor *DesugarTopDefinitionVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	logrus.WithFields(logrus.Fields{
		"constant": definition.Name,
	}).Info("desugaring")

	expression := definition.Value
	DesugarExpressionFrame(expression)
}

type DesugarUnitVisitor struct {
	unit     *ast.Unit
	location *location.UnitLocation
}

func NewDesugarUnitVisitor(unit *ast.Unit, location *location.UnitLocation) *DesugarUnitVisitor {
	visitor := new(DesugarUnitVisitor)
	visitor.unit = unit
	visitor.location = location
	return visitor
}
func (visitor *DesugarUnitVisitor) VisitUnit(unit *ast.Unit) {
	for i, definition := range unit.Definitions {
		definition.Accept(NewDesugarTopDefinitionVisitor(NewFunctionInjector(visitor.unit, visitor.location, i)))
	}
}
