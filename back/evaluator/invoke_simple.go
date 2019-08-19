package evaluator

import (
	"github.com/falcinspire/scriptblock/back/values"
)

func InvokeValue(value values.Value, arguments []values.Value, data *EvaluateData) *InvokeResult {
	visitor := NewInvokeValueVisitor(arguments, data)
	value.Accept(visitor)
	return visitor.Result
}

type InvokeResult struct {
	FunctionModule string
	FunctionUnit   string
	FunctionName   string
	Lines          []string
}

func NewFunctionReferenceInvokeResult(module, unit, name string) *InvokeResult {
	return &InvokeResult{module, unit, name, nil}
}
func NewLinesInvokeResult(lines []string) *InvokeResult {
	return &InvokeResult{"", "", "", lines}
}
func IsFunctionReferenceInvokeResult(result *InvokeResult) bool {
	return result.FunctionModule != ""
}
func IsLinesInvokeResult(result *InvokeResult) bool {
	return result.Lines != nil
}
