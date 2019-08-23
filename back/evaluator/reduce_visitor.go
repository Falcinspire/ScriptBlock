package evaluator

import (
	"fmt"
	"math"

	"github.com/falcinspire/scriptblock/back/dumper"
	"github.com/falcinspire/scriptblock/back/nativefuncs"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/ast"
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
func (visitor *ReduceExpressionVisitor) VisitNumber(number *ast.NumberExpression) {
	visitor.Result = values.NewNumberValue(number.Value)
}
func (visitor *ReduceExpressionVisitor) VisitString(stringExpression *ast.StringExpression) {
	visitor.Result = values.NewStringValue(stringExpression.Value)
}
func (visitor *ReduceExpressionVisitor) VisitIdentifier(identifier *ast.IdentifierExpression) {
	visitor.Result = GetValueForAddress(identifier.Address, visitor.data)
}
func (visitor *ReduceExpressionVisitor) VisitClosure(closure *ast.ClosureExpression) {
	callee := ReduceExpression(closure.Callee, visitor.data).(*values.ClosureReferenceValue)

	captureArgs := make([]values.Value, len(closure.Capture))
	for i, capture := range closure.Capture {
		captureArgs[i] = ReduceIdentifier(capture, visitor.data)
	}
	visitor.Result = values.NewClosureValue(callee, captureArgs)
}
func (visitor *ReduceExpressionVisitor) VisitAdd(add *ast.AddExpression) {
	left := ReduceExpression(add.Left, visitor.data).(*values.NumberValue)
	right := ReduceExpression(add.Right, visitor.data).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value + right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitSubtract(add *ast.SubtractExpression) {
	left := ReduceExpression(add.Left, visitor.data).(*values.NumberValue)
	right := ReduceExpression(add.Right, visitor.data).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value - right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitMultiply(add *ast.MultiplyExpression) {
	left := ReduceExpression(add.Left, visitor.data).(*values.NumberValue)
	right := ReduceExpression(add.Right, visitor.data).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value * right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitDivide(add *ast.DivideExpression) {
	left := ReduceExpression(add.Left, visitor.data).(*values.NumberValue)
	right := ReduceExpression(add.Right, visitor.data).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(left.Value / right.Value)
}
func (visitor *ReduceExpressionVisitor) VisitIntegerDivide(add *ast.IntegerDivideExpression) {
	left := ReduceExpression(add.Left, visitor.data).(*values.NumberValue)
	right := ReduceExpression(add.Right, visitor.data).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(float64(int(left.Value) / int(right.Value)))
}
func (visitor *ReduceExpressionVisitor) VisitPower(add *ast.PowerExpression) {
	left := ReduceExpression(add.Left, visitor.data).(*values.NumberValue)
	right := ReduceExpression(add.Right, visitor.data).(*values.NumberValue)

	visitor.Result = values.NewNumberValue(math.Pow(left.Value, right.Value))
}
func (visitor *ReduceExpressionVisitor) VisitFormatter(formatter *ast.FormatterExpression) {
	valueArray := make([]values.Value, len(formatter.Arguments))
	for i, argument := range formatter.Arguments {
		valueArray[i] = ReduceExpression(argument, visitor.data)
	}
	stringResult := visitor.nativeMap[formatter.Format](valueArray)
	visitor.Result = values.NewStringValue(stringResult)
}
func (visitor *ReduceExpressionVisitor) VisitCall(call *ast.CallExpression) {
	callee := ReduceExpression(call.Identifier, visitor.data)
	arguments := ReduceArgumentList(call.Arguments, visitor.data)
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
