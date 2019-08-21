package desugar

import (
	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
)

type FunctionFrame struct {
	Parameters []string
	Body       []ast.Statement
}

func DesugarUnit(unit *ast.Unit, location *location.UnitLocation) {
	for i, definition := range unit.Definitions {
		definition.Accept(NewDesugarTopDefinitionVisitor(NewFunctionInjector(unit, location, i)))
	}
}

// TODO maybe change depth before entering function, or before any functions are called
func DesugarFunctionFrame(frame *FunctionFrame, injector *FunctionInjector, depth int) *freeVariableSet {
	freeVars := newFreeVariableSet()

	for i, statement := range frame.Body {
		statement.Accept(NewDesugarStatementVisitor(injector, NewStatementRewriter(frame.Body, i), depth+1, freeVars))
	}

	return freeVars
}

func DesugarExpressionFrame(expression ast.Expression) {
	expressionVisitor := NewDesugarExpressionVisitor(0, newFreeVariableSet())
	expression.Accept(expressionVisitor)
}

type desugaredBody struct {
	Module, Unit, Name string
	PassCaptures       []*ast.IdentifierExpression
}

func DesugarTrailing(frame *FunctionFrame, injector *FunctionInjector, depth int, freeVariables *freeVariableSet) ast.Expression {
	toFunction := DesugarBody(frame, injector, depth, freeVariables)

	captures := toFunction.PassCaptures
	for _, aCapture := range captures {
		aCapture.Accept(NewDesugarExpressionVisitor(depth, freeVariables))
	}

	moduleName := toFunction.Module
	unitName := toFunction.Unit
	functionName := toFunction.Name
	callIdentifier := ast.NewIdentifierExpression(functionName, nil)
	callIdentifier.Address = symbols.NewUnitAddressBox(moduleName, unitName, functionName)
	var expression ast.Expression
	if len(captures) > 0 {
		expression = ast.NewClosureExpression(callIdentifier, captures, nil)
	} else {
		expression = callIdentifier
	}
	return expression
}

func DesugarBody(frame *FunctionFrame, injector *FunctionInjector, depth int, freeVariables *freeVariableSet) *desugaredBody {
	childFreeVariables := DesugarFunctionFrame(frame, injector, depth)

	var module string
	var unit string
	var name string
	var closes []string
	var captures []*ast.IdentifierExpression

	if len(ListFreeSet(childFreeVariables)) == 0 && len(frame.Parameters) == 0 {
		closes = []string{}
		captures = []*ast.IdentifierExpression{}

		module, unit, name = injector.InjectFunction(frame.Body)

		logrus.WithFields(logrus.Fields{
			"name": name,
			"type": "function",
		}).Info("injected")

	} else if len(ListFreeSet(childFreeVariables)) == 0 && len(frame.Parameters) > 0 {
		closes = []string{}
		captures = []*ast.IdentifierExpression{}

		module, unit, name = injector.InjectTemplate(frame.Parameters, frame.Body)

		logrus.WithFields(logrus.Fields{
			"name": name,
			"type": "template",
		}).Info("injected")

	} else {
		passChildFreesUp(freeVariables, childFreeVariables, depth)
		closes = makeClosesSet(childFreeVariables)
		captures = makeCaptureSet(childFreeVariables, depth)

		module, unit, name = injector.InjectClosure(frame.Parameters, closes, frame.Body)

		logrus.WithFields(logrus.Fields{
			"name": name,
			"type": "closure",
		}).Info("injected")
	}

	return &desugaredBody{module, unit, name, captures}
}

func passChildFreesUp(parent, child *freeVariableSet, depth int) {
	for _, enclosed := range ListFreeSet(child) {
		if isFree(enclosed, depth) {
			AddToFreeSet(enclosed, parent)
		}
	}
}

func makeClosesSet(child *freeVariableSet) []string {
	closes := make([]string, len(ListFreeSet(child)))
	for i, freeAddress := range ListFreeSet(child) {
		closes[i] = freeAddress.Name
	}
	return closes
}

func makeCaptureSet(child *freeVariableSet, depth int) []*ast.IdentifierExpression {
	captures := make([]*ast.IdentifierExpression, len(ListFreeSet(child)))
	for i, enclosedVariable := range ListFreeSet(child) {
		captures[i] = ast.NewIdentifierExpression(enclosedVariable.Name, nil)
		captures[i].Address = symbols.NewAddressBox(symbols.PARAMETER, enclosedVariable)
		if enclosedVariable.ClosureDepth != depth {
			captures[i].Free = true
		} else {
			captures[i].Free = false
		}
	}
	return captures
}

func isFree(enclosedVariable *symbols.ParameterAddress, depth int) bool {
	return enclosedVariable.ClosureDepth != depth
}
