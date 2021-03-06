package astgen

import (
	"strconv"
	"strings"

	"github.com/falcinspire/scriptblock/compiler/ast/symbol"

	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/front/parser"
)

type expressionConvertVisitor struct {
	*parser.BaseScriptBlockParserListener
	Expression ast.Expression
}

func newExpressionConvertVisitor() *expressionConvertVisitor {
	return new(expressionConvertVisitor)
}

// EnterNumberExpr is the visitor enter method for number expressions.
func (visitor *expressionConvertVisitor) EnterNumberExpr(ctx *parser.NumberExprContext) {
	number, _ := strconv.ParseFloat(ctx.GetText(), 64)
	visitor.Expression = ast.NewNumberExpression(number, convertMetadata(ctx))
}

// EnterStringExpr is the visitor enter method for string expressions.
func (visitor *expressionConvertVisitor) EnterStringExpr(ctx *parser.StringExprContext) {
	stringV := ctx.GetText()
	visitor.Expression = ast.NewStringExpression(stringV[1:len(stringV)-1], convertMetadata(ctx))

}

// EnterIdentifierExpr is the visitor enter method for identifier expressions.
func (visitor *expressionConvertVisitor) EnterIdentifierExpr(ctx *parser.IdentifierExprContext) {
	visitor.Expression = ast.NewIdentifierExpression(ctx.IDENTIFIER().GetText(), convertMetadata(ctx))

}

// EnterAddExpr is the visitor enter method for addition expressions.
func (visitor *expressionConvertVisitor) EnterAddExpr(ctx *parser.AddExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewAddExpression(left, right, convertMetadata(ctx))
}

// EnterSubtractExpr is the visitor enter method for subtraction expressions.
func (visitor *expressionConvertVisitor) EnterSubtractExpr(ctx *parser.SubtractExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewSubtractExpression(left, right, convertMetadata(ctx))
}

// EnterMultiplyExpr is the visitor enter method for multiply expressions.
func (visitor *expressionConvertVisitor) EnterMultiplyExpr(ctx *parser.MultiplyExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewMultiplyExpression(left, right, convertMetadata(ctx))
}

// EnterDivideExpr is the visitor enter method for divide expressions.
func (visitor *expressionConvertVisitor) EnterDivideExpr(ctx *parser.DivideExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewDivideExpression(left, right, convertMetadata(ctx))
}

// EnterIntegerDivideExpr is the visitor enter method for integer divide expressions.
func (visitor *expressionConvertVisitor) EnterIntegerDivideExpr(ctx *parser.IntegerDivideExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewIntegerDivideExpression(left, right, convertMetadata(ctx))
}

// EnterPowerExpr is the visitor enter method for power expressions.
func (visitor *expressionConvertVisitor) EnterPowerExpr(ctx *parser.PowerExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewPowerExpression(left, right, convertMetadata(ctx))
}

// EnterParenthExpr is the visitor enter method for parenthesis expressions.
func (visitor *expressionConvertVisitor) EnterParenthExpr(ctx *parser.ParenthExprContext) {
	visitor.Expression = convertExpression(ctx.Expression())
}

// EnterFormatterExpr is the visitor enter method for formatter expressions.
func (visitor *expressionConvertVisitor) EnterFormatterExpr(ctx *parser.FormatterExprContext) {
	formatterContext := ctx.Formatter().(*parser.FormatterContext)
	visitor.Expression = ast.NewFormatterExpression(formatterContext.IDENTIFIER().GetText(), convertArgumentList(formatterContext.ArgumentList()), convertMetadata(ctx))
}

// EnterCallExpr is the visitor enter method for call expressions.
func (visitor *expressionConvertVisitor) EnterCallExpr(ctx *parser.CallExprContext) {
	identifier := ctx.IDENTIFIER().GetText()
	argumentList := convertArgumentList(ctx.ArgumentList())
	visitor.Expression = ast.NewCallExpression(ast.NewIdentifierExpression(identifier, convertTokenMetadata(ctx.IDENTIFIER())), argumentList, convertMetadata(ctx))
}

type statementConvertVisitor struct {
	*parser.BaseScriptBlockParserListener
	Statement ast.Statement
}

func newStatementConvertVisitor() *statementConvertVisitor {
	return new(statementConvertVisitor)
}

// EnterFunctionCallStatement is the visitor enter method for function call statements.
func (visitor *statementConvertVisitor) EnterFunctionCallStatement(ctx *parser.FunctionCallStatementContext) {
	visitor.Statement = convertFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext))
}

// EnterNativeCallStatement is the visitor enter method for native call statements.
func (visitor *statementConvertVisitor) EnterNativeCallStatement(ctx *parser.NativeCallStatementContext) {
	nativeCall := ctx.NativeCall().(*parser.NativeCallContext)
	arguments := newExpressionConvertVisitor().QuickVisitArgumentList(nativeCall.ArgumentList())
	visitor.Statement = ast.NewNativeCall(arguments, convertMetadata(ctx))
}

func (visitor *statementConvertVisitor) EnterDelayStructureStatement(ictx *parser.DelayStructureStatementContext) {
	ctx := ictx.DelayStructure().(*parser.DelayStructureContext)
	delay := newExpressionConvertVisitor().QuickVisitExpression(ctx.StructureList().(*parser.StructureListContext).Expression(0))
	body := convertFunctionFrame(ctx.FunctionFrame())
	visitor.Statement = ast.NewDelayStatement(delay, body.Body, convertMetadata(ctx))
}

// EnterFunctionCallStatement is the visitor enter method for function call statements.
func (visitor *statementConvertVisitor) EnterRaiseStatement(lctx *parser.RaiseStatementContext) {
	ctx := lctx.Raise().(*parser.RaiseContext)
	identifierContexts := ctx.Tag().(*parser.TagContext).AllIDENTIFIER()
	tag := convertTag(identifierContexts)

	visitor.Statement = ast.NewRaiseStatement(tag, convertMetadata(ctx))
}

//TODO maybe merge these?
func (visitor *expressionConvertVisitor) QuickVisitArgumentList(list parser.IArgumentListContext) []ast.Expression {
	expressionsContext := list.(*parser.ArgumentListContext).AllExpression()
	arguments := make([]ast.Expression, len(expressionsContext))
	for i, argument := range expressionsContext {
		arguments[i] = visitor.QuickVisitExpression(argument)
	}
	return arguments
}

func (visitor *expressionConvertVisitor) QuickVisitStructureList(list parser.IStructureListContext) []ast.Expression {
	expressionsContext := list.(*parser.StructureListContext).AllExpression()
	arguments := make([]ast.Expression, len(expressionsContext))
	for i, argument := range expressionsContext {
		arguments[i] = visitor.QuickVisitExpression(argument)
	}
	return arguments
}

func (visitor *expressionConvertVisitor) QuickVisitExpression(expression parser.IExpressionContext) ast.Expression {
	expression.EnterRule(visitor)
	return visitor.Expression
}

type topConvertVisitor struct {
	*parser.BaseScriptBlockParserListener

	Definition ast.TopDefinition
}

func newTopConvertVisitor() *topConvertVisitor {
	return &topConvertVisitor{nil, nil}
}

// EnterConstantDefinitionTop is the visitor enter method for constant definitions.
func (visitor *topConvertVisitor) EnterConstantDefinitionTop(ctxSwitch *parser.ConstantDefinitionTopContext) {
	ctx := ctxSwitch.ConstantDefinition().(*parser.ConstantDefinitionContext)
	name := ctx.IDENTIFIER().GetText()
	expressionVisitor := newExpressionConvertVisitor()
	ctx.Expression().EnterRule(expressionVisitor)
	expression := expressionVisitor.Expression
	internal := ctx.INTERNAL() != nil
	var docs string
	if ctx.Documentation() != nil {
		docs = convertDoc(ctx.Documentation())
	}
	visitor.Definition = ast.NewConstantDefinition(name, expression, internal, docs, convertMetadata(ctx))
}

// EnterFunctionDefinitionTop is the visitor enter method for function definitions.
func (visitor *topConvertVisitor) EnterFunctionDefinitionTop(ctxSwitch *parser.FunctionDefinitionTopContext) {
	ctx := ctxSwitch.FunctionDefinition().(*parser.FunctionDefinitionContext)
	name := ctx.IDENTIFIER().GetText()
	frame := convertFunctionFrame(ctx.FunctionFrame())
	internal := ctx.INTERNAL() != nil
	var tag ast.Tag
	if ctx.Tag() != nil {
		identifierContexts := ctx.Tag().(*parser.TagContext).AllIDENTIFIER()
		tag = convertTag(identifierContexts)
	}
	var docs string
	if ctx.Documentation() != nil {
		docs = convertDoc(ctx.Documentation())
	}
	visitor.Definition = ast.NewFunctionDefinition(name, frame.Body, internal, tag, docs, convertMetadata(ctx))
}

// EnterTemplateDefinitionTop is the visitor enter method for template definitions.
func (visitor *topConvertVisitor) EnterTemplateDefinitionTop(ctxSwitch *parser.TemplateDefinitionTopContext) {
	ctx := ctxSwitch.TemplateDefinition().(*parser.TemplateDefinitionContext)
	name := ctx.IDENTIFIER().GetText()
	frame := convertFunctionFrame(ctx.FunctionFrame())
	internal := ctx.INTERNAL() != nil
	var docs string
	if ctx.Documentation() != nil {
		docs = convertDoc(ctx.Documentation())
	}
	visitor.Definition = ast.NewTemplateDefinition(name, frame.Parameters, symbol.NoCloses(), frame.Body, internal, docs, convertMetadata(ctx))
}

// EnterFunctionShortcutTop is the visitor enter method for shortcut definitions.
func (visitor *topConvertVisitor) EnterFunctionShortcutTop(ctxSwitch *parser.FunctionShortcutTopContext) {
	ctx := ctxSwitch.FunctionDefinitionShortcut().(*parser.FunctionDefinitionShortcutContext)
	name := ctx.IDENTIFIER().GetText()
	call := convertFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext))
	internal := ctx.INTERNAL() != nil
	var tag ast.Tag
	if ctx.Tag() != nil {
		identifierContexts := ctx.Tag().(*parser.TagContext).AllIDENTIFIER()
		tag = convertTag(identifierContexts)
	}
	var docs string
	if ctx.Documentation() != nil {
		docs = convertDoc(ctx.Documentation())
	}
	visitor.Definition = ast.NewFunctionShortcutDefinition(name, call, internal, tag, docs, convertMetadata(ctx))
}

type unitConvertVisitor struct {
	*parser.BaseScriptBlockParserListener

	Unit *ast.Unit
}

func newUnitConvertVisitor() *unitConvertVisitor {
	return new(unitConvertVisitor)
}
func (visitor *unitConvertVisitor) EnterUnit(ctx *parser.UnitContext) {

	// TODO do not use visitor for now, only one option

	iimportContexts := ctx.AllImportLine()
	importLines := make([]*ast.ImportLine, len(iimportContexts))
	for i, iimportContext := range iimportContexts {
		importContext := iimportContext.(*parser.ImportLineContext)
		entireString := importContext.STRING().GetText()
		stringValue := entireString[1 : len(entireString)-1] // TODO crop quotes method? & check this in front end
		module, unit := parseImportStatement(stringValue)
		importLines[i] = &ast.ImportLine{Module: module, Unit: unit}
	}

	definitionContexts := ctx.AllTopDefinition()
	definitions := make([]ast.TopDefinition, len(definitionContexts))
	for i, definition := range definitionContexts {
		topVisitor := newTopConvertVisitor()
		definition.EnterRule(topVisitor)
		definitions[i] = topVisitor.Definition
	}
	visitor.Unit = ast.NewUnit(importLines, definitions)
}

func parseImportStatement(stringValue string) (module string, unit string) {
	data := strings.Split(stringValue, ":")
	return data[0], data[1]
}

// func parseImportStatement(stringValue string) (location, module, tag, unit string) {
// 	importValue := strings.Replace(stringValue, "/", "\\", -1)
// 	pathText := strings.Split(importValue, "\\")
// 	location = strings.Join(pathText[0:len(pathText)-2], "\\")
// 	moduleSegment := pathText[len(pathText)-2]
// 	if strings.Contains(moduleSegment, "#") {
// 		moduleSplit := strings.Split(moduleSegment, "#")
// 		module = moduleSplit[0]
// 		tag = moduleSplit[1]
// 	} else {
// 		module = moduleSegment
// 		tag = ""
// 	}
// 	unit = pathText[len(pathText)-1]
// 	return
// }
