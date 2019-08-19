package valuepass

import (
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/sirupsen/logrus"
)

type TopTwoValueVisitor struct {
	*ast.BaseTopVisitor

	data *evaluator.EvaluateData
}

func NewTopTwoValueVisitor(data *evaluator.EvaluateData) *TopTwoValueVisitor {
	return &TopTwoValueVisitor{nil, data}
}

func (visitor *TopTwoValueVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
	logrus.WithFields(logrus.Fields{
		"module": visitor.data.Location.Module,
		"unit":   visitor.data.Location.Unit,
		"name":   definition.Name,
	}).Info("Evaluating constant")

	value := evaluator.ReduceExpression(definition.Value, visitor.data)
	var table values.ValueTable
	if definition.Internal {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Internal)
	} else {
		table = values.LookupValueTable(visitor.data.Location.Module, visitor.data.Location.Unit, visitor.data.ValueLibrary.Exported)
	}
	table[definition.Name] = value
}

type UnitTwoValueVisitor struct {
	data *evaluator.EvaluateData
}

func NewUnitTwoValueVisitor(data *evaluator.EvaluateData) *UnitTwoValueVisitor {
	return &UnitTwoValueVisitor{data}
}
func (visitor *UnitTwoValueVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(NewTopTwoValueVisitor(visitor.data))
	}
}
