# litebuf

1. Msg serialzation and deserialization like Serde/flatbuf
2. Storage like ROSBag
3. KeyValue Server like REDIS
4. RPC like GRPC
5. PubSub like REDIS pubsub

```
source         = seq(namespace, optional(import), optional(_statements);

namespace      = seq("namespace", path);

import         = choice(_import_singular, _import_multiple);
_import_singular = seq("import", path);
_import_multiple = seq("import", "{", repeat(path), "}");

_statements    = choice(_def, root, service, channel);

root           = seq("root", ref);

service        = seq("srv", "{", repeat(_fn), "}");
_fn            = seq(identifier, ":", ref, "->", ref);

channel        = seq("chan", identifier, ref);

ref            = choice(type, _def);
type           = choice(scalar, array);
scalar         = choice(identifier, array)
array          = seq("[", scalar, "]");
primitive      = choice("bool", "i8", "u8", "string");

_def           = choice(struct, union, enum);

struct         = seq("struct", "{", repeat(_field), "}");
union          = seq("union", "{", repeat(_field), "}");
_field         = seq(identifier, ":", type, optional(default), optional(tags));
default        = seq("=", literal);
tags           = seq("@", identifier);

enum           = seq("enum", "{", repeat(_variant), "}");
_variant       = seq(identifier, optional(_order));
_order         = seq("=", digit);

path           = /"[^"]*"/;
identifier     = /[a-zA-Z_][a-zA-Z0-9_]*/;
literal        = choice(string, bool, number);
string         = /"[^"]*"/;
bool           = choice("true", "false");
number         = /-?[0-9]+(\.[0-9]+)?/;
digit          = /[0-9]+/;
```

```
grammar Litebuf;

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


// Skip whitespaces
WS: [ \t\r\n]+ -> skip;
```
