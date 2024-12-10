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
	varStmt   *Var

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
	listener.proc.Name = ctx.NON_TYPE_NAME().GetText()
}

func (listener *Listener) EnterProc_type(ctx *parser.Proc_typeContext) {
	listener.proc.ReturnType = &Type{}
	listener.SetType(listener.proc.ReturnType)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (listener *Listener) ExitType(ctx *parser.TypeContext) {
	listener.UnsetType()
}

func (listener *Listener) EnterType_out(ctx *parser.Type_outContext) {
	listener.typ.Out.Name = ctx.TYPE_NAME().GetText()
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (listener *Listener) EnterParam(ctx *parser.ParamContext) {
	listener.typ.In = append(listener.typ.In, Param{
		Name: ctx.NON_TYPE_NAME().GetText(),
	})
	listener.typ.Param = &listener.typ.In[len(listener.typ.In)-1]
}

func (listener *Listener) ExitParam(ctx *parser.ParamContext) {
	listener.typ.Param = nil
}

func (listener *Listener) EnterParam_expr(ctx *parser.Param_exprContext) {
	listener.typ.Param.Expr = make([]Expr, 0)
	listener.exprs = &listener.typ.Param.Expr
}

func (listener *Listener) ExitParam_expr(ctx *parser.Param_exprContext) {
	listener.exprs = nil
}

func (listener *Listener) EnterParam_type(ctx *parser.Param_typeContext) {
	listener.typ.Param.Typ = &Type{}
	listener.SetType(listener.typ.Param.Typ)
}

// ============================================================================================================
// VAR
// ============================================================================================================

func (listener *Listener) EnterVar_stmt(ctx *parser.Var_stmtContext) {
	listener.proc.Statements = append(listener.proc.Statements, Statement{
		Var: &Var{},
	})
	listener.statement = &listener.proc.Statements[len(listener.proc.Statements)-1]
	listener.varStmt = listener.statement.Var
}

func (listener *Listener) ExitVar_stmt(ctx *parser.Var_stmtContext) {
	listener.varStmt = nil
}

func (listener *Listener) EnterVar_expr(ctx *parser.Var_exprContext) {
	listener.varStmt.Exprs = make([]Expr, 0)
	listener.exprs = &listener.varStmt.Exprs
}

func (listener *Listener) ExitVar_expr(ctx *parser.Var_exprContext) {
	listener.exprs = nil
}

func (listener *Listener) EnterVar_name(ctx *parser.Var_nameContext) {
	listener.varStmt.Name = ctx.NON_TYPE_NAME().GetText()
}

func (listener *Listener) EnterVar_type(ctx *parser.Var_typeContext) {
	listener.varStmt.VarType = &Type{}
	listener.SetType(listener.varStmt.VarType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (listener *Listener) EnterCall_stmt(ctx *parser.Call_stmtContext) {
	listener.proc.Statements = append(listener.proc.Statements, Statement{
		Call: &Call{},
	})
	listener.statement = &listener.proc.Statements[len(listener.proc.Statements)-1]
	listener.call = listener.statement.Call
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
	listener.call.Primary = primary
}

func (listener *Listener) EnterCall_secondary(ctx *parser.Call_secondaryContext) {
	listener.statement.Call.Secondary = ctx.NON_TYPE_NAME().GetText()
}

func (listener *Listener) EnterCall_expr(ctx *parser.Call_exprContext) {
	listener.call.Exprs = make([]Expr, 0)
	listener.exprs = &listener.call.Exprs
}

func (listener *Listener) ExitCall_expr(ctx *parser.Call_exprContext) {
	listener.exprs = nil
}

// ============================================================================================================
// RETURN
// ============================================================================================================

func (listener *Listener) EnterReturn_stmt(ctx *parser.Return_stmtContext) {
	listener.proc.Statements = append(listener.proc.Statements, Statement{
		Ret: &Return{},
	})
	listener.statement = &listener.proc.Statements[len(listener.proc.Statements)-1]
}

func (listener *Listener) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	listener.statement = nil
}

func (listener *Listener) EnterReturn_expr(ctx *parser.Return_exprContext) {
	listener.statement.Ret.Exprs = make([]Expr, 0)
	listener.exprs = &listener.statement.Ret.Exprs
}

func (listener *Listener) ExitReturn_expr(ctx *parser.Return_exprContext) {
	listener.exprs = nil
}

// ============================================================================================================
// ASSIGNMENT
// ============================================================================================================

func (listener *Listener) EnterAssignment(ctx *parser.AssignmentContext) {
	listener.proc.Statements = append(listener.proc.Statements, Statement{
		Assignment: &Assignment{},
	})
	listener.statement = &listener.proc.Statements[len(listener.proc.Statements)-1]
}

func (listener *Listener) ExitAssignment(ctx *parser.AssignmentContext) {
	listener.statement = nil
}

func (listener *Listener) EnterAssignment_var(ctx *parser.Assignment_varContext) {
	listener.statement.Assignment.Vars = append(listener.statement.Assignment.Vars, ctx.NON_TYPE_NAME().GetText())
}

func (listener *Listener) EnterAssignment_expr(ctx *parser.Assignment_exprContext) {
	listener.statement.Assignment.Exprs = make([]Expr, 0)
	listener.exprs = &listener.statement.Assignment.Exprs
}

func (listener *Listener) ExitAssignment_expr(ctx *parser.Assignment_exprContext) {
	listener.exprs = nil
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (listener *Listener) EnterExpr_bool(ctx *parser.Expr_boolContext) {
	str := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{Boolean: &str, Unit: "bool"})
}

func (listener *Listener) EnterExpr_num(ctx *parser.Expr_numContext) {
	value := ctx.NUMBER().GetText()
	unit := "num"
	if ctx.UNIT() != nil {
		unit = ctx.UNIT().GetText()[1:]
	}
	*listener.exprs = append(*listener.exprs, Expr{Number: &value, Unit: unit})
}

func (listener *Listener) EnterExpr_hex(ctx *parser.Expr_hexContext) {
	value := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{Hex: &value, Unit: "hex"})
}

func (listener *Listener) EnterExpr_str(ctx *parser.Expr_strContext) {
	value := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{Str: &value, Unit: "str"})
}

func (listener *Listener) EnterExpr_null(ctx *parser.Expr_nullContext) {
	*listener.exprs = append(*listener.exprs, Expr{Null: true})
}

func (listener *Listener) EnterExpr_var(ctx *parser.Expr_varContext) {
	value := ctx.GetText()
	*listener.exprs = append(*listener.exprs, Expr{VarName: &VarName{value}})
}

func (listener *Listener) EnterExpr_call(ctx *parser.Expr_callContext) {
	call := &Call{}
	*listener.exprs = append(*listener.exprs, Expr{Call: call})
	listener.call = call
}
