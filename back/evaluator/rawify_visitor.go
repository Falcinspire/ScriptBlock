package evaluator

import (
	"fmt"

	"github.com/falcinspire/scriptblock/back/values"
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
func (visitor *RawifyValueVisitor) VisitClosure(closureValue *values.ClosureValue) {
	panic(fmt.Errorf("Cannot reduce a closure reference to a value"))

	//TODO move this elsewhere

	// theClosure := addressbook.AddressClosure(closureValue.Module, closureValue.Unit, closureValue.Name, visitor.addresbook)
	// injectBody := CompileClosure(theClosure, []values.Value{}, closureValue.Capture, visitor.unitLocation, visitor.valuelibrary, visitor.addresbook, visitor.output)

	// if len(injectBody) > 1 {
	// 	generatedName := fmt.Sprintf("%s-%s", closureValue.Unit, uuid.New().String())
	// 	DumpFunction(closureValue.Module, closureValue.Unit, generatedName, injectBody, visitor.output)
	// 	visitor.Result = fmt.Sprintf("function %s:%s/%s", closureValue.Module, closureValue.Unit, generatedName)
	// } else {
	// 	visitor.Result = injectBody[0]
	// }
}
func (visitor *RawifyValueVisitor) VisitClosureReference(closureValue *values.ClosureReferenceValue) {
	panic(fmt.Errorf("Cannot reduce a closure reference to a value"))
}
