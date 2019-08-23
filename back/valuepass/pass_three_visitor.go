package valuepass

import (
	"fmt"

	"github.com/falcinspire/scriptblock/back/dumper"
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/sirupsen/logrus"
)

type topThreeValueVisitor struct {
	*ast.BaseTopVisitor

	data *evaluator.EvaluateData
	tags map[string]tags.LocationList
}

func newTopThreeValueVisitor(data *evaluator.EvaluateData, tags map[string]tags.LocationList) *topThreeValueVisitor {
	return &topThreeValueVisitor{nil, data, tags}
}

func (visitor *topThreeValueVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	logrus.WithFields(logrus.Fields{
		"module": visitor.data.Location.Module,
		"unit":   visitor.data.Location.Unit,
		"name":   definition.Name,
	}).Info("Evaluating function")

	stringBody := evaluator.TranslateFunction(definition, visitor.data.Location, visitor.data)
	dumper.DumpFunction(visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name, stringBody, visitor.data.Output)
	if definition.Tag.Namespace != "" {
		informal := location.InformalTagPath(location.NewTagLocation(definition.Tag.Namespace, definition.Tag.Identity))
		visitor.tags[informal] = append(visitor.tags[informal], fmt.Sprintf("%s:%s/%s", visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name))
	}
}

type unitThreeValueVisitor struct {
	data *evaluator.EvaluateData
	tags map[string]tags.LocationList
}

func newUnitThreeValueVisitor(data *evaluator.EvaluateData, tags map[string]tags.LocationList) *unitThreeValueVisitor {
	return &unitThreeValueVisitor{data, tags}
}

func (visitor *unitThreeValueVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(newTopThreeValueVisitor(visitor.data, visitor.tags))
	}
	if len(visitor.data.LoopInject.InjectBody) > 0 {
		module, unit, name := evaluator.GenerateTickFunction(visitor.data.LoopInject, visitor.data.Location, visitor.data.Output)
		loopAddress := fmt.Sprintf("%s:%s/%s", module, unit, name)
		visitor.tags["minecraft:tick"] = append(visitor.tags["minecraft:tick"], loopAddress)
	}
}
