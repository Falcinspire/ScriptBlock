package valuepass

import (
	"fmt"

	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TopThreeValueVisitor struct {
	*ast.BaseTopVisitor

	data *evaluator.EvaluateData
	tags map[string]tags.LocationList
}

func NewTopThreeValueVisitor(data *evaluator.EvaluateData, tags map[string]tags.LocationList) *TopThreeValueVisitor {
	return &TopThreeValueVisitor{nil, data, tags}
}

func (visitor *TopThreeValueVisitor) VisitFunctionDefinition(definition *ast.FunctionDefinition) {
	logrus.WithFields(logrus.Fields{
		"module": visitor.data.Location.Module,
		"unit":   visitor.data.Location.Unit,
		"name":   definition.Name,
	}).Info("Evaluating function")

	stringBody := evaluator.TranslateFunction(definition, visitor.data.Location, visitor.data)
	evaluator.DumpFunction(visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name, stringBody, visitor.data.Output)
	if definition.Tag.Namespace != "" {
		informal := location.InformalTagPath(location.NewTagLocation(definition.Tag.Namespace, definition.Tag.Identity))
		visitor.tags[informal] = append(visitor.tags[informal], fmt.Sprintf("%s:%s/%s", visitor.data.Location.Module, visitor.data.Location.Unit, definition.Name))
	}
}

type UnitThreeValueVisitor struct {
	data *evaluator.EvaluateData
	tags map[string]tags.LocationList
}

func NewUnitThreeValueVisitor(data *evaluator.EvaluateData, tags map[string]tags.LocationList) *UnitThreeValueVisitor {
	return &UnitThreeValueVisitor{data, tags}
}

func (visitor *UnitThreeValueVisitor) VisitUnit(unit *ast.Unit) {
	for _, definition := range unit.Definitions {
		definition.Accept(NewTopThreeValueVisitor(visitor.data, visitor.tags))
	}
	if len(visitor.data.LoopInject.InjectBody) > 0 {
		loopFunctionName := fmt.Sprintf("loop-%s", uuid.New().String())
		loopFunctionBody := visitor.data.LoopInject.InjectBody
		evaluator.DumpFunction(visitor.data.Location.Module, visitor.data.Location.Unit, loopFunctionName, loopFunctionBody, visitor.data.Output)
		loopAddress := fmt.Sprintf("%s:%s/%s", visitor.data.Location.Module, visitor.data.Location.Unit, loopFunctionName)
		visitor.tags["minecraft:tick"] = append(visitor.tags["minecraft:tick"], loopAddress)
	}
}
