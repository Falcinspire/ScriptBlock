package desugar

import (
	"fmt"

	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/google/uuid"
)

type functionInjector struct {
	unit     *ast.Unit
	location *location.UnitLocation
	index    int
}

func newFunctionInjector(unit *ast.Unit, location *location.UnitLocation, index int) *functionInjector {
	return &functionInjector{unit, location, index}
}

func (injector *functionInjector) injectFunction(body []ast.Statement) (module, unit, name string) {
	functionName := fmt.Sprintf("%s-%s", injector.location.Unit, uuid.New().String())
	injector.unit.Definitions = append(injector.unit.Definitions, ast.NewFunctionDefinition(functionName, body, true, ast.Tag{}, "", nil))
	return injector.location.Module, injector.location.Unit, functionName
}
func (injector *functionInjector) injectTemplate(parameters []string, body []ast.Statement) (module, unit, name string) {
	templateName := fmt.Sprintf("%s-%s", injector.location.Unit, uuid.New().String())
	injector.unit.Definitions = append(injector.unit.Definitions, ast.NewTemplateDefinition(templateName, parameters, body, true, "", nil))

	return injector.location.Module, injector.location.Unit, templateName
}
func (injector *functionInjector) injectClosure(parameters []string, captures []string, body []ast.Statement) (module, unit, name string) {
	closureName := fmt.Sprintf("%s-%s", injector.location.Unit, uuid.New().String())
	injector.unit.Definitions = append(injector.unit.Definitions, ast.NewClosureDefinition(closureName, parameters, captures, body, true, nil))

	return injector.location.Module, injector.location.Unit, closureName
}

func (injector *functionInjector) replaceFunction(name string, body []ast.Statement, internal bool, tag ast.Tag, docs string) {
	function := ast.NewFunctionDefinition(name, body, internal, tag, docs, nil)
	injector.unit.Definitions[injector.index] = function
}
