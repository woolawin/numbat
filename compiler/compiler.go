package compiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"numbat/parser"
	"strings"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
}

func (listener *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Println("ERROR")
}

type Compiler struct {
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

func NewCompiler() *Compiler {
	return &Compiler{
		BaseNumbatListener: new(parser.BaseNumbatListener),
	}
}

func (compiler *Compiler) SetType(t *Type) {
	compiler.typeChain = append(compiler.typeChain, t)
	compiler.typ = t
}

func (compiler *Compiler) UnsetType() {
	if len(compiler.typeChain) == 0 {
		return
	}

	if len(compiler.typeChain) == 1 {
		compiler.typ = compiler.typeChain[0]
		compiler.typeChain = nil
		return
	}

	compiler.typeChain = compiler.typeChain[:len(compiler.typeChain)-1]
	compiler.typ = compiler.typeChain[len(compiler.typeChain)-1]
}

// ============================================================================================================
// PROGRAM
// ============================================================================================================

func (compiler *Compiler) EnterProgram(ctx *parser.ProgramContext) {
	compiler.program = &Proc{}
	compiler.proc = compiler.program
}

func (compiler *Compiler) ExitProgram(ctx *parser.ProgramContext) {
	compiler.proc = nil
}

// ============================================================================================================
// PROC
// ============================================================================================================

func (compiler *Compiler) EnterProc(ctx *parser.ProcContext) {
	compiler.procs = append(compiler.procs, Proc{})
	compiler.proc = &compiler.procs[len(compiler.procs)-1]
}

func (compiler *Compiler) ExitProc(ctx *parser.ProcContext) {
	compiler.proc = nil
}

func (compiler *Compiler) EnterProc_name(ctx *parser.Proc_nameContext) {
	compiler.proc.name = ctx.NON_TYPE_NAME().GetText()
}

func (compiler *Compiler) EnterProc_type(ctx *parser.Proc_typeContext) {
	compiler.proc.returnType = &Type{}
	compiler.SetType(compiler.proc.returnType)
}

// ============================================================================================================
// TYPE
// ============================================================================================================

func (compiler *Compiler) ExitType(ctx *parser.TypeContext) {
	compiler.UnsetType()
}

func (compiler *Compiler) EnterType_out(ctx *parser.Type_outContext) {
	compiler.typ.out.name = ctx.TYPE_NAME().GetText()
}

// ============================================================================================================
// PARAM
// ============================================================================================================

func (compiler *Compiler) EnterParam(ctx *parser.ParamContext) {
	compiler.typ.in = append(compiler.typ.in, Param{
		name: ctx.NON_TYPE_NAME().GetText(),
	})
	compiler.typ.param = &compiler.typ.in[len(compiler.typ.in)-1]
}

func (compiler *Compiler) ExitParam(ctx *parser.ParamContext) {
	compiler.typ.param = nil
}

func (compiler *Compiler) EnterParam_expr(ctx *parser.Param_exprContext) {
	compiler.typ.param.expr = make([]Expr, 0)
	compiler.exprs = &compiler.typ.param.expr
}

func (compiler *Compiler) ExitParam_expr(ctx *parser.Param_exprContext) {
	compiler.exprs = nil
}

func (compiler *Compiler) EnterParam_type(ctx *parser.Param_typeContext) {
	compiler.typ.param.typ = &Type{}
	compiler.SetType(compiler.typ.param.typ)
}

// ============================================================================================================
// LET
// ============================================================================================================

func (compiler *Compiler) EnterLet(ctx *parser.LetContext) {
	compiler.proc.statements = append(compiler.proc.statements, Statement{
		let: &Let{},
	})
	compiler.statement = &compiler.proc.statements[len(compiler.proc.statements)-1]
	compiler.let = compiler.statement.let
}

func (compiler *Compiler) ExitLet(ctx *parser.LetContext) {
	compiler.let = nil
}

func (compiler *Compiler) EnterLet_expr(ctx *parser.Let_exprContext) {
	compiler.let.exprs = make([]Expr, 0)
	compiler.exprs = &compiler.let.exprs
}

func (compiler *Compiler) ExitLet_expr(ctx *parser.Let_exprContext) {
	compiler.exprs = nil
}

func (compiler *Compiler) EnterLet_var_name(ctx *parser.Let_var_nameContext) {
	compiler.let.varName = ctx.NON_TYPE_NAME().GetText()
}

func (compiler *Compiler) EnterLet_var_type(ctx *parser.Let_var_typeContext) {
	compiler.let.varType = &Type{}
	compiler.SetType(compiler.let.varType)
}

// ============================================================================================================
// CALL
// ============================================================================================================

func (compiler *Compiler) EnterCall_stmt(ctx *parser.Call_stmtContext) {
	compiler.proc.statements = append(compiler.proc.statements, Statement{
		call: &Call{},
	})
	compiler.statement = &compiler.proc.statements[len(compiler.proc.statements)-1]
	compiler.call = compiler.statement.call
}

func (compiler *Compiler) ExitCall_stmt(ctx *parser.Call_stmtContext) {
	compiler.call = nil
}

func (compiler *Compiler) EnterCall_primary(ctx *parser.Call_primaryContext) {
	primary := ""
	if ctx.NON_TYPE_NAME() != nil {
		primary = ctx.NON_TYPE_NAME().GetText()
	} else if ctx.TYPE_NAME() != nil {
		primary = ctx.TYPE_NAME().GetText()
	}
	compiler.call.primary = primary
}

func (compiler *Compiler) EnterCall_secondary(ctx *parser.Call_secondaryContext) {
	compiler.statement.call.secondary = ctx.NON_TYPE_NAME().GetText()
}

func (compiler *Compiler) EnterCall_expr(ctx *parser.Call_exprContext) {
	compiler.call.exprs = make([]Expr, 0)
	compiler.exprs = &compiler.call.exprs
}

func (compiler *Compiler) ExitCall_expr(ctx *parser.Call_exprContext) {
	compiler.exprs = nil
}

// ============================================================================================================
// RETURN
// ============================================================================================================

func (compiler *Compiler) EnterReturn_stmt(ctx *parser.Return_stmtContext) {
	compiler.proc.statements = append(compiler.proc.statements, Statement{
		ret: &Return{},
	})
	compiler.statement = &compiler.proc.statements[len(compiler.proc.statements)-1]
}

func (compiler *Compiler) ExitReturn_stmt(ctx *parser.Return_stmtContext) {
	compiler.statement = nil
}

func (compiler *Compiler) EnterReturn_expr(ctx *parser.Return_exprContext) {
	compiler.statement.ret.exprs = make([]Expr, 0)
	compiler.exprs = &compiler.statement.ret.exprs
}

func (compiler *Compiler) ExitReturn_expr(ctx *parser.Return_exprContext) {
	compiler.exprs = nil
}

// ============================================================================================================
// ASSIGNMENT
// ============================================================================================================

func (compiler *Compiler) EnterAssignment(ctx *parser.AssignmentContext) {
	compiler.proc.statements = append(compiler.proc.statements, Statement{
		assignment: &Assignment{},
	})
	compiler.statement = &compiler.proc.statements[len(compiler.proc.statements)-1]
}

func (compiler *Compiler) ExitAssignment(ctx *parser.AssignmentContext) {
	compiler.statement = nil
}

func (compiler *Compiler) EnterAssignment_var(ctx *parser.Assignment_varContext) {
	compiler.statement.assignment.vars = append(compiler.statement.assignment.vars, ctx.NON_TYPE_NAME().GetText())
}

func (compiler *Compiler) EnterAssignment_expr(ctx *parser.Assignment_exprContext) {
	compiler.statement.assignment.exprs = make([]Expr, 0)
	compiler.exprs = &compiler.statement.assignment.exprs
}

func (compiler *Compiler) ExitAssignment_expr(ctx *parser.Assignment_exprContext) {
	compiler.exprs = nil
}

// ============================================================================================================
// EXPR
// ============================================================================================================

func (compiler *Compiler) EnterExpr_bool(ctx *parser.Expr_boolContext) {
	str := ctx.GetText()
	*compiler.exprs = append(*compiler.exprs, Expr{boolean: &str})
}

func (compiler *Compiler) EnterExpr_num(ctx *parser.Expr_numContext) {
	value := ctx.GetText()
	*compiler.exprs = append(*compiler.exprs, Expr{number: &value})
}

func (compiler *Compiler) EnterExpr_hex(ctx *parser.Expr_hexContext) {
	value := ctx.GetText()
	*compiler.exprs = append(*compiler.exprs, Expr{hex: &value})
}

func (compiler *Compiler) EnterExpr_str(ctx *parser.Expr_strContext) {
	value := ctx.GetText()
	*compiler.exprs = append(*compiler.exprs, Expr{str: &value})
}

func (compiler *Compiler) EnterExpr_null(ctx *parser.Expr_nullContext) {
	*compiler.exprs = append(*compiler.exprs, Expr{null: true})
}

func (compiler *Compiler) EnterExpr_var(ctx *parser.Expr_varContext) {
	value := ctx.GetText()
	*compiler.exprs = append(*compiler.exprs, Expr{varName: &VarName{value}})
}

func (compiler *Compiler) EnterExpr_call(ctx *parser.Expr_callContext) {
	call := &Call{}
	*compiler.exprs = append(*compiler.exprs, Expr{call: call})
	compiler.call = call
}

func (compiler *Compiler) Print() {
	if compiler.program != nil {
		fmt.Printf("PROGRAM %s", compiler.program.String())
	}
	for _, proc := range compiler.procs {
		fmt.Printf("PROC %s", proc.String())
	}
}

func (proc *Proc) String() string {
	var str strings.Builder
	str.WriteString(proc.name)
	str.WriteString(" ")
	if proc.returnType != nil {
		str.WriteString(proc.returnType.String())
	}
	str.WriteString("\n")
	for _, stmt := range proc.statements {

		if stmt.call != nil {
			str.WriteString("\t CALL ")
			str.WriteString(stmt.call.String())
			str.WriteString("\n")
		}

		if stmt.let != nil {
			str.WriteString("\t LET ")
			str.WriteString(stmt.let.varName)
			if stmt.let.varType != nil {
				str.WriteString(" ")
				str.WriteString(stmt.let.varType.String())
			}
			for _, exp := range stmt.let.exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}

		if stmt.ret != nil {
			str.WriteString("\t RET")
			for _, exp := range stmt.ret.exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}

		if stmt.assignment != nil {
			str.WriteString("\t ASN")
			for _, varName := range stmt.assignment.vars {
				str.WriteString(" ")
				str.WriteString(varName)
			}
			for _, exp := range stmt.assignment.exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}
	}
	str.WriteString("\n")
	return str.String()
}

func (call *Call) String() string {
	var str strings.Builder
	str.WriteString(call.primary)
	if call.secondary != "" {
		str.WriteString(".")
		str.WriteString(call.secondary)
	}
	for _, exp := range call.exprs {
		str.WriteString(" ")
		str.WriteString(exp.String())
	}
	return str.String()
}

func (typ *Type) String() string {
	var str strings.Builder
	if len(typ.in) != 0 {
		str.WriteString("(")
		for idx, param := range typ.in {
			str.WriteString(param.name)
			if param.typ != nil {
				str.WriteString(" ")
				str.WriteString(param.typ.String())
			}

			for _, expr := range param.expr {
				str.WriteString(" ")
				str.WriteString(expr.String())
			}

			if idx != len(typ.in)-1 {
				str.WriteString(", ")
			}
		}
		str.WriteString(") ")
	}
	str.WriteString(typ.out.name)
	return str.String()
}

func (exp *Expr) String() string {
	if exp.boolean != nil {
		return "#bool " + *exp.boolean
	} else if exp.number != nil {
		return "#num " + *exp.number
	} else if exp.hex != nil {
		return "#hex " + *exp.hex
	} else if exp.str != nil {
		return "#str " + *exp.str
	} else if exp.null {
		return "null"
	} else if exp.varName != nil {
		return "$" + exp.varName.name
	} else if exp.call != nil {
		return "&(" + exp.call.String() + ")"
	}
	return "!!error!!"
}
