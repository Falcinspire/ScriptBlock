package valuepass

import (
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/addressbook"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/sirupsen/logrus"
)

type TopOneValueVisitor struct {
	*ast.BaseTopVisitor

	data *evaluator.EvaluateData
}

func NewOneValueVisitor(data *evaluator.EvaluateData) *TopOneValueVisitor {
	return &TopOneValueVisitor{nil, data}
}

func (visitor *TopOneValueVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
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
func (visitor *TopOneValueVisitor) VisitTemplateDefinition(definition *ast.TemplateDefinition) {
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

	addresstable := addressbook.LookupAddressTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.AddressBook)
	addresstable.Templates[definition.Name] = definition
}
func (visitor *TopOneValueVisitor) VisitClosureDefinition(definition *ast.ClosureDefinition) {
	logrus.WithFields(logrus.Fields{
		"module": visitor.data.Location.Module,
		"unit":   visitor.data.Location.Unit,
		"name":   definition.Name,
	}).Info("Evaluating closure reference")

	value := values.NewClosureReferenceValue(visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name)
	var table values.ValueTable
	//TODO closures are all internal?
	if definition.Internal {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Internal)
	} else {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Exported)
	}
	table[definition.Name] = value

	addresstable := addressbook.LookupAddressTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.AddressBook)
	addresstable.Closures[definition.Name] = definition
}

type UnitOneValueVisitor struct {
	data *evaluator.EvaluateData
}

func NewUnitOneValueVisitor(data *evaluator.EvaluateData) *UnitOneValueVisitor {
	return &UnitOneValueVisitor{data}
}
func (visitor *UnitOneValueVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(NewOneValueVisitor(visitor.data))
	}
}
