parser grammar LitebufParser;

options {
  tokenVocab = LitebufLexer;
}

// Define parser rules
source: namespace (include)? (def | root | service | channel)* EOF;

namespace: NAMESPACE identifier;

include: IMPORT (URL | LBRACE URL+ RBRACE);

def : (struct | union | enum);

struct: STRUCT identifier LBRACE field+ RBRACE;

union: UNION identifier LBRACE field+ RBRACE;

field: identifier COLON type;

enum: ENUM identifier LBRACE variant+ RBRACE;

variant: identifier (order)?;

order: EQUALS ORDER;

root: ROOT (identifier | def);

service: SERVICE identifier LBRACE fn* RBRACE;

fn: identifier COLON (identifier | type | def) RARROW (identifier | type | def);

channel: CHANNEL identifier (identifier | type | def);

type: identifier | primitive;

primitive: BOOLEAN | INT8 | UINT8 | STRING;

identifier: IDENTIFIER;

