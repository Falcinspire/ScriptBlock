parser grammar ScriptBlockParser;
options { tokenVocab=ScriptBlockLexer; }

unit: (importLine NEWLINES)* topDefinition* EOF;
documentation: DOC_START NEWLINES DOC_LINE* DOC_END NEWLINES;
parameterList: OPAREN (IDENTIFIER (COMMA IDENTIFIER)*)? CPAREN;
argumentList: OPAREN (expression (COMMA expression)*)? CPAREN;
structureList: OSQUARE (expression (COMMA expression)*)? CSQUARE;
nativeCall: NATIVE argumentList;
functionCall: IDENTIFIER ((argumentList functionFrame?) | functionFrame);
functionFrame: OCURLY (parameterList ARROW)? NEWLINES statement* CCURLY;
functionDefinition: documentation? tag? INTERNAL? FUNCTION IDENTIFIER functionFrame;
functionDefinitionShortcut: documentation? tag? INTERNAL? FUNCTION IDENTIFIER EQUALS functionCall;
tag: POUND IDENTIFIER COLON IDENTIFIER;
templateDefinition: documentation? INTERNAL? TEMPLATE IDENTIFIER functionFrame;
constantDefinition: documentation? INTERNAL? CONSTANT IDENTIFIER EQUALS expression;
formatter: LESS_THAN IDENTIFIER GREATER_THAN argumentList;

importLine: IMPORT IDENTIFIER COLON IDENTIFIER;

topDefinition
: constantDefinition NEWLINES #constantDefinitionTop
| functionDefinition NEWLINES #functionDefinitionTop
| templateDefinition NEWLINES #templateDefinitionTop
| functionDefinitionShortcut NEWLINES #functionShortcutTop
;

statement
: (functionCall NEWLINES) #functionCallStatement
| (nativeCall NEWLINES) #nativeCallStatement
| (delayStructure NEWLINES) #delayStructureStatement
| (raise NEWLINES) #raiseStatement
;

delayStructure
: DELAY structureList functionFrame
;

raise
: RAISE tag
;

expression
: number #numberExpr
| STRING #stringExpr
| IDENTIFIER #identifierExpr
| OPAREN expression CPAREN #parenthExpr
| formatter #formatterExpr
| <assoc=right> expression POWER expression #powerExpr
| expression MULTIPLY expression #multiplyExpr
| expression DIVIDE expression #divideExpr
| expression INTEGER_DIVIDE expression #integerDivideExpr
| expression PLUS expression #addExpr
| expression SUBTRACT expression #subtractExpr
| RUN IDENTIFIER argumentList #callExpr
;

number: (SIGN)? DIGITS (DOT DIGITS)?;