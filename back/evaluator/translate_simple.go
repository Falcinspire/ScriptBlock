package evaluator

import (
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
)

func TranslateFrame(frame *CallFrame, data *EvaluateData) []string {
	mappedArguments := make(map[string]values.Value)
	for i, parameter := range frame.Parameters {
		mappedArguments[parameter] = frame.Arguments[i]
	}
	mappedCaptures := make(map[string]values.Value)
	for i, capture := range frame.Closes {
		mappedCaptures[capture] = frame.Captures[i]
	}

	newData := &EvaluateData{
		frame.Location,
		values.NewLocalValueTable(mappedArguments, mappedCaptures),
		data.ValueLibrary,
		data.AddressBook,
		data.LoopInject,
		data.ModulePath,
		data.Output,
	}

	stringBody := make([]string, 0)
	statementVisitor := NewTranslateStatementVisitor(newData)
	for _, statement := range frame.Body {
		stringBody = append(stringBody, statementVisitor.QuickVisitStatement(statement)...)
	}
	return stringBody
}

func TranslateTemplate(definition *ast.TemplateDefinition, arguments []values.Value, location *location.UnitLocation, data *EvaluateData) []string {
	frame := &CallFrame{location, definition.Name, definition.Body, definition.Parameters, symbols.NoCloses(), arguments, values.NoCaptures()}
	return TranslateFrame(frame, data)
}

func TranslateClosure(definition *ast.ClosureDefinition, arguments []values.Value, location *location.UnitLocation, captures []values.Value, data *EvaluateData) []string {
	frame := &CallFrame{location, definition.Name, definition.Body, definition.Parameters, definition.Capture, arguments, captures}
	return TranslateFrame(frame, data)
}

func TranslateFunction(definition *ast.FunctionDefinition, location *location.UnitLocation, data *EvaluateData) []string {
	frame := &CallFrame{location, definition.Name, definition.Body, symbols.NoParameters(), symbols.NoCloses(), values.NoArguments(), values.NoCaptures()}
	return TranslateFrame(frame, data)
}
