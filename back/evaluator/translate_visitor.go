package evaluator

import (
	"fmt"
	"strings"

	"github.com/falcinspire/scriptblock/back/values"

	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/google/uuid"
)

type TranslateStatementVisitor struct {
	Lines []string

	data *EvaluateData
}

func NewTranslateStatementVisitor(data *EvaluateData) *TranslateStatementVisitor {
	visitor := new(TranslateStatementVisitor)
	visitor.data = data
	return visitor
}
func (visitor *TranslateStatementVisitor) VisitFunctionCall(call *ast.FunctionCall) {
	invoker := ReduceExpression(call.Callee, visitor.data)
	arguments := ReduceArgumentList(call.Arguments, visitor.data)
	result := InvokeValue(invoker, arguments, visitor.data)
	if IsFunctionReferenceInvokeResult(result) {
		visitor.Lines = []string{fmt.Sprintf("function %s:%s/%s", result.FunctionModule, result.FunctionUnit, result.FunctionName)}
	} else {
		visitor.Lines = result.Lines
	}
}
func (visitor *TranslateStatementVisitor) VisitNativeCall(call *ast.NativeCall) {
	arguments := ReduceArgumentList(call.Arguments, visitor.data)

	stringArgs := make([]string, len(arguments))
	for i, argument := range arguments {
		stringArgs[i] = RawifyValue(argument)
	}
	visitor.Lines = []string{strings.Join(stringArgs, " ")}
}
func (visitor *TranslateStatementVisitor) VisitDelay(delay *ast.DelayStatement) {
	tickDelay := int(ReduceExpression(delay.TickDelay, visitor.data).(*values.NumberValue).Value)
	cloudId := uuid.New().String()
	invokeValue := ReduceExpression(ast.NewCallExpression(delay.FunctionCall.Callee, delay.FunctionCall.Arguments), visitor.data)
	translateValue := RawifyValue(invokeValue)
	testCloud := fmt.Sprintf("execute as @e[type=minecraft:area_effect_cloud,tag=%s,nbt={Age:%d}] run %s", cloudId, tickDelay, translateValue)
	summonCloud := fmt.Sprintf("execute at @p run summon minecraft:area_effect_cloud ~ ~ ~ {Tags:[\"%s\"],Duration:%d}", cloudId, tickDelay+1)
	visitor.data.LoopInject.InjectBody = append(visitor.data.LoopInject.InjectBody, testCloud)
	visitor.Lines = []string{summonCloud}
}
