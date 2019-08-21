package desugar

import (
	"fmt"

	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/google/uuid"
)

type FunctionInjector struct {
	unit     *ast.Unit
	location *location.UnitLocation
	index    int
}

func NewFunctionInjector(unit *ast.Unit, location *location.UnitLocation, index int) *FunctionInjector {
	return &FunctionInjector{unit, location, index}
}

func (injector *FunctionInjector) InjectFunction(body []ast.Statement) (module, unit, name string) {
	functionName := fmt.Sprintf("%s-%s", injector.location.Unit, uuid.New().String())
	injector.unit.Definitions = append(injector.unit.Definitions, ast.NewFunctionDefinition(functionName, body, true, ast.Tag{}, "", nil))
	return injector.location.Module, injector.location.Unit, functionName
}
func (injector *FunctionInjector) InjectTemplate(parameters []string, body []ast.Statement) (module, unit, name string) {
	templateName := fmt.Sprintf("%s-%s", injector.location.Unit, uuid.New().String())
	injector.unit.Definitions = append(injector.unit.Definitions, ast.NewTemplateDefinition(templateName, parameters, body, true, "", nil))

	return injector.location.Module, injector.location.Unit, templateName
}
func (injector *FunctionInjector) InjectClosure(parameters []string, captures []string, body []ast.Statement) (module, unit, name string) {
	closureName := fmt.Sprintf("%s-%s", injector.location.Unit, uuid.New().String())
	injector.unit.Definitions = append(injector.unit.Definitions, ast.NewClosureDefinition(closureName, parameters, captures, body, true, nil))

	return injector.location.Module, injector.location.Unit, closureName
}

func (injector *FunctionInjector) ReplaceFunction(name string, body []ast.Statement, internal bool, tag ast.Tag, docs string) {
	function := ast.NewFunctionDefinition(name, body, internal, tag, docs, nil)
	injector.unit.Definitions[injector.index] = function
}

// TODO why is this here?
func MakeIdentifierListFromNeededCaptures(captures []string) []*ast.IdentifierExpression {
	freeCapture := make([]*ast.IdentifierExpression, len(captures))
	for i, freeIdentifier := range captures {
		freeCapture[i] = ast.NewIdentifierExpression(freeIdentifier, nil)
	}
	return freeCapture
}
