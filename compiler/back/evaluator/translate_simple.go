package evaluator

import (
	"fmt"

	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/falcinspire/scriptblock/compiler/ast/symbol"
	"github.com/falcinspire/scriptblock/compiler/back/values"
)

func TranslateFrame(frame *CallFrame, data *EvaluateData) []string {

	// TODO move or not?
	PushCallStack(frame.Name, data.CallStack)
	fmt.Println(ListCallStack(data.CallStack))

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
		data.CallStack,
		data.LoopInject,
		data.ModulePath,
		data.Output,
	}

	stringBody := make([]string, 0)
	statementVisitor := NewTranslateStatementVisitor(newData)
	for _, statement := range frame.Body {
		stringBody = append(stringBody, statementVisitor.QuickVisitStatement(statement)...)
	}

	PopCallStack(data.CallStack)

	return stringBody
}

func TranslateFunctor(definition *ast.TemplateDefinition, arguments []values.Value, location *location.UnitLocation, captures []values.Value, data *EvaluateData) []string {
	frame := &CallFrame{location, definition.Name, definition.Body, definition.Parameters, definition.Capture, arguments, captures}
	return TranslateFrame(frame, data)
}

func TranslateFunction(definition *ast.FunctionDefinition, location *location.UnitLocation, data *EvaluateData) []string {
	frame := &CallFrame{location, definition.Name, definition.Body, symbol.NoParameters(), symbol.NoCloses(), values.NoArguments(), values.NoCaptures()}
	return TranslateFrame(frame, data)
}
