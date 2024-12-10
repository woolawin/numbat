package read

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
}

func (listener *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Println("ERROR")
}

type Source struct {
	Program *Proc
	Procs   []Proc
}

func Read(code string) *Source {
	listener := NewListener()
	listener.Exec(code)
	src := listener.Source()
	return src
}

func (listener *Listener) Source() *Source {
	src := &Source{
		Program: listener.program,
		Procs:   listener.procs,
	}
	return src
}

func (src *Source) InferTypes() bool {
	return false
}

func (src *Source) String() string {
	var str strings.Builder
	if src.Program != nil {
		str.WriteString("PROGRAM")
		str.WriteString(src.Program.String())
	}
	for _, proc := range src.Procs {
		str.WriteString("PROC ")
		str.WriteString(proc.String())
	}
	return str.String()
}

func (proc *Proc) String() string {
	var str strings.Builder
	str.WriteString(proc.Name)
	str.WriteString("\n")
	if proc.ReturnType != nil {
		for _, in := range proc.ReturnType.In {
			str.WriteString("\tIN ")
			str.WriteString(in.Name)
			str.WriteString(" ")
			str.WriteString(in.Typ.String())

			for _, expr := range in.Expr {
				str.WriteString(" ")
				str.WriteString(expr.String())
			}
			str.WriteString("\n")
		}
		if proc.ReturnType.Out.Name != "" {
			str.WriteString("\tOUT ")
			str.WriteString(proc.ReturnType.Out.Name)
			str.WriteString("\n")
		}
	}
	str.WriteString("\tDO\n")
	for _, stmt := range proc.Statements {

		if stmt.Call != nil {
			str.WriteString("\tCALL ")
			str.WriteString(stmt.Call.String())
			str.WriteString("\n")
		}

		if stmt.Var != nil {
			str.WriteString("\tVAR ")
			str.WriteString(stmt.Var.Name)
			if stmt.Var.VarType != nil {
				str.WriteString(" ")
				str.WriteString(stmt.Var.VarType.String())
			}
			for _, exp := range stmt.Var.Exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}

		if stmt.Ret != nil {
			str.WriteString("\tRET")
			for _, exp := range stmt.Ret.Exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}

		if stmt.Assignment != nil {
			str.WriteString("\tASN")
			for _, varName := range stmt.Assignment.Vars {
				str.WriteString(" ")
				str.WriteString(varName)
			}
			for _, exp := range stmt.Assignment.Exprs {
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
	str.WriteString(call.Primary)
	if call.Secondary != "" {
		str.WriteString(".")
		str.WriteString(call.Secondary)
	}
	for _, exp := range call.Exprs {
		str.WriteString(" ")
		str.WriteString(exp.String())
	}
	return str.String()
}

func (typ *Type) String() string {
	var str strings.Builder
	if len(typ.In) != 0 {
		str.WriteString("(")
		for idx, param := range typ.In {
			str.WriteString(param.Name)
			if param.Typ != nil {
				str.WriteString(" ")
				str.WriteString(param.Typ.String())
			}

			for _, expr := range param.Expr {
				str.WriteString(" ")
				str.WriteString(expr.String())
			}

			if idx != len(typ.In)-1 {
				str.WriteString(", ")
			}
		}
		str.WriteString(") ")
	}
	str.WriteString(typ.Out.Name)
	return str.String()
}

func (exp *Expr) String() string {
	var str strings.Builder
	str.WriteString("#")
	str.WriteString(exp.Unit)
	str.WriteString(" ")

	if exp.Boolean != nil {
		str.WriteString(*exp.Boolean)
	} else if exp.Number != nil {
		str.WriteString(*exp.Number)
	} else if exp.Hex != nil {
		str.WriteString(*exp.Hex)
	} else if exp.Str != nil {
		str.WriteString(*exp.Str)
	} else if exp.Null {
		return "null"
	} else if exp.VarName != nil {
		return "$" + exp.VarName.Name
	} else if exp.Call != nil {
		return "&(" + exp.Call.String() + ")"
	}

	return str.String()
}
