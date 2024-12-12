package read

import (
	"github.com/antlr4-go/antlr/v4"
	parser2 "numbat/internal/parser"
)

type SourceReader struct {
	*parser2.BaseNumbatListener

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

func NewSourceReader() *SourceReader {
	return &SourceReader{
		BaseNumbatListener: new(parser2.BaseNumbatListener),
	}
}

func (reader *SourceReader) Read(code string) {
	input := antlr.NewInputStream(code)
	lexer := parser2.NewNumbatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser2.NewNumbatParser(stream)

	errorListener := &ErrorListener{}
	p.AddErrorListener(errorListener)

	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(reader, tree)

	reader.Reset()
}

func (reader *SourceReader) Reset() {
	reader.proc = nil
	reader.statement = nil
	reader.call = nil
	reader.exprs = nil
	reader.varStmt = nil

	reader.typ = nil
	reader.typeChain = nil
}

func (reader *SourceReader) SetType(t *Type) {
	reader.typeChain = append(reader.typeChain, t)
	reader.typ = t
}

func (reader *SourceReader) UnsetType() {
	if len(reader.typeChain) == 0 {
		return
	}

	if len(reader.typeChain) == 1 {
		reader.typ = reader.typeChain[0]
		reader.typeChain = nil
		return
	}

	reader.typeChain = reader.typeChain[:len(reader.typeChain)-1]
	reader.typ = reader.typeChain[len(reader.typeChain)-1]
}

// ============================================================================================================
// PROGRAM
// ============================================================================================================

func (reader *SourceReader) EnterProgram(ctx *parser2.ProgramContext) {
	reader.program = &Proc{}
	reader.proc = reader.program
}

func (reader *SourceReader) ExitProgram(ctx *parser2.ProgramContext) {
	reader.proc = nil
}

// ============================================================================================================
// PROC
// ============================================================================================================

func (reader *SourceReader) EnterProc(ctx *parser2.ProcContext) {
	reader.procs = append(reader.procs, Proc{})
	reader.proc = &reader.procs[len(reader.procs)-1]
}

func (reader *SourceReader) ExitProc(ctx *parser2.ProcContext) {
	reader.proc = nil
}

func (reader *SourceReader) EnterProc_name(ctx *parser2.Proc_nameContext) {
	reader.proc.Name = ctx.NON_TYPE_NAME().GetText()
}

func (reader *SourceReader) EnterProc_type(ctx *parser2.Proc_typeContext) {
	reader.proc.ReturnType = &Type{}
	reader.SetType(reader.proc.ReturnType)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (reader *SourceReader) ExitType(ctx *parser2.TypeContext) {
	reader.UnsetType()
}

func (reader *SourceReader) EnterType_out(ctx *parser2.Type_outContext) {
	reader.typ.Out.Name = ctx.TYPE_NAME().GetText()
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (reader *SourceReader) EnterParam(ctx *parser2.ParamContext) {
	reader.typ.In = append(reader.typ.In, Param{
		Name: ctx.NON_TYPE_NAME().GetText(),
	})
	reader.typ.Param = &reader.typ.In[len(reader.typ.In)-1]
}

func (reader *SourceReader) ExitParam(ctx *parser2.ParamContext) {
	reader.typ.Param = nil
}

func (reader *SourceReader) EnterParam_expr(ctx *parser2.Param_exprContext) {
	reader.typ.Param.Expr = make([]Expr, 0)
	reader.exprs = &reader.typ.Param.Expr
}

func (reader *SourceReader) ExitParam_expr(ctx *parser2.Param_exprContext) {
	reader.exprs = nil
}

func (reader *SourceReader) EnterParam_type(ctx *parser2.Param_typeContext) {
	reader.typ.Param.Typ = &Type{}
	reader.SetType(reader.typ.Param.Typ)
}

// ============================================================================================================
// VAR
// ============================================================================================================

func (reader *SourceReader) EnterVar_stmt(ctx *parser2.Var_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Var: &Var{},
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
	reader.varStmt = reader.statement.Var
}

func (reader *SourceReader) ExitVar_stmt(ctx *parser2.Var_stmtContext) {
	reader.varStmt = nil
}

func (reader *SourceReader) EnterVar_expr(ctx *parser2.Var_exprContext) {
	reader.varStmt.Exprs = make([]Expr, 0)
	reader.exprs = &reader.varStmt.Exprs
}

func (reader *SourceReader) ExitVar_expr(ctx *parser2.Var_exprContext) {
	reader.exprs = nil
}

func (reader *SourceReader) EnterVar_name(ctx *parser2.Var_nameContext) {
	reader.varStmt.Name = ctx.NON_TYPE_NAME().GetText()
}

func (reader *SourceReader) EnterVar_type(ctx *parser2.Var_typeContext) {
	reader.varStmt.VarType = &Type{}
	reader.SetType(reader.varStmt.VarType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (reader *SourceReader) EnterCall_stmt(ctx *parser2.Call_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Call: &Call{},
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
	reader.call = reader.statement.Call
}

func (reader *SourceReader) ExitCall_stmt(ctx *parser2.Call_stmtContext) {
	reader.call = nil
}

func (reader *SourceReader) EnterCall_primary(ctx *parser2.Call_primaryContext) {
	primary := ""
	if ctx.NON_TYPE_NAME() != nil {
		primary = ctx.NON_TYPE_NAME().GetText()
	} else if ctx.TYPE_NAME() != nil {
		primary = ctx.TYPE_NAME().GetText()
	}
	reader.call.Primary = primary
}

func (reader *SourceReader) EnterCall_secondary(ctx *parser2.Call_secondaryContext) {
	reader.call.Secondary = ctx.NON_TYPE_NAME().GetText()
}

func (reader *SourceReader) EnterCall_expr(ctx *parser2.Call_exprContext) {
	reader.call.Exprs = make([]Expr, 0)
	reader.exprs = &reader.call.Exprs
}

func (reader *SourceReader) ExitCall_expr(ctx *parser2.Call_exprContext) {
	reader.exprs = nil
}

// ============================================================================================================
// RETURN
// ============================================================================================================

func (reader *SourceReader) EnterReturn_stmt(ctx *parser2.Return_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Ret: &Return{},
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
}

func (reader *SourceReader) ExitReturn_stmt(ctx *parser2.Return_stmtContext) {
	reader.statement = nil
}

func (reader *SourceReader) EnterReturn_expr(ctx *parser2.Return_exprContext) {
	reader.statement.Ret.Exprs = make([]Expr, 0)
	reader.exprs = &reader.statement.Ret.Exprs
}

func (reader *SourceReader) ExitReturn_expr(ctx *parser2.Return_exprContext) {
	reader.exprs = nil
}

// ============================================================================================================
// ASSIGNMENT
// ============================================================================================================

func (reader *SourceReader) EnterAssignment(ctx *parser2.AssignmentContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Assignment: &Assignment{},
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
}

func (reader *SourceReader) ExitAssignment(ctx *parser2.AssignmentContext) {
	reader.statement = nil
}

func (reader *SourceReader) EnterAssignment_var(ctx *parser2.Assignment_varContext) {
	reader.statement.Assignment.Vars = append(reader.statement.Assignment.Vars, ctx.NON_TYPE_NAME().GetText())
}

func (reader *SourceReader) EnterAssignment_expr(ctx *parser2.Assignment_exprContext) {
	reader.statement.Assignment.Exprs = make([]Expr, 0)
	reader.exprs = &reader.statement.Assignment.Exprs
}

func (reader *SourceReader) ExitAssignment_expr(ctx *parser2.Assignment_exprContext) {
	reader.exprs = nil
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (reader *SourceReader) EnterExpr_bool(ctx *parser2.Expr_boolContext) {
	str := ctx.GetText()
	*reader.exprs = append(*reader.exprs, Expr{Boolean: &str, Unit: "bool"})
}

func (reader *SourceReader) EnterExpr_num(ctx *parser2.Expr_numContext) {
	value := ctx.NUMBER().GetText()
	unit := "num"
	if ctx.UNIT() != nil {
		unit = ctx.UNIT().GetText()[1:]
	}
	*reader.exprs = append(*reader.exprs, Expr{Number: &value, Unit: unit})
}

func (reader *SourceReader) EnterExpr_hex(ctx *parser2.Expr_hexContext) {
	value := ctx.GetText()
	*reader.exprs = append(*reader.exprs, Expr{Hex: &value, Unit: "hex"})
}

func (reader *SourceReader) EnterExpr_str(ctx *parser2.Expr_strContext) {
	value := ctx.GetText()
	*reader.exprs = append(*reader.exprs, Expr{Str: &value, Unit: "str"})
}

func (reader *SourceReader) EnterExpr_null(ctx *parser2.Expr_nullContext) {
	*reader.exprs = append(*reader.exprs, Expr{Null: true})
}

func (reader *SourceReader) EnterExpr_var(ctx *parser2.Expr_varContext) {
	value := ctx.GetText()
	*reader.exprs = append(*reader.exprs, Expr{VarName: &VarName{value}})
}

func (reader *SourceReader) EnterExpr_call(ctx *parser2.Expr_callContext) {
	call := &Call{}
	*reader.exprs = append(*reader.exprs, Expr{Call: call})
	reader.call = call
}
