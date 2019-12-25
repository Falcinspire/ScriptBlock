package evaluator

import (
	"fmt"
	"math"

	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/back/dumper"
	"github.com/falcinspire/scriptblock/compiler/back/nativefuncs"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/google/uuid"
)

type ReduceExpressionVisitor struct {
	*ast.BaseExpressionVisitor

	Result values.Value

	nativeMap map[string]nativefuncs.NativeFunction

	data *EvaluateData
}

func NewReduceExpressionVisitor(data *EvaluateData) *ReduceExpressionVisitor {
	visitor := new(ReduceExpressionVisitor)
	visitor.nativeMap = nativefuncs.CreateNativeMap()
	visitor.data = data
	return visitor
}
func (visitor *ReduceExpressionVisitor) QuickVisitExpression(expression ast.Expression) values.Value {
	expression.Accept(visitor)
	return visitor.Result
}
func (visitor *ReduceExpressionVisitor) QuickVisitArgumentList(argumentList []ast.Expression) []values.Value {
	argumentValues := make([]values.Value, len(argumentList))
	for i, argument := range argumentList {
		argumentValues[i] = visitor.QuickVisitExpression(argument)
	}
	return argumentValues
}
func (visitor *ReduceExpressionVisitor) VisitNumber(number *ast.NumberExpression) {
	visitor.Result = values.NewNumberValue(number.Value)
}
func (visitor *ReduceExpressionVisitor) VisitString(stringExpression *ast.StringExpression) {
	visitor.Result = values.NewStringValue(stringExpression.Value)
}
func (visitor *ReduceExpressionVisitor) VisitIdentifier(identifier *ast.IdentifierExpression) {
	visitor.Result = GetValueForAddress(identifier.Address, visitor.data)
}
func (visitor *ReduceExpressionVisitor) VisitFunctor(closure *ast.FunctorExpression) {
	callee := visitor.QuickVisitExpression(closure.Callee).(*values.TemplateValue)

	captureArgs := make([]values.Value, len(closure.Capture))
	for i, capture := range closure.Capture {
		captureArgs[i] = ReduceIdentifier(capture, visitor.data)
	}
	visitor.Result = values.NewFunctorValue(callee, captureArgs)
}
func (visitor *ReduceExpressionVisitor) VisitAdd(add *ast.AddExpression) {
	left := visitor.QuickVisitExpression(add.Left).(*values.NumberValue)
	right := visitor.QuickVisitExpression(add.Right).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value + right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitSubtract(add *ast.SubtractExpression) {
	left := visitor.QuickVisitExpression(add.Left).(*values.NumberValue)
	right := visitor.QuickVisitExpression(add.Right).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value - right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitMultiply(add *ast.MultiplyExpression) {
	left := visitor.QuickVisitExpression(add.Left).(*values.NumberValue)
	right := visitor.QuickVisitExpression(add.Right).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value * right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitDivide(add *ast.DivideExpression) {
	left := visitor.QuickVisitExpression(add.Left).(*values.NumberValue)
	right := visitor.QuickVisitExpression(add.Right).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value / right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitIntegerDivide(add *ast.IntegerDivideExpression) {
	left := visitor.QuickVisitExpression(add.Left).(*values.NumberValue)
	right := visitor.QuickVisitExpression(add.Right).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(float64(int(left.Value) / int(right.Value)))
}
func (visitor *ReduceExpressionVisitor) VisitPower(add *ast.PowerExpression) {
	left := visitor.QuickVisitExpression(add.Left).(*values.NumberValue)
	right := visitor.QuickVisitExpression(add.Right).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(math.Pow(left.Value, right.Value))
}
func (visitor *ReduceExpressionVisitor) VisitFormatter(formatter *ast.FormatterExpression) {
	valueArray := visitor.QuickVisitArgumentList(formatter.Arguments)
	stringResult := visitor.nativeMap[formatter.Format](visitor.data.ModulePath, valueArray)
	visitor.Result = values.NewStringValue(stringResult)
}
func (visitor *ReduceExpressionVisitor) VisitCall(call *ast.CallExpression) {
	callee := visitor.QuickVisitExpression(call.Identifier)
	arguments := visitor.QuickVisitArgumentList(call.Arguments)
	result := InvokeValue(callee, arguments, visitor.data)
	if IsLinesInvokeResult(result) {
		if len(result.Lines) == 1 {
			//TODO is this the right value?
			visitor.Result = values.NewStringValue(result.Lines[0])
		} else {
			generatedName := fmt.Sprintf("%s-%s", visitor.data.Location.Unit, uuid.New().String())
			dumper.DumpFunction(visitor.data.Location.Module, visitor.data.Location.Unit, generatedName, result.Lines, visitor.data.Output)
			visitor.Result = values.NewFunctionValue(visitor.data.Location.Module, visitor.data.Location.Unit, generatedName)
		}
	} else {
		visitor.Result = values.NewFunctionValue(result.FunctionModule, result.FunctionUnit, result.FunctionName)
	}
}
