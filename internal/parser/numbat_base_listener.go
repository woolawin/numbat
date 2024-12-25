// Code generated from Numbat.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Numbat

import "github.com/antlr4-go/antlr/v4"

// BaseNumbatListener is a complete listener for a parse tree produced by NumbatParser.
type BaseNumbatListener struct{}

var _ NumbatListener = &BaseNumbatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNumbatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNumbatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNumbatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNumbatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseNumbatListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseNumbatListener) ExitProg(ctx *ProgContext) {}

// EnterExpr_bool is called when production expr_bool is entered.
func (s *BaseNumbatListener) EnterExpr_bool(ctx *Expr_boolContext) {}

// ExitExpr_bool is called when production expr_bool is exited.
func (s *BaseNumbatListener) ExitExpr_bool(ctx *Expr_boolContext) {}

// EnterExpr_num is called when production expr_num is entered.
func (s *BaseNumbatListener) EnterExpr_num(ctx *Expr_numContext) {}

// ExitExpr_num is called when production expr_num is exited.
func (s *BaseNumbatListener) ExitExpr_num(ctx *Expr_numContext) {}

// EnterExpr_hex is called when production expr_hex is entered.
func (s *BaseNumbatListener) EnterExpr_hex(ctx *Expr_hexContext) {}

// ExitExpr_hex is called when production expr_hex is exited.
func (s *BaseNumbatListener) ExitExpr_hex(ctx *Expr_hexContext) {}

// EnterExpr_str is called when production expr_str is entered.
func (s *BaseNumbatListener) EnterExpr_str(ctx *Expr_strContext) {}

// ExitExpr_str is called when production expr_str is exited.
func (s *BaseNumbatListener) ExitExpr_str(ctx *Expr_strContext) {}

// EnterExpr_null is called when production expr_null is entered.
func (s *BaseNumbatListener) EnterExpr_null(ctx *Expr_nullContext) {}

// ExitExpr_null is called when production expr_null is exited.
func (s *BaseNumbatListener) ExitExpr_null(ctx *Expr_nullContext) {}

// EnterExpr_constant is called when production expr_constant is entered.
func (s *BaseNumbatListener) EnterExpr_constant(ctx *Expr_constantContext) {}

// ExitExpr_constant is called when production expr_constant is exited.
func (s *BaseNumbatListener) ExitExpr_constant(ctx *Expr_constantContext) {}

// EnterExpr_var is called when production expr_var is entered.
func (s *BaseNumbatListener) EnterExpr_var(ctx *Expr_varContext) {}

// ExitExpr_var is called when production expr_var is exited.
func (s *BaseNumbatListener) ExitExpr_var(ctx *Expr_varContext) {}

// EnterExpr_call is called when production expr_call is entered.
func (s *BaseNumbatListener) EnterExpr_call(ctx *Expr_callContext) {}

// ExitExpr_call is called when production expr_call is exited.
func (s *BaseNumbatListener) ExitExpr_call(ctx *Expr_callContext) {}

// EnterExpr_seq is called when production expr_seq is entered.
func (s *BaseNumbatListener) EnterExpr_seq(ctx *Expr_seqContext) {}

// ExitExpr_seq is called when production expr_seq is exited.
func (s *BaseNumbatListener) ExitExpr_seq(ctx *Expr_seqContext) {}

// EnterExpr_all is called when production expr_all is entered.
func (s *BaseNumbatListener) EnterExpr_all(ctx *Expr_allContext) {}

// ExitExpr_all is called when production expr_all is exited.
func (s *BaseNumbatListener) ExitExpr_all(ctx *Expr_allContext) {}

// EnterParam_expr is called when production param_expr is entered.
func (s *BaseNumbatListener) EnterParam_expr(ctx *Param_exprContext) {}

// ExitParam_expr is called when production param_expr is exited.
func (s *BaseNumbatListener) ExitParam_expr(ctx *Param_exprContext) {}

// EnterParam_type is called when production param_type is entered.
func (s *BaseNumbatListener) EnterParam_type(ctx *Param_typeContext) {}

// ExitParam_type is called when production param_type is exited.
func (s *BaseNumbatListener) ExitParam_type(ctx *Param_typeContext) {}

// EnterParam is called when production param is entered.
func (s *BaseNumbatListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseNumbatListener) ExitParam(ctx *ParamContext) {}

// EnterType_super_atomic_seq is called when production type_super_atomic_seq is entered.
func (s *BaseNumbatListener) EnterType_super_atomic_seq(ctx *Type_super_atomic_seqContext) {}

// ExitType_super_atomic_seq is called when production type_super_atomic_seq is exited.
func (s *BaseNumbatListener) ExitType_super_atomic_seq(ctx *Type_super_atomic_seqContext) {}

// EnterType_super_atomic is called when production type_super_atomic is entered.
func (s *BaseNumbatListener) EnterType_super_atomic(ctx *Type_super_atomicContext) {}

// ExitType_super_atomic is called when production type_super_atomic is exited.
func (s *BaseNumbatListener) ExitType_super_atomic(ctx *Type_super_atomicContext) {}

// EnterType_out is called when production type_out is entered.
func (s *BaseNumbatListener) EnterType_out(ctx *Type_outContext) {}

// ExitType_out is called when production type_out is exited.
func (s *BaseNumbatListener) ExitType_out(ctx *Type_outContext) {}

// EnterType_in is called when production type_in is entered.
func (s *BaseNumbatListener) EnterType_in(ctx *Type_inContext) {}

// ExitType_in is called when production type_in is exited.
func (s *BaseNumbatListener) ExitType_in(ctx *Type_inContext) {}

// EnterType is called when production type is entered.
func (s *BaseNumbatListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseNumbatListener) ExitType(ctx *TypeContext) {}

// EnterProc_body is called when production proc_body is entered.
func (s *BaseNumbatListener) EnterProc_body(ctx *Proc_bodyContext) {}

// ExitProc_body is called when production proc_body is exited.
func (s *BaseNumbatListener) ExitProc_body(ctx *Proc_bodyContext) {}

// EnterProc_name is called when production proc_name is entered.
func (s *BaseNumbatListener) EnterProc_name(ctx *Proc_nameContext) {}

// ExitProc_name is called when production proc_name is exited.
func (s *BaseNumbatListener) ExitProc_name(ctx *Proc_nameContext) {}

// EnterProc_type is called when production proc_type is entered.
func (s *BaseNumbatListener) EnterProc_type(ctx *Proc_typeContext) {}

// ExitProc_type is called when production proc_type is exited.
func (s *BaseNumbatListener) ExitProc_type(ctx *Proc_typeContext) {}

// EnterProc_def is called when production proc_def is entered.
func (s *BaseNumbatListener) EnterProc_def(ctx *Proc_defContext) {}

// ExitProc_def is called when production proc_def is exited.
func (s *BaseNumbatListener) ExitProc_def(ctx *Proc_defContext) {}

// EnterProc is called when production proc is entered.
func (s *BaseNumbatListener) EnterProc(ctx *ProcContext) {}

// ExitProc is called when production proc is exited.
func (s *BaseNumbatListener) ExitProc(ctx *ProcContext) {}

// EnterCall_arg is called when production call_arg is entered.
func (s *BaseNumbatListener) EnterCall_arg(ctx *Call_argContext) {}

// ExitCall_arg is called when production call_arg is exited.
func (s *BaseNumbatListener) ExitCall_arg(ctx *Call_argContext) {}

// EnterCall_args is called when production call_args is entered.
func (s *BaseNumbatListener) EnterCall_args(ctx *Call_argsContext) {}

// ExitCall_args is called when production call_args is exited.
func (s *BaseNumbatListener) ExitCall_args(ctx *Call_argsContext) {}

// EnterCall_secondary is called when production call_secondary is entered.
func (s *BaseNumbatListener) EnterCall_secondary(ctx *Call_secondaryContext) {}

// ExitCall_secondary is called when production call_secondary is exited.
func (s *BaseNumbatListener) ExitCall_secondary(ctx *Call_secondaryContext) {}

// EnterCall_primary is called when production call_primary is entered.
func (s *BaseNumbatListener) EnterCall_primary(ctx *Call_primaryContext) {}

// ExitCall_primary is called when production call_primary is exited.
func (s *BaseNumbatListener) ExitCall_primary(ctx *Call_primaryContext) {}

// EnterCall is called when production call is entered.
func (s *BaseNumbatListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseNumbatListener) ExitCall(ctx *CallContext) {}

// EnterCall_stmt is called when production call_stmt is entered.
func (s *BaseNumbatListener) EnterCall_stmt(ctx *Call_stmtContext) {}

// ExitCall_stmt is called when production call_stmt is exited.
func (s *BaseNumbatListener) ExitCall_stmt(ctx *Call_stmtContext) {}

// EnterVar_expr is called when production var_expr is entered.
func (s *BaseNumbatListener) EnterVar_expr(ctx *Var_exprContext) {}

// ExitVar_expr is called when production var_expr is exited.
func (s *BaseNumbatListener) ExitVar_expr(ctx *Var_exprContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseNumbatListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseNumbatListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterVar_name is called when production var_name is entered.
func (s *BaseNumbatListener) EnterVar_name(ctx *Var_nameContext) {}

// ExitVar_name is called when production var_name is exited.
func (s *BaseNumbatListener) ExitVar_name(ctx *Var_nameContext) {}

// EnterVar_stmt is called when production var_stmt is entered.
func (s *BaseNumbatListener) EnterVar_stmt(ctx *Var_stmtContext) {}

// ExitVar_stmt is called when production var_stmt is exited.
func (s *BaseNumbatListener) ExitVar_stmt(ctx *Var_stmtContext) {}

// EnterAssignment_expr is called when production assignment_expr is entered.
func (s *BaseNumbatListener) EnterAssignment_expr(ctx *Assignment_exprContext) {}

// ExitAssignment_expr is called when production assignment_expr is exited.
func (s *BaseNumbatListener) ExitAssignment_expr(ctx *Assignment_exprContext) {}

// EnterAssignment_var is called when production assignment_var is entered.
func (s *BaseNumbatListener) EnterAssignment_var(ctx *Assignment_varContext) {}

// ExitAssignment_var is called when production assignment_var is exited.
func (s *BaseNumbatListener) ExitAssignment_var(ctx *Assignment_varContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseNumbatListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseNumbatListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterReturn_expr is called when production return_expr is entered.
func (s *BaseNumbatListener) EnterReturn_expr(ctx *Return_exprContext) {}

// ExitReturn_expr is called when production return_expr is exited.
func (s *BaseNumbatListener) ExitReturn_expr(ctx *Return_exprContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseNumbatListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseNumbatListener) ExitReturn(ctx *ReturnContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseNumbatListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseNumbatListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseNumbatListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseNumbatListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseNumbatListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseNumbatListener) ExitStatement(ctx *StatementContext) {}

// EnterObject is called when production object is entered.
func (s *BaseNumbatListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseNumbatListener) ExitObject(ctx *ObjectContext) {}
