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
func (visitor *InvokeValueVisitor) VisitFunctor(functorValue *values.FunctorValue) {
	logrus.WithFields(logrus.Fields{
		"argslength":    len(visitor.arguments),
		"capturelength": len(functorValue.Capture),
		"module":        functorValue.Callee.Module,
		"unit":          functorValue.Callee.Unit,
		"name":          functorValue.Callee.Name,
	}).Info("invoking functor")

	templateReference := functorValue.Callee
	theFunctor := addressbook.AddressTemplate(templateReference.Module, templateReference.Unit, templateReference.Name, visitor.data.AddressBook)
	injectBody := TranslateFunctor(theFunctor, visitor.arguments, location.NewUnitLocation(templateReference.Module, templateReference.Unit), functorValue.Capture, visitor.data)

	visitor.Result = NewLinesInvokeResult(injectBody)
}
func (visitor *InvokeValueVisitor) VisitTemplate(templateValue *values.TemplateValue) {
	logrus.WithFields(logrus.Fields{
		"module":     templateValue.Module,
		"unit":       templateValue.Unit,
		"name":       templateValue.Name,
		"argslength": len(visitor.arguments),
	}).Info("invoking template")

	theTemplate := addressbook.AddressTemplate(templateValue.Module, templateValue.Unit, templateValue.Name, visitor.data.AddressBook)
	injectBody := TranslateFunctor(theTemplate, visitor.arguments, location.NewUnitLocation(templateValue.Module, templateValue.Unit), []values.Value{}, visitor.data) // TODO make this different location?

	visitor.Result = NewLinesInvokeResult(injectBody)
}
