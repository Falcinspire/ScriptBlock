package astgen

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/front/parser"
)

// PSTtoAST converts the given parse tree into an unannotated abstract syntax tree
func PSTtoAST(root parser.IUnitContext) *ast.Unit {
	visitor := newUnitConvertVisitor()
	root.EnterRule(visitor)
	return visitor.Unit
}

func convertExpression(expression parser.IExpressionContext) ast.Expression {
	visitor := newExpressionConvertVisitor()
	expression.EnterRule(visitor)
	return visitor.Expression
}

func convertStatement(statement parser.IStatementContext) ast.Statement {
	visitor := newStatementConvertVisitor()
	statement.EnterRule(visitor)
	return visitor.Statement
}

func convertArgumentList(arguments parser.IArgumentListContext) []ast.Expression {
	argumentContexts := arguments.(*parser.ArgumentListContext).AllExpression()
	expressionArguments := make([]ast.Expression, len(argumentContexts))
	for i, argumentContext := range argumentContexts {
		expressionArguments[i] = convertExpression(argumentContext)
	}
	return expressionArguments
}

func convertParameterList(ictx parser.IParameterListContext) []string {
	parameterContexts := ictx.(*parser.ParameterListContext).AllIDENTIFIER()
	parameterList := make([]string, len(parameterContexts))
	for i, parameterContext := range parameterContexts {
		parameterList[i] = parameterContext.GetText()
	}
	return parameterList
}

type functionFrame struct {
	Parameters []string
	Body       []ast.Statement
}

func convertFunctionFrame(ictx parser.IFunctionFrameContext) *functionFrame {
	ctx := ictx.(*parser.FunctionFrameContext)

	var parameterList []string
	parameterListContext := ctx.ParameterList()
	if parameterListContext != nil {
		parameterList = convertParameterList(parameterListContext)
	} else {
		parameterList = make([]string, 0)
	}

	statementContexts := ctx.AllStatement()
	body := make([]ast.Statement, len(statementContexts))
	for i, statementContext := range statementContexts {
		body[i] = convertStatement(statementContext)
	}

	return &functionFrame{parameterList, body}
}

func convertFunctionCall(ctx *parser.FunctionCallContext) *ast.FunctionCall {
	var trailingFrame *ast.TrailingFunction

	trailingContext := ctx.FunctionFrame()
	if trailingContext != nil {
		frame := convertFunctionFrame(trailingContext)
		trailingFrame = &ast.TrailingFunction{Parameters: frame.Parameters, Body: frame.Body}
	}

	var arguments []ast.Expression
	argumentListContext := ctx.ArgumentList()
	if argumentListContext != nil {
		arguments = convertArgumentList(argumentListContext)
	} else {
		arguments = make([]ast.Expression, 0)
	}

	identifierName := ctx.IDENTIFIER().GetText()
	return ast.NewFunctionCall(ast.NewIdentifierExpression(identifierName, convertTokenMetadata(ctx.IDENTIFIER())), arguments, trailingFrame, convertMetadata(ctx))
}

func convertTag(identifierCtxs []antlr.TerminalNode) ast.Tag {
	namespace := identifierCtxs[0].GetText()
	identity := identifierCtxs[1].GetText()
	return ast.Tag{Namespace: namespace, Identity: identity}
}

func convertDoc(ictx parser.IDocumentationContext) string {
	docLineContexts := ictx.(*parser.DocumentationContext).AllDOC_LINE()
	docLines := make([]string, len(docLineContexts))
	for i, docLineContext := range docLineContexts {
		text := docLineContext.GetText()
		// TODO may be unnecessary to trim
		docLines[i] = strings.TrimRight(text[3:], "\n\r")
	}
	return strings.Join(docLines, "\n")
}

func convertMetadata(node antlr.ParserRuleContext) *ast.Metadata {
	return &ast.Metadata{
		StartLine:   node.GetStart().GetLine(),
		StartColumn: node.GetStart().GetColumn(),
		EndLine:     node.GetStop().GetLine(),
		EndColumn:   node.GetStop().GetColumn() + len(node.GetStop().GetText()),
		Text:        node.GetText(),
	}
}

func convertTokenMetadata(inode antlr.TerminalNode) *ast.Metadata {
	node := inode.GetSymbol()
	return &ast.Metadata{
		StartLine:   node.GetLine(),
		StartColumn: node.GetColumn(),
		EndLine:     node.GetLine(),
		EndColumn:   node.GetColumn() + len(node.GetText()),
		Text:        node.GetText(),
	}
}
