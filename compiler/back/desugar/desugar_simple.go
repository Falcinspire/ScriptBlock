package desugar

import (
	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/falcinspire/scriptblock/compiler/ast/symbol"
)

type functionFrame struct {
	Parameters []string
	Body       []ast.Statement
}

// Unit removes the syntactic sugar from the tree, performing tree rewrites/injections as needed
func Unit(unit *ast.Unit, location *location.UnitLocation) {
	for i, definition := range unit.Definitions {
		definition.Accept(newDesugarTopDefinitionVisitor(newFunctionInjector(unit, location, i)))
	}
}

func desugarFunctionFrame(frame *functionFrame, injector *functionInjector, depth int) *freeVariableSet {
	freeVars := newFreeVariableSet()

	for i, statement := range frame.Body {
		statement.Accept(newDesugarStatementVisitor(injector, NewStatementRewriter(frame.Body, i), depth, freeVars))
	}

	return freeVars
}

func desugarExpressionFrame(expression ast.Expression) {
	expressionVisitor := newDesugarExpressionVisitor(0, newFreeVariableSet())
	expression.Accept(expressionVisitor)
}

type desugaredBody struct {
	Module, Unit, Name string
	PassCaptures       []*ast.IdentifierExpression
}

func desugarTrailing(frame *functionFrame, injector *functionInjector, depth int) ast.Expression {
	freeVariables := newFreeVariableSet()
	toFunction := desugarBody(frame, injector, depth, freeVariables)

	captures := toFunction.PassCaptures

	moduleName := toFunction.Module
	unitName := toFunction.Unit
	functionName := toFunction.Name
	callIdentifier := ast.NewIdentifierExpression(functionName, nil)
	callIdentifier.Address = symbol.NewUnitAddressBox(moduleName, unitName, functionName)
	var expression ast.Expression
	if len(captures) > 0 {
		expression = ast.NewFunctorExpression(callIdentifier, captures, nil)
	} else {
		expression = callIdentifier
	}
	return expression
}

func desugarBody(frame *functionFrame, injector *functionInjector, depth int, freeVariables *freeVariableSet) *desugaredBody {
	childFreeVariables := desugarFunctionFrame(frame, injector, depth)

	var module string
	var unit string
	var name string
	var closes []string
	var captures []*ast.IdentifierExpression

	if len(ListFreeSet(childFreeVariables)) == 0 && len(frame.Parameters) == 0 {
		closes = []string{}
		captures = []*ast.IdentifierExpression{}

		module, unit, name = injector.injectFunction(frame.Body)

		logrus.WithFields(logrus.Fields{
			"name": name,
			"type": "function",
		}).Info("injected")

	} else if len(ListFreeSet(childFreeVariables)) == 0 && len(frame.Parameters) > 0 {
		closes = []string{}
		captures = []*ast.IdentifierExpression{}

		module, unit, name = injector.injectFunctor(frame.Parameters, symbol.NoCloses(), frame.Body)

		logrus.WithFields(logrus.Fields{
			"name": name,
			"type": "template",
		}).Info("injected")

	} else {
		passChildFreesUp(freeVariables, childFreeVariables, depth)
		closes = makeClosesSet(childFreeVariables)
		captures = makeCaptureSet(childFreeVariables, depth)

		module, unit, name = injector.injectFunctor(frame.Parameters, closes, frame.Body)

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
		captures[i].Address = symbol.NewAddressBox(symbol.PARAMETER, enclosedVariable)
	}
	return captures
}

func isFree(enclosedVariable *symbol.ParameterAddress, depth int) bool {
	return enclosedVariable.ClosureDepth != depth
}
