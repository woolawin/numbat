grammar Numbat;

prog: (WS | NEWLINE | object)*;

expr_bool: ('true' | 'false');
expr_num: NUMBER(UNIT)?;
expr_hex: HEX;
expr_str: STRING;
expr_null: 'null';

expr_constant: (expr_bool | expr_num | expr_hex | expr_str | expr_null);
expr_var: NON_TYPE_NAME;
expr_call: call;
expr_all: (expr_var | expr_constant | expr_call);

param_expr: '=' expr_constant;
param_type: type;
param: NON_TYPE_NAME param_type param_expr?;


type_out: TYPE_NAME;
type_in: '(' param (',' param)* ')';
type: (type_in? type_out | type_in type_out?);

proc_body: 'do' (WS | NEWLINE | statement)* 'end';
proc_name: NON_TYPE_NAME;
proc_type: type;
proc_def: 'proc' proc_name proc_type?;
proc: proc_def proc_body;


call_arg: (expr_var | expr_constant);
call_args: expr_all(',' expr_all)*;
call_secondary: NON_TYPE_NAME;
call_primary: (TYPE_NAME | NON_TYPE_NAME);
call: call_primary(':'call_secondary)?'('call_args?')';
call_stmt: call;

var_expr: '=' expr_all;
var_type: type;
var_name: NON_TYPE_NAME;
var_stmt: 'var' var_name var_type? var_expr?;

assignment_expr: '=' expr_all (',' expr_all)?;
assignment_var: NON_TYPE_NAME;
assignment: assignment_var (',' assignment_var)? assignment_expr;

return_expr: expr_all;
return: 'return' return_expr?;
return_stmt: return;

program: 'program' proc_body;

statement: (call_stmt | var_stmt | return_stmt | assignment);
object: (program | proc);

fragment UNIT_NAME: [a-zA-Z]+[a-zA-Z0-9]*;
fragment UNIT_POWER: '^'[23];

UNIT            : '#'UNIT_NAME(UNIT_POWER('/'UNIT_NAME)? | '/'UNIT_NAME(UNIT_POWER)?)?;
NEWLINE         : [\r\n]+;
NUMBER          : '-'? [0-9]+ ('.' [0-9]+)? ([e] [+-]? [0-9]+)?;
HEX             : '-'? '0' [x] [0-9A-F]+;
STRING          : '"' (~["\\] | '\\' .)* '"';
TYPE_NAME       : [A-Z][a-zA-Z0-9]+;
NON_TYPE_NAME   : [a-z][a-z_0-9]*;
WS              : [ \t]+ -> skip;
COMMENT         : '/*' .*? '*/' -> skip;
LINE_COMMENT    : '//' ~[\r\n]* -> skip;
