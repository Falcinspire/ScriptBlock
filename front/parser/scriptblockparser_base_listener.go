// Generated from ScriptBlockParser.g4 by ANTLR 4.7.

package parser // ScriptBlockParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseScriptBlockParserListener is a complete listener for a parse tree produced by ScriptBlockParser.
type BaseScriptBlockParserListener struct{}

var _ ScriptBlockParserListener = &BaseScriptBlockParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseScriptBlockParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseScriptBlockParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseScriptBlockParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseScriptBlockParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseScriptBlockParserListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseScriptBlockParserListener) ExitUnit(ctx *UnitContext) {}

// EnterDocumentation is called when production documentation is entered.
func (s *BaseScriptBlockParserListener) EnterDocumentation(ctx *DocumentationContext) {}

// ExitDocumentation is called when production documentation is exited.
func (s *BaseScriptBlockParserListener) ExitDocumentation(ctx *DocumentationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseScriptBlockParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseScriptBlockParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseScriptBlockParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseScriptBlockParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterStructureList is called when production structureList is entered.
func (s *BaseScriptBlockParserListener) EnterStructureList(ctx *StructureListContext) {}

// ExitStructureList is called when production structureList is exited.
func (s *BaseScriptBlockParserListener) ExitStructureList(ctx *StructureListContext) {}

// EnterNativeCall is called when production nativeCall is entered.
func (s *BaseScriptBlockParserListener) EnterNativeCall(ctx *NativeCallContext) {}

// ExitNativeCall is called when production nativeCall is exited.
func (s *BaseScriptBlockParserListener) ExitNativeCall(ctx *NativeCallContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionFrame is called when production functionFrame is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionFrame(ctx *FunctionFrameContext) {}

// ExitFunctionFrame is called when production functionFrame is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionFrame(ctx *FunctionFrameContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionDefinitionShortcut is called when production functionDefinitionShortcut is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionDefinitionShortcut(ctx *FunctionDefinitionShortcutContext) {
}

// ExitFunctionDefinitionShortcut is called when production functionDefinitionShortcut is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionDefinitionShortcut(ctx *FunctionDefinitionShortcutContext) {
}

// EnterTag is called when production tag is entered.
func (s *BaseScriptBlockParserListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseScriptBlockParserListener) ExitTag(ctx *TagContext) {}

// EnterTemplateDefinition is called when production templateDefinition is entered.
func (s *BaseScriptBlockParserListener) EnterTemplateDefinition(ctx *TemplateDefinitionContext) {}

// ExitTemplateDefinition is called when production templateDefinition is exited.
func (s *BaseScriptBlockParserListener) ExitTemplateDefinition(ctx *TemplateDefinitionContext) {}

// EnterConstantDefinition is called when production constantDefinition is entered.
func (s *BaseScriptBlockParserListener) EnterConstantDefinition(ctx *ConstantDefinitionContext) {}

// ExitConstantDefinition is called when production constantDefinition is exited.
func (s *BaseScriptBlockParserListener) ExitConstantDefinition(ctx *ConstantDefinitionContext) {}

// EnterFormatter is called when production formatter is entered.
func (s *BaseScriptBlockParserListener) EnterFormatter(ctx *FormatterContext) {}

// ExitFormatter is called when production formatter is exited.
func (s *BaseScriptBlockParserListener) ExitFormatter(ctx *FormatterContext) {}

// EnterImportLine is called when production importLine is entered.
func (s *BaseScriptBlockParserListener) EnterImportLine(ctx *ImportLineContext) {}

// ExitImportLine is called when production importLine is exited.
func (s *BaseScriptBlockParserListener) ExitImportLine(ctx *ImportLineContext) {}

// EnterConstantDefinitionTop is called when production constantDefinitionTop is entered.
func (s *BaseScriptBlockParserListener) EnterConstantDefinitionTop(ctx *ConstantDefinitionTopContext) {
}

// ExitConstantDefinitionTop is called when production constantDefinitionTop is exited.
func (s *BaseScriptBlockParserListener) ExitConstantDefinitionTop(ctx *ConstantDefinitionTopContext) {}

// EnterFunctionDefinitionTop is called when production functionDefinitionTop is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionDefinitionTop(ctx *FunctionDefinitionTopContext) {
}

// ExitFunctionDefinitionTop is called when production functionDefinitionTop is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionDefinitionTop(ctx *FunctionDefinitionTopContext) {}

// EnterTemplateDefinitionTop is called when production templateDefinitionTop is entered.
func (s *BaseScriptBlockParserListener) EnterTemplateDefinitionTop(ctx *TemplateDefinitionTopContext) {
}

// ExitTemplateDefinitionTop is called when production templateDefinitionTop is exited.
func (s *BaseScriptBlockParserListener) ExitTemplateDefinitionTop(ctx *TemplateDefinitionTopContext) {}

// EnterFunctionShortcutTop is called when production functionShortcutTop is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionShortcutTop(ctx *FunctionShortcutTopContext) {}

// ExitFunctionShortcutTop is called when production functionShortcutTop is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionShortcutTop(ctx *FunctionShortcutTopContext) {}

// EnterFunctionCallStatement is called when production functionCallStatement is entered.
func (s *BaseScriptBlockParserListener) EnterFunctionCallStatement(ctx *FunctionCallStatementContext) {
}

// ExitFunctionCallStatement is called when production functionCallStatement is exited.
func (s *BaseScriptBlockParserListener) ExitFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// EnterNativeCallStatement is called when production nativeCallStatement is entered.
func (s *BaseScriptBlockParserListener) EnterNativeCallStatement(ctx *NativeCallStatementContext) {}

// ExitNativeCallStatement is called when production nativeCallStatement is exited.
func (s *BaseScriptBlockParserListener) ExitNativeCallStatement(ctx *NativeCallStatementContext) {}

// EnterDelayStructureStatement is called when production delayStructureStatement is entered.
func (s *BaseScriptBlockParserListener) EnterDelayStructureStatement(ctx *DelayStructureStatementContext) {
}

// ExitDelayStructureStatement is called when production delayStructureStatement is exited.
func (s *BaseScriptBlockParserListener) ExitDelayStructureStatement(ctx *DelayStructureStatementContext) {
}

// EnterRaiseStatement is called when production raiseStatement is entered.
func (s *BaseScriptBlockParserListener) EnterRaiseStatement(ctx *RaiseStatementContext) {}

// ExitRaiseStatement is called when production raiseStatement is exited.
func (s *BaseScriptBlockParserListener) ExitRaiseStatement(ctx *RaiseStatementContext) {}

// EnterDelayStructure is called when production delayStructure is entered.
func (s *BaseScriptBlockParserListener) EnterDelayStructure(ctx *DelayStructureContext) {}

// ExitDelayStructure is called when production delayStructure is exited.
func (s *BaseScriptBlockParserListener) ExitDelayStructure(ctx *DelayStructureContext) {}

// EnterRaise is called when production raise is entered.
func (s *BaseScriptBlockParserListener) EnterRaise(ctx *RaiseContext) {}

// ExitRaise is called when production raise is exited.
func (s *BaseScriptBlockParserListener) ExitRaise(ctx *RaiseContext) {}

// EnterStringExpr is called when production stringExpr is entered.
func (s *BaseScriptBlockParserListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production stringExpr is exited.
func (s *BaseScriptBlockParserListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterDivideExpr is called when production divideExpr is entered.
func (s *BaseScriptBlockParserListener) EnterDivideExpr(ctx *DivideExprContext) {}

// ExitDivideExpr is called when production divideExpr is exited.
func (s *BaseScriptBlockParserListener) ExitDivideExpr(ctx *DivideExprContext) {}

// EnterIntegerDivideExpr is called when production integerDivideExpr is entered.
func (s *BaseScriptBlockParserListener) EnterIntegerDivideExpr(ctx *IntegerDivideExprContext) {}

// ExitIntegerDivideExpr is called when production integerDivideExpr is exited.
func (s *BaseScriptBlockParserListener) ExitIntegerDivideExpr(ctx *IntegerDivideExprContext) {}

// EnterSubtractExpr is called when production subtractExpr is entered.
func (s *BaseScriptBlockParserListener) EnterSubtractExpr(ctx *SubtractExprContext) {}

// ExitSubtractExpr is called when production subtractExpr is exited.
func (s *BaseScriptBlockParserListener) ExitSubtractExpr(ctx *SubtractExprContext) {}

// EnterPowerExpr is called when production powerExpr is entered.
func (s *BaseScriptBlockParserListener) EnterPowerExpr(ctx *PowerExprContext) {}

// ExitPowerExpr is called when production powerExpr is exited.
func (s *BaseScriptBlockParserListener) ExitPowerExpr(ctx *PowerExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseScriptBlockParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseScriptBlockParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterNumberExpr is called when production numberExpr is entered.
func (s *BaseScriptBlockParserListener) EnterNumberExpr(ctx *NumberExprContext) {}

// ExitNumberExpr is called when production numberExpr is exited.
func (s *BaseScriptBlockParserListener) ExitNumberExpr(ctx *NumberExprContext) {}

// EnterMultiplyExpr is called when production multiplyExpr is entered.
func (s *BaseScriptBlockParserListener) EnterMultiplyExpr(ctx *MultiplyExprContext) {}

// ExitMultiplyExpr is called when production multiplyExpr is exited.
func (s *BaseScriptBlockParserListener) ExitMultiplyExpr(ctx *MultiplyExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseScriptBlockParserListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseScriptBlockParserListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterFormatterExpr is called when production formatterExpr is entered.
func (s *BaseScriptBlockParserListener) EnterFormatterExpr(ctx *FormatterExprContext) {}

// ExitFormatterExpr is called when production formatterExpr is exited.
func (s *BaseScriptBlockParserListener) ExitFormatterExpr(ctx *FormatterExprContext) {}

// EnterIdentifierExpr is called when production identifierExpr is entered.
func (s *BaseScriptBlockParserListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production identifierExpr is exited.
func (s *BaseScriptBlockParserListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterParenthExpr is called when production parenthExpr is entered.
func (s *BaseScriptBlockParserListener) EnterParenthExpr(ctx *ParenthExprContext) {}

// ExitParenthExpr is called when production parenthExpr is exited.
func (s *BaseScriptBlockParserListener) ExitParenthExpr(ctx *ParenthExprContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseScriptBlockParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseScriptBlockParserListener) ExitNumber(ctx *NumberContext) {}
