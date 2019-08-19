package valuepass

import (
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/sirupsen/logrus"
)

func ValuePassOne(unit *ast.Unit, data *evaluator.EvaluateData) {
	logrus.WithFields(logrus.Fields{
		"module": data.Location.Module,
		"unit":   data.Location.Unit,
	}).Info("Value pass 1")

	unit.Accept(NewUnitOneValueVisitor(data))
}
