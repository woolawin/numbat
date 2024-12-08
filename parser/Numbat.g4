grammar Numbat;

prog: (WS | NEWLINE | object)*;


expr_bool: ('true' | 'false');
expr_num: NUMBER;
expr_hex: HEX;
expr_str: STRING;
expr_null: 'null';

expr_constant: (expr_bool | expr_num | expr_hex | expr_str | expr_null);
expr_var: NON_TYPE_NAME;
expr_call: call;
expr_all: (expr_var | expr_constant | expr_call);

param_expr: '=' expr_constant;
param_type: type;
param: NON_TYPE_NAME param_type? param_expr?;


type_out: TYPE_NAME;
type_in: '(' param (',' param)* ')';
type: (type_in? type_out | type_in type_out?);

proc_body: 'do' (WS | NEWLINE | statement)* 'end';
proc_name: NON_TYPE_NAME;
proc_type: type;
proc_def: 'proc' proc_name proc_type?;
proc: proc_def proc_body;


call_expr: ':' (expr_var | expr_constant)+;
call_secondary: NON_TYPE_NAME;
call_primary: (TYPE_NAME | NON_TYPE_NAME);
call: '(' call_primary call_secondary? call_expr?')';
call_stmt: call;

let_expr: '=' expr_all;
let_var_type: type;
let_var_name: NON_TYPE_NAME;
let: 'let' let_var_name let_var_type? let_expr?;

assignment_expr: '=' expr_all (',' expr_all)?;
assignment_var: NON_TYPE_NAME;
assignment: assignment_var (',' assignment_var)? assignment_expr;

return_expr: expr_all;
return: 'return' return_expr?;
return_stmt: return;

program: 'program' proc_body;

statement: (call_stmt | let | return_stmt | assignment);
object: (program | proc);


NEWLINE         : [\r\n]+;
NUMBER          : '-'? [0-9]+ ('.' [0-9]+)? ([e] [+-]? [0-9]+)?;
HEX             : '-'? '0' [x] [0-9A-F]+;
STRING          : '"' (~["\\] | '\\' .)* '"';
TYPE_NAME       : [A-Z][a-zA-Z_0-9]+;
NON_TYPE_NAME   : [a-z][a-z_0-9]+;
WS              : [ \t]+ -> skip;
COMMENT         : '/*' .*? '*/' -> skip;
LINE_COMMENT    : '//' ~[\r\n]* -> skip;
