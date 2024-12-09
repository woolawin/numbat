package read

import (
	"numbat/parser"
)

type Listener struct {
	*parser.BaseNumbatListener

	program *Proc

	procs []Proc

	proc      *Proc
	statement *Statement
	call      *Call
	exprs     *[]Expr
	let       *Let

	typ       *Type
	typeChain []*Type
}

func NewListener() *Listener {
	return &Listener{
		BaseNumbatListener: new(parser.BaseNumbatListener),
	}
}

func (listener *Listener) SetType(t *Type) {
	listener.typeChain = append(listener.typeChain, t)
	listener.typ = t
}

func (listener *Listener) UnsetType() {
	if len(listener.typeChain) == 0 {
		return
	}

	if len(listener.typeChain) == 1 {
		listener.typ = listener.typeChain[0]
		listener.typeChain = nil
		return
	}

	listener.typeChain = listener.typeChain[:len(listener.typeChain)-1]
	listener.typ = listener.typeChain[len(listener.typeChain)-1]
}

// ============================================================================================================
// PROGRAM
// ============================================================================================================

func (listener *Listener) EnterProgram(ctx *parser.ProgramContext) {
	listener.program = &Proc{}
	listener.proc = listener.program
}

func (listener *Listener) ExitProgram(ctx *parser.ProgramContext) {
	listener.proc = nil
}

// ============================================================================================================
// PROC
// ============================================================================================================

func (listener *Listener) EnterProc(ctx *parser.ProcContext) {
	listener.procs = append(listener.procs, Proc{})
	listener.proc = &listener.procs[len(listener.procs)-1]
}

func (listener *Listener) ExitProc(ctx *parser.ProcContext) {
	listener.proc = nil
}

func (listener *Listener) EnterProc_name(ctx *parser.Proc_nameContext) {
	listener.proc.name = ctx.NON_TYPE_NAME().GetText()
}

func (listener *Listener) EnterProc_type(ctx *parser.Proc_typeContext) {
	listener.proc.returnType = &Type{}
	listener.SetType(listener.proc.returnType)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (listener *Listener) ExitType(ctx *parser.TypeContext) {
	listener.UnsetType()
}

func (listener *Listener) EnterType_out(ctx *parser.Type_outContext) {
	listener.typ.out.name = ctx.TYPE_NAME().GetText()
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (listener *Listener) EnterParam(ctx *parser.ParamContext) {
	listener.typ.in = append(listener.typ.in, Param{
		name: ctx.NON_TYPE_NAME().GetText(),
	})
	listener.typ.param = &listener.typ.in[len(listener.typ.in)-1]
}

func (listener *Listener) ExitParam(ctx *parser.ParamContext) {
	listener.typ.param = nil
}

func (listener *Listener) EnterParam_expr(ctx *parser.Param_exprContext) {
	listener.typ.param.expr = make([]Expr, 0)
	listener.exprs = &listener.typ.param.expr
}

func (listener *Listener) ExitParam_expr(ctx *parser.Param_exprContext) {
	listener.exprs = nil
}

func (listener *Listener) EnterParam_type(ctx *parser.Param_typeContext) {
	listener.typ.param.typ = &Type{}
	listener.SetType(listener.typ.param.typ)
}

// ============================================================================================================
// LET
// ============================================================================================================

func (listener *Listener) EnterLet(ctx *parser.LetContext) {
	listener.proc.statements = append(listener.proc.statements, Statement{
		let: &Let{},
	})
	listener.statement = &listener.proc.statements[len(listener.proc.statements)-1]
	listener.let = listener.statement.let
}

func (listener *Listener) ExitLet(ctx *parser.LetContext) {
	listener.let = nil
}

func (listener *Listener) EnterLet_expr(ctx *parser.Let_exprContext) {
	listener.let.exprs = make([]Expr, 0)
	listener.exprs = &listener.let.exprs
}

func (listener *Listener) ExitLet_expr(ctx *parser.Let_exprContext) {
	listener.exprs = nil
}

func (listener *Listener) EnterLet_var_name(ctx *parser.Let_var_nameContext) {
	listener.let.varName = ctx.NON_TYPE_NAME().GetText()
}

func (listener *Listener) EnterLet_var_type(ctx *parser.Let_var_typeContext) {
	listener.let.varType = &Type{}
	listener.SetType(listener.let.varType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (listener *Listener) EnterCall_stmt(ctx *parser.Call_stmtContext) {
	listener.proc.statements = append(listener.proc.statements, Statement{
		call: &Call{},
	})
	listener.statement = &listener.proc.statements[len(listener.proc.statements)-1]
	listener.call = listener.statement.call
}

func (listener *Listener) ExitCall_stmt(ctx *parser.Call_stmtContext) {
	listener.call = nil
}

func (listener *Listener) EnterCall_primary(ctx *parser.Call_primaryContext) {
	primary := ""
	if ctx.NON_TYPE_NAME() != nil {
		primary = ctx.NON_TYPE_NAME().GetText()
	} else if ctx.TYPE_NAME() != nil {
		primary = ctx.TYPE_NAME().GetText()
	}
	listener.call.primary = primary
}

func (listener *Listener) EnterCall_secondary(ctx *parser.Call_secondaryContext) {
	listener.statement.call.secondary = ctx.NON_TYPE_NAME().GetText()
}

func (listener *Listener) EnterCall_expr(ctx *parser.Call_exprContext) {
	listener.call.exprs = make([]Expr, 0)
	listener.exprs = &listener.call.exprs
}

func (listener *Listener) ExitCall_expr(ctx *parser.Call_exprContext) {
	listener.exprs = nil
}

// ============================================================================================================
// RETURN
// ============================================================================================================

func (listener *Listener) EnterReturn_stmt(ctx *parser.Return_stmtContext) {
	listener.proc.statements = append(listener.proc.statements, Statement{
		ret: &Return{},
	})
	listener.statement = &listener.proc.statements[len(listener.proc.statements)-1]
}

func (listener *Listener) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	listener.statement = nil
}

func (listener *Listener) EnterReturn_expr(ctx *parser.Return_exprContext) {
	listener.statement.ret.exprs = make([]Expr, 0)
	listener.exprs = &listener.statement.ret.exprs
}

func (listener *Listener) ExitReturn_expr(ctx *parser.Return_exprContext) {
	listener.exprs = nil
}

// ============================================================================================================
// ASSIGNMENT
// ============================================================================================================

func (listener *Listener) EnterAssignment(ctx *parser.AssignmentContext) {
	listener.proc.statements = append(listener.proc.statements, Statement{
		assignment: &Assignment{},
	})
	listener.statement = &listener.proc.statements[len(listener.proc.statements)-1]
}

func (listener *Listener) ExitAssignment(ctx *parser.AssignmentContext) {
	listener.statement = nil
}

func (listener *Listener) EnterAssignment_var(ctx *parser.Assignment_varContext) {
	listener.statement.assignment.vars = append(listener.statement.assignment.vars, ctx.NON_TYPE_NAME().GetText())
}

func (listener *Listener) EnterAssignment_expr(ctx *parser.Assignment_exprContext) {
	listener.statement.assignment.exprs = make([]Expr, 0)
	listener.exprs = &listener.statement.assignment.exprs
}

func (listener *Listener) ExitAssignment_expr(ctx *parser.Assignment_exprContext) {
	listener.exprs = nil
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (listener *Listener) EnterExpr_bool(ctx *parser.Expr_boolContext) {
	str := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{boolean: &str, unit: "bool"})
}

func (listener *Listener) EnterExpr_num(ctx *parser.Expr_numContext) {
	value := ctx.NUMBER().GetText()
	unit := "num"
	if ctx.Unit() != nil {
		unit = ctx.Unit().NON_TYPE_NAME().GetText()
	}
	*listener.exprs = append(*listener.exprs, Expr{number: &value, unit: unit})
}

func (listener *Listener) EnterExpr_hex(ctx *parser.Expr_hexContext) {
	value := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{hex: &value, unit: "hex"})
}

func (listener *Listener) EnterExpr_str(ctx *parser.Expr_strContext) {
	value := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{str: &value, unit: "str"})
}

func (listener *Listener) EnterExpr_null(ctx *parser.Expr_nullContext) {
	*listener.exprs = append(*listener.exprs, Expr{null: true})
}

func (listener *Listener) EnterExpr_var(ctx *parser.Expr_varContext) {
	value := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{varName: &VarName{value}})
}

func (listener *Listener) EnterExpr_call(ctx *parser.Expr_callContext) {
	call := &Call{}
	*listener.exprs = append(*listener.exprs, Expr{call: call})
	listener.call = call
}
