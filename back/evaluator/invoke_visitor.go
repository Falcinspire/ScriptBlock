package evaluator

import (
	"errors"

	"github.com/falcinspire/scriptblock/ast/location"
	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/back/addressbook"

	"github.com/falcinspire/scriptblock/back/values"
)

type InvokeValueVisitor struct {
	*values.BaseValueVisitor

	Result *InvokeResult

	arguments []values.Value

	data *EvaluateData
}

func NewInvokeValueVisitor(arguments []values.Value, data *EvaluateData) *InvokeValueVisitor {
	return &InvokeValueVisitor{nil, nil, arguments, data}
}

func (visitor *InvokeValueVisitor) VisitFunction(functionValue *values.FunctionValue) {
	logrus.WithFields(logrus.Fields{
		"module": functionValue.Module,
		"unit":   functionValue.Unit,
		"name":   functionValue.Name,
	}).Info("invoking function")

	visitor.Result = NewFunctionReferenceInvokeResult(functionValue.Module, functionValue.Unit, functionValue.Name)
}
func (visitor *InvokeValueVisitor) VisitTemplate(templateValue *values.TemplateValue) {
	logrus.WithFields(logrus.Fields{
		"module":     templateValue.Module,
		"unit":       templateValue.Unit,
		"name":       templateValue.Name,
		"argslength": len(visitor.arguments),
	}).Info("invoking template")

	theTemplate := addressbook.AddressTemplate(templateValue.Module, templateValue.Unit, templateValue.Name, visitor.data.AddressBook)
	injectBody := TranslateTemplate(theTemplate, visitor.arguments, location.NewUnitLocation(templateValue.Module, templateValue.Unit), visitor.data) // TODO make this different location?

	visitor.Result = NewLinesInvokeResult(injectBody)
}
func (visitor *InvokeValueVisitor) VisitClosure(closureValue *values.ClosureValue) {
	logrus.WithFields(logrus.Fields{
		"argslength":    len(visitor.arguments),
		"capturelength": len(closureValue.Capture),
	}).Info("invoking closure")

	closureReference := closureValue.Callee
	theClosure := addressbook.AddressClosure(closureReference.Module, closureReference.Unit, closureReference.Name, visitor.data.AddressBook)
	injectBody := TranslateClosure(theClosure, visitor.arguments, location.NewUnitLocation(closureReference.Module, closureReference.Unit), closureValue.Capture, visitor.data)

	visitor.Result = NewLinesInvokeResult(injectBody)
}
func (visitor *InvokeValueVisitor) VisitClosureReference(closureReference *values.ClosureReferenceValue) {
	panic(errors.New("cannot invoke a closure function without a capture list"))
}
