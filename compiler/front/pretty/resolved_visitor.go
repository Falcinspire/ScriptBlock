package pretty

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/symbol"
)

type ResolvedTopDefinitionVisitor struct {
	*ast.BaseTopVisitor

	Result ResolvedTopDefinition
}

func (visitor *ResolvedTopDefinitionVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	value := QuickVisitExpression(definition.Value, &ResolvedExpressionVisitor{nil, nil})
	visitor.Result = &ResolvedConstantDefinition{"constant", definition.Name, definition.Internal, definition.Documentation, value}
}

func (visitor *ResolvedTopDefinitionVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	body := QuickVisitBody(definition.Body)
	visitor.Result = &ResolvedFunctionDefinition{"function", definition.Name, definition.Internal, definition.Documentation, body}
}

func (visitor *ResolvedTopDefinitionVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	body := QuickVisitBody(definition.Body)
	visitor.Result = &ResolvedTemplateDefinition{"template", definition.Name, definition.Internal, definition.Documentation, definition.Parameters, definition.Capture, body}
}

func QuickVisitExpressionList(expressionsAst []ast.Expression, visitor *ResolvedExpressionVisitor) []ResolvedExpression {
	expressions := make([]ResolvedExpression, len(expressionsAst))
	for i, expressionaST := range expressionsAst {
		expressionaST.Accept(visitor)
		expressions[i] = visitor.Result
	}
	return expressions
}

func QuickVisitExpression(expression ast.Expression, visitor *ResolvedExpressionVisitor) ResolvedExpression {
	expression.Accept(visitor)
	return visitor.Result
}

func QuickVisitBody(body []ast.Statement) []ResolvedStatement {
	statements := make([]ResolvedStatement, len(body))
	for i, statement := range body {
		visitor := &ResolvedStatementVisitor{nil, nil}
		statement.Accept(visitor)
		statements[i] = visitor.Result
	}
	return statements
}

type ResolvedStatementVisitor struct {
	*ast.BaseStatementVisitor

	Result ResolvedStatement
}

func (visitor *ResolvedStatementVisitor) VisitFunctionCall(functionCall *ast.FunctionCall) {
	expressionVisitor := &ResolvedExpressionVisitor{nil, nil}
	callee := QuickVisitExpression(functionCall.Callee, expressionVisitor)
	arguments := QuickVisitExpressionList(functionCall.Arguments, expressionVisitor)
	var trailing *TrailingFunction
	if functionCall.Trailing != nil {
		trailing = &TrailingFunction{"trailing-funtion", functionCall.Trailing.Parameters, QuickVisitBody(functionCall.Trailing.Body)}
	}
	visitor.Result = &ResolvedFunctionCall{"function-call", callee, arguments, trailing}
}
func (visitor *ResolvedStatementVisitor) VisitNativeCall(nativeCall *ast.NativeCall) {
	expressionVisitor := &ResolvedExpressionVisitor{nil, nil}
	arguments := QuickVisitExpressionList(nativeCall.Arguments, expressionVisitor)
	visitor.Result = &ResolvedNativeCall{"native-call", arguments}
}

type ResolvedExpressionVisitor struct {
	*ast.BaseExpressionVisitor

	Result ResolvedExpression
}

func (visitor *ResolvedExpressionVisitor) VisitIdentifier(call *ast.IdentifierExpression) {
	address := QuickVisitAddress(call.Address)
	visitor.Result = &ResolvedIdentifierExpression{"identifier", address}
}

func QuickVisitAddress(address *symbol.AddressBox) ResolvedAddress {
	switch address.Type {
	case symbol.UNIT:
		data := address.Data.(*symbol.UnitAddress)
		return &ResolvedUnitAddress{data.Module, data.Unit, data.Name}
	case symbol.PARAMETER:
		data := address.Data.(*symbol.ParameterAddress)
		return &ResolvedParameterAddress{data.ClosureDepth, data.Name}
	case symbol.CAPTURE:
		data := address.Data.(*symbol.CaptureAddress)
		return &ResolvedCaptureAddress{data.Name}
	}

	return nil
}

func (visitor *ResolvedExpressionVisitor) VisitAdd(add *ast.AddExpression) {
	visitor.Result = &ResolvedOperatorExpression{"+", QuickVisitExpression(add.Left, visitor), QuickVisitExpression(add.Right, visitor)}
}
func (visitor *ResolvedExpressionVisitor) VisitSubtract(add *ast.SubtractExpression) {
	visitor.Result = &ResolvedOperatorExpression{"-", QuickVisitExpression(add.Left, visitor), QuickVisitExpression(add.Right, visitor)}
}
func (visitor *ResolvedExpressionVisitor) VisitMultiply(add *ast.MultiplyExpression) {
	visitor.Result = &ResolvedOperatorExpression{"*", QuickVisitExpression(add.Left, visitor), QuickVisitExpression(add.Right, visitor)}
}
func (visitor *ResolvedExpressionVisitor) VisitDivide(add *ast.DivideExpression) {
	visitor.Result = &ResolvedOperatorExpression{"/", QuickVisitExpression(add.Left, visitor), QuickVisitExpression(add.Right, visitor)}
}
func (visitor *ResolvedExpressionVisitor) VisitIntegerDivide(add *ast.IntegerDivideExpression) {
	visitor.Result = &ResolvedOperatorExpression{"//", QuickVisitExpression(add.Left, visitor), QuickVisitExpression(add.Right, visitor)}
}
func (visitor *ResolvedExpressionVisitor) VisitPower(add *ast.PowerExpression) {
	visitor.Result = &ResolvedOperatorExpression{"^", QuickVisitExpression(add.Left, visitor), QuickVisitExpression(add.Right, visitor)}
}
func (visitor *ResolvedExpressionVisitor) VisitNumber(number *ast.NumberExpression) {
	visitor.Result = &ResolvedNumberExpression{"number", number.Value}
}
func (visitor *ResolvedExpressionVisitor) VisitString(number *ast.StringExpression) {
	visitor.Result = &ResolvedStringExpression{"string", number.Value}
}
