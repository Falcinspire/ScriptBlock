package valuepass

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/back/addressbook"
	"github.com/falcinspire/scriptblock/compiler/back/evaluator"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/sirupsen/logrus"
)

type topOneValueVisitor struct {
	*ast.BaseTopVisitor

	data *evaluator.EvaluateData
}

func newOneValueVisitor(data *evaluator.EvaluateData) *topOneValueVisitor {
	return &topOneValueVisitor{nil, data}
}

func (visitor *topOneValueVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	logrus.WithFields(logrus.Fields{
		"module": visitor.data.Location.Module,
		"unit":   visitor.data.Location.Unit,
		"name":   definition.Name,
	}).Info("Evaluating function reference")

	value := values.NewFunctionValue(visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name)
	var table values.ValueTable
	if definition.Internal {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Internal)
	} else {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Exported)
	}
	table[definition.Name] = value
}
func (visitor *topOneValueVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
	logrus.WithFields(logrus.Fields{
		"module": visitor.data.Location.Module,
		"unit":   visitor.data.Location.Unit,
		"name":   definition.Name,
	}).Info("Evaluating template reference")

	value := values.NewTemplateValue(visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name)
	var table values.ValueTable
	if definition.Internal {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Internal)
	} else {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Exported)
	}
	table[definition.Name] = value

	addressbook.InsertFunctorAddress(definition, visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name, visitor.data.AddressBook)
}

type unitOneValueVisitor struct {
	data *evaluator.EvaluateData
}

func newUnitOneValueVisitor(data *evaluator.EvaluateData) *unitOneValueVisitor {
	return &unitOneValueVisitor{data}
}
func (visitor *unitOneValueVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(newOneValueVisitor(visitor.data))
	}
}
