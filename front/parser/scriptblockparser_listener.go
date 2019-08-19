// Generated from ScriptBlockParser.g4 by ANTLR 4.7.

package parser // ScriptBlockParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ScriptBlockParserListener is a complete listener for a parse tree produced by ScriptBlockParser.
type ScriptBlockParserListener interface {
	antlr.ParseTreeListener

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterDocumentation is called when entering the documentation production.
	EnterDocumentation(c *DocumentationContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterStructureList is called when entering the structureList production.
	EnterStructureList(c *StructureListContext)

	// EnterNativeCall is called when entering the nativeCall production.
	EnterNativeCall(c *NativeCallContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionFrame is called when entering the functionFrame production.
	EnterFunctionFrame(c *FunctionFrameContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionDefinitionShortcut is called when entering the functionDefinitionShortcut production.
	EnterFunctionDefinitionShortcut(c *FunctionDefinitionShortcutContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterTemplateDefinition is called when entering the templateDefinition production.
	EnterTemplateDefinition(c *TemplateDefinitionContext)

	// EnterConstantDefinition is called when entering the constantDefinition production.
	EnterConstantDefinition(c *ConstantDefinitionContext)

	// EnterFormatter is called when entering the formatter production.
	EnterFormatter(c *FormatterContext)

	// EnterImportLine is called when entering the importLine production.
	EnterImportLine(c *ImportLineContext)

	// EnterConstantDefinitionTop is called when entering the constantDefinitionTop production.
	EnterConstantDefinitionTop(c *ConstantDefinitionTopContext)

	// EnterFunctionDefinitionTop is called when entering the functionDefinitionTop production.
	EnterFunctionDefinitionTop(c *FunctionDefinitionTopContext)

	// EnterTemplateDefinitionTop is called when entering the templateDefinitionTop production.
	EnterTemplateDefinitionTop(c *TemplateDefinitionTopContext)

	// EnterFunctionShortcutTop is called when entering the functionShortcutTop production.
	EnterFunctionShortcutTop(c *FunctionShortcutTopContext)

	// EnterFunctionCallStatement is called when entering the functionCallStatement production.
	EnterFunctionCallStatement(c *FunctionCallStatementContext)

	// EnterNativeCallStatement is called when entering the nativeCallStatement production.
	EnterNativeCallStatement(c *NativeCallStatementContext)

	// EnterDelayStructureStatement is called when entering the delayStructureStatement production.
	EnterDelayStructureStatement(c *DelayStructureStatementContext)

	// EnterDelayStructure is called when entering the delayStructure production.
	EnterDelayStructure(c *DelayStructureContext)

	// EnterStringExpr is called when entering the stringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterDivideExpr is called when entering the divideExpr production.
	EnterDivideExpr(c *DivideExprContext)

	// EnterIntegerDivideExpr is called when entering the integerDivideExpr production.
	EnterIntegerDivideExpr(c *IntegerDivideExprContext)

	// EnterSubtractExpr is called when entering the subtractExpr production.
	EnterSubtractExpr(c *SubtractExprContext)

	// EnterPowerExpr is called when entering the powerExpr production.
	EnterPowerExpr(c *PowerExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterNumberExpr is called when entering the numberExpr production.
	EnterNumberExpr(c *NumberExprContext)

	// EnterMultiplyExpr is called when entering the multiplyExpr production.
	EnterMultiplyExpr(c *MultiplyExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterFormatterExpr is called when entering the formatterExpr production.
	EnterFormatterExpr(c *FormatterExprContext)

	// EnterIdentifierExpr is called when entering the identifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterParenthExpr is called when entering the parenthExpr production.
	EnterParenthExpr(c *ParenthExprContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitDocumentation is called when exiting the documentation production.
	ExitDocumentation(c *DocumentationContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitStructureList is called when exiting the structureList production.
	ExitStructureList(c *StructureListContext)

	// ExitNativeCall is called when exiting the nativeCall production.
	ExitNativeCall(c *NativeCallContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionFrame is called when exiting the functionFrame production.
	ExitFunctionFrame(c *FunctionFrameContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionDefinitionShortcut is called when exiting the functionDefinitionShortcut production.
	ExitFunctionDefinitionShortcut(c *FunctionDefinitionShortcutContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitTemplateDefinition is called when exiting the templateDefinition production.
	ExitTemplateDefinition(c *TemplateDefinitionContext)

	// ExitConstantDefinition is called when exiting the constantDefinition production.
	ExitConstantDefinition(c *ConstantDefinitionContext)

	// ExitFormatter is called when exiting the formatter production.
	ExitFormatter(c *FormatterContext)

	// ExitImportLine is called when exiting the importLine production.
	ExitImportLine(c *ImportLineContext)

	// ExitConstantDefinitionTop is called when exiting the constantDefinitionTop production.
	ExitConstantDefinitionTop(c *ConstantDefinitionTopContext)

	// ExitFunctionDefinitionTop is called when exiting the functionDefinitionTop production.
	ExitFunctionDefinitionTop(c *FunctionDefinitionTopContext)

	// ExitTemplateDefinitionTop is called when exiting the templateDefinitionTop production.
	ExitTemplateDefinitionTop(c *TemplateDefinitionTopContext)

	// ExitFunctionShortcutTop is called when exiting the functionShortcutTop production.
	ExitFunctionShortcutTop(c *FunctionShortcutTopContext)

	// ExitFunctionCallStatement is called when exiting the functionCallStatement production.
	ExitFunctionCallStatement(c *FunctionCallStatementContext)

	// ExitNativeCallStatement is called when exiting the nativeCallStatement production.
	ExitNativeCallStatement(c *NativeCallStatementContext)

	// ExitDelayStructureStatement is called when exiting the delayStructureStatement production.
	ExitDelayStructureStatement(c *DelayStructureStatementContext)

	// ExitDelayStructure is called when exiting the delayStructure production.
	ExitDelayStructure(c *DelayStructureContext)

	// ExitStringExpr is called when exiting the stringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitDivideExpr is called when exiting the divideExpr production.
	ExitDivideExpr(c *DivideExprContext)

	// ExitIntegerDivideExpr is called when exiting the integerDivideExpr production.
	ExitIntegerDivideExpr(c *IntegerDivideExprContext)

	// ExitSubtractExpr is called when exiting the subtractExpr production.
	ExitSubtractExpr(c *SubtractExprContext)

	// ExitPowerExpr is called when exiting the powerExpr production.
	ExitPowerExpr(c *PowerExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitNumberExpr is called when exiting the numberExpr production.
	ExitNumberExpr(c *NumberExprContext)

	// ExitMultiplyExpr is called when exiting the multiplyExpr production.
	ExitMultiplyExpr(c *MultiplyExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitFormatterExpr is called when exiting the formatterExpr production.
	ExitFormatterExpr(c *FormatterExprContext)

	// ExitIdentifierExpr is called when exiting the identifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitParenthExpr is called when exiting the parenthExpr production.
	ExitParenthExpr(c *ParenthExprContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
