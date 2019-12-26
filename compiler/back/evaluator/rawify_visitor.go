package evaluator

import (
	"fmt"

	"github.com/falcinspire/scriptblock/compiler/back/values"
)

type RawifyValueVisitor struct {
	Result string
}

func NewRawifyValueVisitor() *RawifyValueVisitor {
	return &RawifyValueVisitor{}
}

func (visitor *RawifyValueVisitor) VisitNumber(numberValue *values.NumberValue) {
	visitor.Result = fmt.Sprintf("%.2f", numberValue.Value)
}
func (visitor *RawifyValueVisitor) VisitString(stringValue *values.StringValue) {
	visitor.Result = stringValue.Value
}
func (visitor *RawifyValueVisitor) VisitFunction(functionValue *values.FunctionValue) {
	visitor.Result = fmt.Sprintf("function %s:%s/%s", functionValue.Module, functionValue.Unit, functionValue.Name)
}
func (visitor *RawifyValueVisitor) VisitTemplate(templateValue *values.TemplateValue) {
	panic(fmt.Errorf("Cannot reduce a template reference to a value"))
}
func (visitor *RawifyValueVisitor) VisitFunctor(functorValue *values.FunctorValue) {
	panic(fmt.Errorf("Cannot reduce a closure reference to a value"))
}
