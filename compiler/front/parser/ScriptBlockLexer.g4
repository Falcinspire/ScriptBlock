lexer grammar ScriptBlockLexer;

WS
: [ \t] + -> skip
;

COMMENT: '##' ~[\\\n\r]*? [\n\r]+ -> skip;

DOC_START: '/*';
DOC_END: '*/';
DOC_LINE: '**' ~[\n\r]*? [\n\r]+;

NEWLINES
: [\n\r] +
;

FUNCTION: 'func';
TEMPLATE: 'script';
NATIVE: 'command';
CONSTANT: 'const';

RUN: 'run';
RAISE: 'raise';

DELAY: 'delay';

IMPORT: 'import';

INTERNAL: 'internal';

ARROW: '->';
OPAREN: '(';
CPAREN: ')';
OCURLY: '{';
CCURLY: '}';
OSQUARE: '[';
CSQUARE: ']';
COMMA: ',';
EQUALS: '=';
SEMICOLON: ';';
COLON: ':';
POWER: '^';
MULTIPLY: '*';
DIVIDE: '/';
INTEGER_DIVIDE: '//';
PLUS: '+';
SUBTRACT: '-';
DOT: '.';
GREATER_THAN: '>';
LESS_THAN: '<';
POUND: '#';

IDENTIFIER: [a-zA-Z_]+;

SIGN:	[+-];
DIGITS: [0-9]+;

STRING: '\'' ~[']*? '\'';
