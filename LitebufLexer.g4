lexer grammar LitebufLexer;

// Define tokens
NAMESPACE: 'namespace';
IMPORT: 'import';
STRUCT: 'struct';
UNION: 'union';
ENUM: 'enum';
ROOT: 'root';
SERVICE: 'srv';
CHANNEL: 'chan';
LBRACE: '{';
RBRACE: '}';
COLON: ':';
EQUALS: '=';
RARROW: '->';
BOOLEAN: 'bool';
INT8: 'i8';
UINT8: 'u8';
STRING: 'str';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
ORDER: [0-9]+;
URL: '"' (~["])* '"';

// Skip whitespaces
WS: [ \t\r\n]+ -> skip;

