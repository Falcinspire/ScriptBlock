package valuepass

import (
	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/sirupsen/logrus"
)

type topTwoValueVisitor struct {
	*ast.BaseTopVisitor

	data *evaluator.EvaluateData
}

func newTopTwoValueVisitor(data *evaluator.EvaluateData) *topTwoValueVisitor {
	return &topTwoValueVisitor{nil, data}
}

func (visitor *topTwoValueVisitor) VisitConstantDefinition(definition *ast.ConstantDefinition) {
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

type unitTwoValueVisitor struct {
	data *evaluator.EvaluateData
}

func newUnitTwoValueVisitor(data *evaluator.EvaluateData) *unitTwoValueVisitor {
	return &unitTwoValueVisitor{data}
}
func (visitor *unitTwoValueVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(newTopTwoValueVisitor(visitor.data))
	}
}
