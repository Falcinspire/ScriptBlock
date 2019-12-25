package evaluator

import (
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/compiler/back/addressbook"

	"github.com/falcinspire/scriptblock/compiler/back/values"
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
func (visitor *InvokeValueVisitor) VisitFunctor(closureValue *values.FunctorValue) {
	logrus.WithFields(logrus.Fields{
		"argslength":    len(visitor.arguments),
		"capturelength": len(closureValue.Capture),
	}).Info("invoking closure")

	closureReference := closureValue.Callee
	theFunctor := addressbook.AddressFunctor(closureReference.Module, closureReference.Unit, closureReference.Name, visitor.data.AddressBook)
	injectBody := TranslateFunctor(theFunctor, visitor.arguments, location.NewUnitLocation(closureReference.Module, closureReference.Unit), closureValue.Capture, visitor.data)

	visitor.Result = NewLinesInvokeResult(injectBody)
}
func (visitor *InvokeValueVisitor) VisitTemplate(closureReference *values.TemplateValue) {
	logrus.WithFields(logrus.Fields{
		"module":     closureReference.Module,
		"unit":       closureReference.Unit,
		"name":       closureReference.Name,
		"argslength": len(visitor.arguments),
	}).Info("invoking template")

	theTemplate := addressbook.AddressFunctor(closureReference.Module, closureReference.Unit, closureReference.Name, visitor.data.AddressBook)
	injectBody := TranslateFunctor(theTemplate, visitor.arguments, location.NewUnitLocation(closureReference.Module, closureReference.Unit), []values.Value{}, visitor.data) // TODO make this different location?

	visitor.Result = NewLinesInvokeResult(injectBody)
}
