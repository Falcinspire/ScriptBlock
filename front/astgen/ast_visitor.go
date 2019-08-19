package astgen

import (
	"strconv"

	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/parser"
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
	visitor.Expression = ast.NewNumberExpression(number)
}

// EnterStringExpr is the visitor enter method for string expressions.
func (visitor *expressionConvertVisitor) EnterStringExpr(ctx *parser.StringExprContext) {
	stringV := ctx.GetText()
	visitor.Expression = ast.NewStringExpression(stringV[1 : len(stringV)-1])

}

// EnterIdentifierExpr is the visitor enter method for identifier expressions.
func (visitor *expressionConvertVisitor) EnterIdentifierExpr(ctx *parser.IdentifierExprContext) {
	visitor.Expression = ast.NewIdentifierExpression(ctx.IDENTIFIER().GetText())

}

// EnterAddExpr is the visitor enter method for addition expressions.
func (visitor *expressionConvertVisitor) EnterAddExpr(ctx *parser.AddExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewAddExpression(left, right)
}

// EnterSubtractExpr is the visitor enter method for subtraction expressions.
func (visitor *expressionConvertVisitor) EnterSubtractExpr(ctx *parser.SubtractExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewSubtractExpression(left, right)
}

// EnterMultiplyExpr is the visitor enter method for multiply expressions.
func (visitor *expressionConvertVisitor) EnterMultiplyExpr(ctx *parser.MultiplyExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewMultiplyExpression(left, right)
}

// EnterDivideExpr is the visitor enter method for divide expressions.
func (visitor *expressionConvertVisitor) EnterDivideExpr(ctx *parser.DivideExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewDivideExpression(left, right)
}

// EnterIntegerDivideExpr is the visitor enter method for integer divide expressions.
func (visitor *expressionConvertVisitor) EnterIntegerDivideExpr(ctx *parser.IntegerDivideExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewIntegerDivideExpression(left, right)
}

// EnterPowerExpr is the visitor enter method for power expressions.
func (visitor *expressionConvertVisitor) EnterPowerExpr(ctx *parser.PowerExprContext) {
	left := convertExpression(ctx.Expression(0))
	right := convertExpression(ctx.Expression(1))

	visitor.Expression = ast.NewPowerExpression(left, right)
}

// EnterParenthExpr is the visitor enter method for parenthesis expressions.
func (visitor *expressionConvertVisitor) EnterParenthExpr(ctx *parser.ParenthExprContext) {
	visitor.Expression = convertExpression(ctx.Expression())
}

// EnterFormatterExpr is the visitor enter method for formatter expressions.
func (visitor *expressionConvertVisitor) EnterFormatterExpr(ctx *parser.FormatterExprContext) {
	formatterContext := ctx.Formatter().(*parser.FormatterContext)
	visitor.Expression = ast.NewFormatterExpression(formatterContext.IDENTIFIER().GetText(), convertArgumentList(formatterContext.ArgumentList()))
}

// EnterCallExpr is the visitor enter method for call expressions.
func (visitor *expressionConvertVisitor) EnterCallExpr(ctx *parser.CallExprContext) {
	identifier := ctx.IDENTIFIER().GetText()
	argumentList := convertArgumentList(ctx.ArgumentList())
	visitor.Expression = ast.NewCallExpression(ast.NewIdentifierExpression(identifier), argumentList)
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
	visitor.Statement = ast.NewNativeCall(arguments)
}

func (visitor *statementConvertVisitor) EnterDelayStructureStatement(ictx *parser.DelayStructureStatementContext) {
	ctx := ictx.DelayStructure().(*parser.DelayStructureContext)
	delay := newExpressionConvertVisitor().QuickVisitExpression(ctx.StructureList().(*parser.StructureListContext).Expression(0))
	body := convertFunctionFrame(ctx.FunctionFrame())
	visitor.Statement = ast.NewDelayStatement(delay, body.Body)
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
	visitor.Definition = ast.NewConstantDefinition(name, expression, internal, docs)
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
	visitor.Definition = ast.NewFunctionDefinition(name, frame.Body, internal, tag, docs)
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
	visitor.Definition = ast.NewTemplateDefinition(name, frame.Parameters, frame.Body, internal, docs)
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
	visitor.Definition = ast.NewFunctionShortcutDefinition(name, call, internal, tag, docs)
}

type unitConvertVisitor struct {
	*parser.BaseScriptBlockParserListener

	Unit *ast.Unit
}

func newUnitConvertVisitor() *unitConvertVisitor {
	return new(unitConvertVisitor)
}
func (visitor *unitConvertVisitor) EnterUnit(ctx *parser.UnitContext) {

	// do not use visitor for now, only one option

	iimportContexts := ctx.AllImportLine()
	importLines := make([]*ast.ImportLine, len(iimportContexts))
	for i, iimportContext := range iimportContexts {
		importContext := iimportContext.(*parser.ImportLineContext)
		importLines[i] = &ast.ImportLine{Module: importContext.IDENTIFIER(0).GetText(), Unit: importContext.IDENTIFIER(1).GetText()}
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
