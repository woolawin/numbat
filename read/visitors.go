package read

import "numbat/parser"

// ============================================================================================================
// PROGRAM
// ============================================================================================================

func (source *Source) EnterProgram(ctx *parser.ProgramContext) {
	source.program = &Proc{}
	source.proc = source.program
}

func (source *Source) ExitProgram(ctx *parser.ProgramContext) {
	source.proc = nil
}

// ============================================================================================================
// PROC
// ============================================================================================================

func (source *Source) EnterProc(ctx *parser.ProcContext) {
	source.procs = append(source.procs, Proc{})
	source.proc = &source.procs[len(source.procs)-1]
}

func (source *Source) ExitProc(ctx *parser.ProcContext) {
	source.proc = nil
}

func (source *Source) EnterProc_name(ctx *parser.Proc_nameContext) {
	source.proc.name = ctx.NON_TYPE_NAME().GetText()
}

func (source *Source) EnterProc_type(ctx *parser.Proc_typeContext) {
	source.proc.returnType = &Type{}
	source.SetType(source.proc.returnType)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (source *Source) ExitType(ctx *parser.TypeContext) {
	source.UnsetType()
}

func (source *Source) EnterType_out(ctx *parser.Type_outContext) {
	source.typ.out.name = ctx.TYPE_NAME().GetText()
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (source *Source) EnterParam(ctx *parser.ParamContext) {
	source.typ.in = append(source.typ.in, Param{
		name: ctx.NON_TYPE_NAME().GetText(),
	})
	source.typ.param = &source.typ.in[len(source.typ.in)-1]
}

func (source *Source) ExitParam(ctx *parser.ParamContext) {
	source.typ.param = nil
}

func (source *Source) EnterParam_expr(ctx *parser.Param_exprContext) {
	source.typ.param.expr = make([]Expr, 0)
	source.exprs = &source.typ.param.expr
}

func (source *Source) ExitParam_expr(ctx *parser.Param_exprContext) {
	source.exprs = nil
}

func (source *Source) EnterParam_type(ctx *parser.Param_typeContext) {
	source.typ.param.typ = &Type{}
	source.SetType(source.typ.param.typ)
}

// ============================================================================================================
// LET
// ============================================================================================================

func (source *Source) EnterLet(ctx *parser.LetContext) {
	source.proc.statements = append(source.proc.statements, Statement{
		let: &Let{},
	})
	source.statement = &source.proc.statements[len(source.proc.statements)-1]
	source.let = source.statement.let
}

func (source *Source) ExitLet(ctx *parser.LetContext) {
	source.let = nil
}

func (source *Source) EnterLet_expr(ctx *parser.Let_exprContext) {
	source.let.exprs = make([]Expr, 0)
	source.exprs = &source.let.exprs
}

func (source *Source) ExitLet_expr(ctx *parser.Let_exprContext) {
	source.exprs = nil
}

func (source *Source) EnterLet_var_name(ctx *parser.Let_var_nameContext) {
	source.let.varName = ctx.NON_TYPE_NAME().GetText()
}

func (source *Source) EnterLet_var_type(ctx *parser.Let_var_typeContext) {
	source.let.varType = &Type{}
	source.SetType(source.let.varType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (source *Source) EnterCall_stmt(ctx *parser.Call_stmtContext) {
	source.proc.statements = append(source.proc.statements, Statement{
		call: &Call{},
	})
	source.statement = &source.proc.statements[len(source.proc.statements)-1]
	source.call = source.statement.call
}

func (source *Source) ExitCall_stmt(ctx *parser.Call_stmtContext) {
	source.call = nil
}

func (source *Source) EnterCall_primary(ctx *parser.Call_primaryContext) {
	primary := ""
	if ctx.NON_TYPE_NAME() != nil {
		primary = ctx.NON_TYPE_NAME().GetText()
	} else if ctx.TYPE_NAME() != nil {
		primary = ctx.TYPE_NAME().GetText()
	}
	source.call.primary = primary
}

func (source *Source) EnterCall_secondary(ctx *parser.Call_secondaryContext) {
	source.statement.call.secondary = ctx.NON_TYPE_NAME().GetText()
}

func (source *Source) EnterCall_expr(ctx *parser.Call_exprContext) {
	source.call.exprs = make([]Expr, 0)
	source.exprs = &source.call.exprs
}

func (source *Source) ExitCall_expr(ctx *parser.Call_exprContext) {
	source.exprs = nil
}

// ============================================================================================================
// RETURN
// ============================================================================================================

func (source *Source) EnterReturn_stmt(ctx *parser.Return_stmtContext) {
	source.proc.statements = append(source.proc.statements, Statement{
		ret: &Return{},
	})
	source.statement = &source.proc.statements[len(source.proc.statements)-1]
}

func (source *Source) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	source.statement = nil
}

func (source *Source) EnterReturn_expr(ctx *parser.Return_exprContext) {
	source.statement.ret.exprs = make([]Expr, 0)
	source.exprs = &source.statement.ret.exprs
}

func (source *Source) ExitReturn_expr(ctx *parser.Return_exprContext) {
	source.exprs = nil
}

// ============================================================================================================
// ASSIGNMENT
// ============================================================================================================

func (source *Source) EnterAssignment(ctx *parser.AssignmentContext) {
	source.proc.statements = append(source.proc.statements, Statement{
		assignment: &Assignment{},
	})
	source.statement = &source.proc.statements[len(source.proc.statements)-1]
}

func (source *Source) ExitAssignment(ctx *parser.AssignmentContext) {
	source.statement = nil
}

func (source *Source) EnterAssignment_var(ctx *parser.Assignment_varContext) {
	source.statement.assignment.vars = append(source.statement.assignment.vars, ctx.NON_TYPE_NAME().GetText())
}

func (source *Source) EnterAssignment_expr(ctx *parser.Assignment_exprContext) {
	source.statement.assignment.exprs = make([]Expr, 0)
	source.exprs = &source.statement.assignment.exprs
}

func (source *Source) ExitAssignment_expr(ctx *parser.Assignment_exprContext) {
	source.exprs = nil
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (source *Source) EnterExpr_bool(ctx *parser.Expr_boolContext) {
	str := ctx.GetText()
	*source.exprs = append(*source.exprs, Expr{boolean: &str})
}

func (source *Source) EnterExpr_num(ctx *parser.Expr_numContext) {
	value := ctx.GetText()
	*source.exprs = append(*source.exprs, Expr{number: &value})
}

func (source *Source) EnterExpr_hex(ctx *parser.Expr_hexContext) {
	value := ctx.GetText()
	*source.exprs = append(*source.exprs, Expr{hex: &value})
}

func (source *Source) EnterExpr_str(ctx *parser.Expr_strContext) {
	value := ctx.GetText()
	*source.exprs = append(*source.exprs, Expr{str: &value})
}

func (source *Source) EnterExpr_null(ctx *parser.Expr_nullContext) {
	*source.exprs = append(*source.exprs, Expr{null: true})
}

func (source *Source) EnterExpr_var(ctx *parser.Expr_varContext) {
	value := ctx.GetText()
	*source.exprs = append(*source.exprs, Expr{varName: &VarName{value}})
}

func (source *Source) EnterExpr_call(ctx *parser.Expr_callContext) {
	call := &Call{}
	*source.exprs = append(*source.exprs, Expr{call: call})
	source.call = call
}

