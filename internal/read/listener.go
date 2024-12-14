package read

import (
	"github.com/antlr4-go/antlr/v4"
	"numbat/internal/common"
	"numbat/internal/parser"
)

type SourceReader struct {
	*parser.BaseNumbatListener
	fileName string

	program *Proc

	procs []Proc

	proc      *Proc
	statement *Statement
	call      *Call
	callChain []*Call
	exprs     *[]Expr
	varStmt   *Var

	typ       *Type
	typeChain []*Type
}

func NewSourceReader() *SourceReader {
	return &SourceReader{
		BaseNumbatListener: new(parser.BaseNumbatListener),
	}
}

func (reader *SourceReader) Read(code string, fileName string) {
	reader.Reset()
	reader.fileName = fileName
	input := antlr.NewInputStream(code)
	lexer := parser.NewNumbatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewNumbatParser(stream)

	errorListener := &ErrorListener{}
	p.AddErrorListener(errorListener)

	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(reader, tree)

	reader.Reset()
}

func (reader *SourceReader) Reset() {
	reader.fileName = ""
	reader.proc = nil
	reader.statement = nil
	reader.call = nil
	reader.exprs = nil
	reader.varStmt = nil

	reader.typ = nil
	reader.typeChain = nil
}

func (reader *SourceReader) PushCall(call *Call) {
	reader.callChain = append(reader.callChain, call)
	reader.call = call
	reader.exprs = &reader.call.Exprs
}

func (reader *SourceReader) PopCall() {
	if len(reader.callChain) == 0 {
		return
	}

	if len(reader.callChain) == 1 {
		reader.call = nil
		reader.exprs = nil
		reader.callChain = nil
		return
	}

	reader.callChain = reader.callChain[:len(reader.callChain)-1]
	reader.call = reader.callChain[len(reader.callChain)-1]
	reader.exprs = &reader.call.Exprs
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

func (reader *SourceReader) EnterProgram(ctx *parser.ProgramContext) {
	reader.program = &Proc{
		Location: common.Location{
			File:   reader.fileName,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
	}
	reader.proc = reader.program
}

func (reader *SourceReader) ExitProgram(ctx *parser.ProgramContext) {
	reader.proc = nil
}

// ============================================================================================================
// PROC
// ============================================================================================================

func (reader *SourceReader) EnterProc(ctx *parser.ProcContext) {
	reader.procs = append(reader.procs, Proc{
		Location: common.Location{
			File:   reader.fileName,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
	})
	reader.proc = &reader.procs[len(reader.procs)-1]
}

func (reader *SourceReader) ExitProc(ctx *parser.ProcContext) {
	reader.proc = nil
}

func (reader *SourceReader) EnterProc_name(ctx *parser.Proc_nameContext) {
	reader.proc.Name = Name{
		Value:    ctx.NON_TYPE_NAME().GetText(),
		Location: reader.location(ctx.BaseParserRuleContext),
	}
}

func (reader *SourceReader) EnterProc_type(ctx *parser.Proc_typeContext) {
	reader.proc.Type = &Type{}
	reader.SetType(reader.proc.Type)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (reader *SourceReader) ExitType(ctx *parser.TypeContext) {
	reader.UnsetType()
}

func (reader *SourceReader) EnterType_out(ctx *parser.Type_outContext) {
	reader.typ.Out.Name = ctx.TYPE_NAME().GetText()
	reader.typ.Out.Location = reader.location(ctx.BaseParserRuleContext)
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (reader *SourceReader) EnterParam(ctx *parser.ParamContext) {
	reader.typ.In = append(reader.typ.In, Param{
		Name: Name{
			Value:    ctx.NON_TYPE_NAME().GetText(),
			Location: reader.location(ctx.BaseParserRuleContext),
		},
	})
	reader.typ.Param = &reader.typ.In[len(reader.typ.In)-1]
}

func (reader *SourceReader) ExitParam(ctx *parser.ParamContext) {
	reader.typ.Param = nil
}

func (reader *SourceReader) EnterParam_expr(ctx *parser.Param_exprContext) {
	reader.typ.Param.Expr = make([]Expr, 0)
	reader.exprs = &reader.typ.Param.Expr
}

func (reader *SourceReader) ExitParam_expr(ctx *parser.Param_exprContext) {
	reader.exprs = nil
}

func (reader *SourceReader) EnterParam_type(ctx *parser.Param_typeContext) {
	reader.typ.Param.Typ = &Type{}
	reader.SetType(reader.typ.Param.Typ)
}

// ============================================================================================================
// VAR
// ============================================================================================================

func (reader *SourceReader) EnterVar_stmt(ctx *parser.Var_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Var:      &Var{},
		Location: reader.location(ctx.BaseParserRuleContext),
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
	reader.varStmt = reader.statement.Var
}

func (reader *SourceReader) ExitVar_stmt(ctx *parser.Var_stmtContext) {
	reader.varStmt = nil
}

func (reader *SourceReader) EnterVar_expr(ctx *parser.Var_exprContext) {
	reader.varStmt.Exprs = make([]Expr, 0)
	reader.exprs = &reader.varStmt.Exprs
}

func (reader *SourceReader) ExitVar_expr(ctx *parser.Var_exprContext) {
	reader.exprs = nil
}

func (reader *SourceReader) EnterVar_name(ctx *parser.Var_nameContext) {
	reader.varStmt.Name = Name{
		Value:    ctx.NON_TYPE_NAME().GetText(),
		Location: reader.location(ctx.BaseParserRuleContext),
	}
}

func (reader *SourceReader) EnterVar_type(ctx *parser.Var_typeContext) {
	reader.varStmt.VarType = &Type{}
	reader.SetType(reader.varStmt.VarType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (reader *SourceReader) EnterCall_stmt(ctx *parser.Call_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Call: &Call{
			Exprs: make([]Expr, 0),
		},
		Location: reader.location(ctx.BaseParserRuleContext),
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
	reader.PushCall(reader.statement.Call)
}

func (reader *SourceReader) ExitCall_stmt(ctx *parser.Call_stmtContext) {
	reader.PopCall()
}

func (reader *SourceReader) EnterCall_primary(ctx *parser.Call_primaryContext) {
	primary := ""
	if ctx.NON_TYPE_NAME() != nil {
		primary = ctx.NON_TYPE_NAME().GetText()
	} else if ctx.TYPE_NAME() != nil {
		primary = ctx.TYPE_NAME().GetText()
	}
	reader.call.Primary = Name{
		Value:    primary,
		Location: reader.location(ctx.BaseParserRuleContext),
	}
}

func (reader *SourceReader) EnterCall_secondary(ctx *parser.Call_secondaryContext) {
	reader.call.Secondary = Name{
		Value:    ctx.NON_TYPE_NAME().GetText(),
		Location: reader.location(ctx.BaseParserRuleContext),
	}
}

// ============================================================================================================
// RETURN
// ============================================================================================================

func (reader *SourceReader) EnterReturn_stmt(ctx *parser.Return_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Ret:      &Return{},
		Location: reader.location(ctx.BaseParserRuleContext),
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
}

func (reader *SourceReader) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	reader.statement = nil
}

func (reader *SourceReader) EnterReturn_expr(ctx *parser.Return_exprContext) {
	reader.statement.Ret.Exprs = make([]Expr, 0)
	reader.exprs = &reader.statement.Ret.Exprs
}

func (reader *SourceReader) ExitReturn_expr(ctx *parser.Return_exprContext) {
	reader.exprs = nil
}

// ============================================================================================================
// ASSIGNMENT
// ============================================================================================================

func (reader *SourceReader) EnterAssignment(ctx *parser.AssignmentContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Assignment: &Assignment{},
		Location:   reader.location(ctx.BaseParserRuleContext),
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
}

func (reader *SourceReader) ExitAssignment(ctx *parser.AssignmentContext) {
	reader.statement = nil
}

func (reader *SourceReader) EnterAssignment_var(ctx *parser.Assignment_varContext) {
	reader.statement.Assignment.Vars = append(reader.statement.Assignment.Vars, ctx.NON_TYPE_NAME().GetText())
}

func (reader *SourceReader) EnterAssignment_expr(ctx *parser.Assignment_exprContext) {
	reader.statement.Assignment.Exprs = make([]Expr, 0)
	reader.exprs = &reader.statement.Assignment.Exprs
}

func (reader *SourceReader) ExitAssignment_expr(ctx *parser.Assignment_exprContext) {
	reader.exprs = nil
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (reader *SourceReader) EnterExpr_bool(ctx *parser.Expr_boolContext) {
	str := ctx.GetText()

	*reader.exprs = append(
		*reader.exprs,
		Expr{Boolean: &str, Unit: "bool", Location: reader.location(ctx.BaseParserRuleContext)},
	)
}

func (reader *SourceReader) location(ctx antlr.BaseParserRuleContext) common.Location {
	return common.Location{
		File:   reader.fileName,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn() + 1,
	}
}

func (reader *SourceReader) EnterExpr_num(ctx *parser.Expr_numContext) {
	value := ctx.NUMBER().GetText()
	unit := ""
	if ctx.UNIT() != nil {
		unit = ctx.UNIT().GetText()[1:]
	}
	*reader.exprs = append(
		*reader.exprs,
		Expr{Number: &value, Unit: unit, Location: reader.location(ctx.BaseParserRuleContext)},
	)
}

func (reader *SourceReader) EnterExpr_hex(ctx *parser.Expr_hexContext) {
	value := ctx.GetText()
	*reader.exprs = append(
		*reader.exprs,
		Expr{Hex: &value, Unit: "hex", Location: reader.location(ctx.BaseParserRuleContext)},
	)
}

func (reader *SourceReader) EnterExpr_str(ctx *parser.Expr_strContext) {
	value := ctx.GetText()
	*reader.exprs = append(
		*reader.exprs, Expr{Str: &value, Unit: "str", Location: reader.location(ctx.BaseParserRuleContext)},
	)
}

func (reader *SourceReader) EnterExpr_null(ctx *parser.Expr_nullContext) {
	*reader.exprs = append(
		*reader.exprs,
		Expr{Null: true, Location: reader.location(ctx.BaseParserRuleContext)},
	)
}

func (reader *SourceReader) EnterExpr_var(ctx *parser.Expr_varContext) {
	value := ctx.GetText()
	*reader.exprs = append(
		*reader.exprs,
		Expr{
			VarName: &Name{
				Value:    value,
				Location: reader.location(ctx.BaseParserRuleContext),
			},
			Location: reader.location(ctx.BaseParserRuleContext),
		},
	)
}

func (reader *SourceReader) EnterExpr_call(ctx *parser.Expr_callContext) {
	call := &Call{
		Exprs: make([]Expr, 0),
	}
	*reader.exprs = append(
		*reader.exprs,
		Expr{Call: call, Location: reader.location(ctx.BaseParserRuleContext)},
	)
	reader.PushCall(call)
}

func (reader *SourceReader) ExitExpr_call(ctx *parser.Expr_callContext) {
	reader.PopCall()
}
