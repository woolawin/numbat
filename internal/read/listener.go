package read

import (
	"github.com/antlr4-go/antlr/v4"
	"numbat/internal/common"
	"numbat/internal/parser"
)

type SourceReader struct {
	*parser.BaseNumbatListener
	fileName string

	programs []Proc

	procs []Proc

	program   *Proc
	proc      *Proc
	statement *Statement

	exprs   ExprGroupStack
	calls   CallStack
	types   TypeStack
	varStmt *Var
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
	reader.exprs = ExprGroupStack{}
	reader.calls = CallStack{exprStack: &reader.exprs}
	reader.types = TypeStack{}
	reader.varStmt = nil
}

// ============================================================================================================
// PROGRAM
// ============================================================================================================

func (reader *SourceReader) EnterProgram(ctx *parser.ProgramContext) {
	reader.programs = append(reader.programs, Proc{
		Location: common.Location{
			File:   reader.fileName,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
	})
	reader.program = &reader.programs[len(reader.programs)-1]
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
	reader.types.Push(reader.proc.Type)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (reader *SourceReader) EnterType(ctx *parser.TypeContext) {
	reader.types.Current().Location = reader.location(ctx.BaseParserRuleContext)
}

func (reader *SourceReader) ExitType(ctx *parser.TypeContext) {
	reader.types.Pop()
}

func (reader *SourceReader) EnterType_super_atomic(ctx *parser.Type_super_atomicContext) {
	reader.types.Current().Out.Name = ctx.TYPE_NAME().GetText()
	reader.types.Current().Out.Location = reader.location(ctx.BaseParserRuleContext)
}

func (reader *SourceReader) EnterType_super_atomic_seq(ctx *parser.Type_super_atomic_seqContext) {
	reader.types.Current().Out.Sequence = true
	if ctx.NUMBER() != nil {
		reader.types.Current().Out.SequenceSize = ctx.NUMBER().GetText()
		reader.types.Current().Out.SequenceLocation = reader.location(ctx.BaseParserRuleContext)
	}
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (reader *SourceReader) EnterParam(ctx *parser.ParamContext) {
	reader.types.Current().In = append(reader.types.Current().In, Param{
		Name: Name{
			Value:    ctx.NON_TYPE_NAME().GetText(),
			Location: reader.location(ctx.BaseParserRuleContext),
		},
	})
	reader.types.Current().Param = &reader.types.Current().In[len(reader.types.Current().In)-1]
}

func (reader *SourceReader) ExitParam(ctx *parser.ParamContext) {
	reader.types.Current().Param = nil
}

func (reader *SourceReader) EnterParam_expr(ctx *parser.Param_exprContext) {
	reader.types.Current().Param.Value = ExprGroup{}
	reader.exprs.Push(&reader.types.Current().Param.Value)
}

func (reader *SourceReader) ExitParam_expr(ctx *parser.Param_exprContext) {
	reader.exprs.Pop()
}

func (reader *SourceReader) EnterParam_type(ctx *parser.Param_typeContext) {
	reader.types.Current().Param.Typ = &Type{}
	reader.types.Push(reader.types.Current().Param.Typ)
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
	reader.exprs.Push(&reader.varStmt.Value)
}

func (reader *SourceReader) ExitVar_expr(ctx *parser.Var_exprContext) {
	reader.exprs.Pop()
}

func (reader *SourceReader) EnterVar_name(ctx *parser.Var_nameContext) {
	reader.varStmt.Name = Name{
		Value:    ctx.NON_TYPE_NAME().GetText(),
		Location: reader.location(ctx.BaseParserRuleContext),
	}
}

func (reader *SourceReader) EnterVar_type(ctx *parser.Var_typeContext) {
	reader.varStmt.VarType = &Type{}
	reader.types.Push(reader.varStmt.VarType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (reader *SourceReader) EnterCall_stmt(ctx *parser.Call_stmtContext) {
	reader.proc.Statements = append(reader.proc.Statements, Statement{
		Call: &Call{
			Arguments: ExprGroup{
				Exprs: make([]Expr, 0),
			},
		},
		Location: reader.location(ctx.BaseParserRuleContext),
	})
	reader.statement = &reader.proc.Statements[len(reader.proc.Statements)-1]
	reader.calls.Push(reader.statement.Call)
}

func (reader *SourceReader) ExitCall_stmt(ctx *parser.Call_stmtContext) {
	reader.calls.Pop()
}

func (reader *SourceReader) EnterCall_primary(ctx *parser.Call_primaryContext) {
	primary := ""
	if ctx.NON_TYPE_NAME() != nil {
		primary = ctx.NON_TYPE_NAME().GetText()
	} else if ctx.TYPE_NAME() != nil {
		primary = ctx.TYPE_NAME().GetText()
	}
	reader.calls.Current().Primary = Name{
		Value:    primary,
		Location: reader.location(ctx.BaseParserRuleContext),
	}
}

func (reader *SourceReader) EnterCall_secondary(ctx *parser.Call_secondaryContext) {
	reader.calls.Current().Secondary = Name{
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
	reader.exprs.Push(&reader.statement.Ret.Value)
}

func (reader *SourceReader) ExitReturn_expr(ctx *parser.Return_exprContext) {
	reader.exprs.Pop()
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
	reader.exprs.Push(&reader.statement.Assignment.Values)
}

func (reader *SourceReader) ExitAssignment_expr(ctx *parser.Assignment_exprContext) {
	reader.exprs.Pop()
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (reader *SourceReader) EnterExpr_bool(ctx *parser.Expr_boolContext) {
	str := ctx.GetText()

	reader.exprs.Current().Add(Expr{Boolean: &str, Unit: "bool", Location: reader.location(ctx.BaseParserRuleContext)})
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
	reader.exprs.Current().Add(Expr{Number: &value, Unit: unit, Location: reader.location(ctx.BaseParserRuleContext)})
}

func (reader *SourceReader) EnterExpr_hex(ctx *parser.Expr_hexContext) {
	value := ctx.GetText()
	reader.exprs.Current().Add(Expr{Hex: &value, Unit: "hex", Location: reader.location(ctx.BaseParserRuleContext)})
}

func (reader *SourceReader) EnterExpr_str(ctx *parser.Expr_strContext) {
	value := ctx.GetText()
	reader.exprs.Current().Add(Expr{Str: &value, Unit: "str", Location: reader.location(ctx.BaseParserRuleContext)})
}

func (reader *SourceReader) EnterExpr_null(ctx *parser.Expr_nullContext) {
	reader.exprs.Current().Add(Expr{Null: true, Location: reader.location(ctx.BaseParserRuleContext)})
}

func (reader *SourceReader) EnterExpr_var(ctx *parser.Expr_varContext) {
	reader.exprs.Current().Add(
		Expr{
			VarName: &Name{
				Value:    ctx.GetText(),
				Location: reader.location(ctx.BaseParserRuleContext),
			},
			Location: reader.location(ctx.BaseParserRuleContext),
		},
	)
}

func (reader *SourceReader) EnterExpr_call(ctx *parser.Expr_callContext) {
	call := &Call{
		Arguments: ExprGroup{
			Exprs: make([]Expr, 0),
		},
	}
	reader.exprs.Current().Add(Expr{Call: call, Location: reader.location(ctx.BaseParserRuleContext)})
	reader.calls.Push(call)
}

func (reader *SourceReader) ExitExpr_call(ctx *parser.Expr_callContext) {
	reader.calls.Pop()
}

func (reader *SourceReader) EnterExpr_seq(ctx *parser.Expr_seqContext) {
	seq := Expr{
		Seq:      &ExprGroup{},
		Location: reader.location(ctx.BaseParserRuleContext),
	}

	reader.exprs.Current().Add(seq)
	reader.exprs.Push(seq.Seq)
}

func (reader *SourceReader) ExitExpr_seq(ctx *parser.Expr_seqContext) {
	reader.exprs.Pop()
}
