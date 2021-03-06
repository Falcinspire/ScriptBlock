package valuepass

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/back/evaluator"
	"github.com/falcinspire/scriptblock/compiler/back/tags"
	"github.com/sirupsen/logrus"
)

func PassThree(unit *ast.Unit, data *evaluator.EvaluateData, tags map[string]tags.LocationList) {
	logrus.WithFields(logrus.Fields{
		"module": data.Location.Module,
		"unit":   data.Location.Unit,
	}).Info("Value pass 3")

	unit.Accept(newUnitThreeValueVisitor(data, tags))
}
