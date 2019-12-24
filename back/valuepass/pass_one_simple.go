package valuepass

import (
	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/sirupsen/logrus"
)

func PassOne(unit *ast.Unit, data *evaluator.EvaluateData) {
	logrus.WithFields(logrus.Fields{
		"module": data.Location.Module,
		"unit":   data.Location.Unit,
	}).Info("Value pass 1")

	unit.Accept(newUnitOneValueVisitor(data))
}
