// Code generated from Numbat.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Numbat

import "github.com/antlr4-go/antlr/v4"

// NumbatListener is a complete listener for a parse tree produced by NumbatParser.
type NumbatListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr_bool is called when entering the expr_bool production.
	EnterExpr_bool(c *Expr_boolContext)

	// EnterExpr_num is called when entering the expr_num production.
	EnterExpr_num(c *Expr_numContext)

	// EnterExpr_hex is called when entering the expr_hex production.
	EnterExpr_hex(c *Expr_hexContext)

	// EnterExpr_str is called when entering the expr_str production.
	EnterExpr_str(c *Expr_strContext)

	// EnterExpr_null is called when entering the expr_null production.
	EnterExpr_null(c *Expr_nullContext)

	// EnterExpr_constant is called when entering the expr_constant production.
	EnterExpr_constant(c *Expr_constantContext)

	// EnterExpr_var is called when entering the expr_var production.
	EnterExpr_var(c *Expr_varContext)

	// EnterExpr_call is called when entering the expr_call production.
	EnterExpr_call(c *Expr_callContext)

	// EnterExpr_all is called when entering the expr_all production.
	EnterExpr_all(c *Expr_allContext)

	// EnterParam_expr is called when entering the param_expr production.
	EnterParam_expr(c *Param_exprContext)

	// EnterParam_type is called when entering the param_type production.
	EnterParam_type(c *Param_typeContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterType_out is called when entering the type_out production.
	EnterType_out(c *Type_outContext)

	// EnterType_in is called when entering the type_in production.
	EnterType_in(c *Type_inContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterProc_body is called when entering the proc_body production.
	EnterProc_body(c *Proc_bodyContext)

	// EnterProc_name is called when entering the proc_name production.
	EnterProc_name(c *Proc_nameContext)

	// EnterProc_type is called when entering the proc_type production.
	EnterProc_type(c *Proc_typeContext)

	// EnterProc_def is called when entering the proc_def production.
	EnterProc_def(c *Proc_defContext)

	// EnterProc is called when entering the proc production.
	EnterProc(c *ProcContext)

	// EnterCall_expr is called when entering the call_expr production.
	EnterCall_expr(c *Call_exprContext)

	// EnterCall_secondary is called when entering the call_secondary production.
	EnterCall_secondary(c *Call_secondaryContext)

	// EnterCall_primary is called when entering the call_primary production.
	EnterCall_primary(c *Call_primaryContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterCall_stmt is called when entering the call_stmt production.
	EnterCall_stmt(c *Call_stmtContext)

	// EnterVar_expr is called when entering the var_expr production.
	EnterVar_expr(c *Var_exprContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterVar_name is called when entering the var_name production.
	EnterVar_name(c *Var_nameContext)

	// EnterVar_stmt is called when entering the var_stmt production.
	EnterVar_stmt(c *Var_stmtContext)

	// EnterAssignment_expr is called when entering the assignment_expr production.
	EnterAssignment_expr(c *Assignment_exprContext)

	// EnterAssignment_var is called when entering the assignment_var production.
	EnterAssignment_var(c *Assignment_varContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterReturn_expr is called when entering the return_expr production.
	EnterReturn_expr(c *Return_exprContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr_bool is called when exiting the expr_bool production.
	ExitExpr_bool(c *Expr_boolContext)

	// ExitExpr_num is called when exiting the expr_num production.
	ExitExpr_num(c *Expr_numContext)

	// ExitExpr_hex is called when exiting the expr_hex production.
	ExitExpr_hex(c *Expr_hexContext)

	// ExitExpr_str is called when exiting the expr_str production.
	ExitExpr_str(c *Expr_strContext)

	// ExitExpr_null is called when exiting the expr_null production.
	ExitExpr_null(c *Expr_nullContext)

	// ExitExpr_constant is called when exiting the expr_constant production.
	ExitExpr_constant(c *Expr_constantContext)

	// ExitExpr_var is called when exiting the expr_var production.
	ExitExpr_var(c *Expr_varContext)

	// ExitExpr_call is called when exiting the expr_call production.
	ExitExpr_call(c *Expr_callContext)

	// ExitExpr_all is called when exiting the expr_all production.
	ExitExpr_all(c *Expr_allContext)

	// ExitParam_expr is called when exiting the param_expr production.
	ExitParam_expr(c *Param_exprContext)

	// ExitParam_type is called when exiting the param_type production.
	ExitParam_type(c *Param_typeContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitType_out is called when exiting the type_out production.
	ExitType_out(c *Type_outContext)

	// ExitType_in is called when exiting the type_in production.
	ExitType_in(c *Type_inContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitProc_body is called when exiting the proc_body production.
	ExitProc_body(c *Proc_bodyContext)

	// ExitProc_name is called when exiting the proc_name production.
	ExitProc_name(c *Proc_nameContext)

	// ExitProc_type is called when exiting the proc_type production.
	ExitProc_type(c *Proc_typeContext)

	// ExitProc_def is called when exiting the proc_def production.
	ExitProc_def(c *Proc_defContext)

	// ExitProc is called when exiting the proc production.
	ExitProc(c *ProcContext)

	// ExitCall_expr is called when exiting the call_expr production.
	ExitCall_expr(c *Call_exprContext)

	// ExitCall_secondary is called when exiting the call_secondary production.
	ExitCall_secondary(c *Call_secondaryContext)

	// ExitCall_primary is called when exiting the call_primary production.
	ExitCall_primary(c *Call_primaryContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitCall_stmt is called when exiting the call_stmt production.
	ExitCall_stmt(c *Call_stmtContext)

	// ExitVar_expr is called when exiting the var_expr production.
	ExitVar_expr(c *Var_exprContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitVar_name is called when exiting the var_name production.
	ExitVar_name(c *Var_nameContext)

	// ExitVar_stmt is called when exiting the var_stmt production.
	ExitVar_stmt(c *Var_stmtContext)

	// ExitAssignment_expr is called when exiting the assignment_expr production.
	ExitAssignment_expr(c *Assignment_exprContext)

	// ExitAssignment_var is called when exiting the assignment_var production.
	ExitAssignment_var(c *Assignment_varContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitReturn_expr is called when exiting the return_expr production.
	ExitReturn_expr(c *Return_exprContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)
}
