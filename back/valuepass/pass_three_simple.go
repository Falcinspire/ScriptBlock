package valuepass

import (
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/sirupsen/logrus"
)

func PassThree(unit *ast.Unit, data *evaluator.EvaluateData, tags map[string]tags.LocationList) {
	logrus.WithFields(logrus.Fields{
		"module": data.Location.Module,
		"unit":   data.Location.Unit,
	}).Info("Value pass 3")

	unit.Accept(newUnitThreeValueVisitor(data, tags))
}
