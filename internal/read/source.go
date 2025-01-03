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
	Programs []Proc
	Procs    []Proc
}

func (reader *SourceReader) Source() *Source {
	src := &Source{
		Programs: reader.programs,
		Procs:    reader.procs,
	}
	return src
}

func (src *Source) String() string {
	var str strings.Builder
	for _, prog := range src.Programs {
		str.WriteString("PROGRAM")
		str.WriteString(prog.String())
	}
	for _, proc := range src.Procs {
		str.WriteString("PROC ")
		str.WriteString(proc.String())
	}
	return str.String()
}

func (proc *Proc) String() string {
	var str strings.Builder
	str.WriteString(proc.Name.Value)
	str.WriteString("\n")
	if proc.Type != nil {
		for _, in := range proc.Type.In {
			str.WriteString("\tIN ")
			str.WriteString(in.Name.Value)
			str.WriteString(" ")
			str.WriteString(in.Typ.String())

			for _, expr := range in.Value.Exprs {
				str.WriteString(" ")
				str.WriteString(expr.String())
			}
			str.WriteString("\n")
		}
		if proc.Type.Out.Name != "" {
			str.WriteString("\tOUT ")
			str.WriteString(proc.Type.Out.Name)
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
			str.WriteString(stmt.Var.Name.Value)
			if stmt.Var.VarType != nil {
				str.WriteString(" ")
				str.WriteString(stmt.Var.VarType.String())
			}
			for _, exp := range stmt.Var.Value.Exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}

		if stmt.Ret != nil {
			str.WriteString("\tRET")
			for _, exp := range stmt.Ret.Value.Exprs {
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
			for _, exp := range stmt.Assignment.Values.Exprs {
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
	str.WriteString(call.Primary.Value)
	if call.Secondary.Value != "" {
		str.WriteString(".")
		str.WriteString(call.Secondary.Value)
	}
	for _, exp := range call.Arguments.Exprs {
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
			str.WriteString(param.Name.Value)
			if param.Typ != nil {
				str.WriteString(" ")
				str.WriteString(param.Typ.String())
			}

			for _, expr := range param.Value.Exprs {
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
	if typ.Out.Sequence {
		str.WriteString("[")
		str.WriteString(typ.Out.SequenceSize)
		str.WriteString("]")
	}
	return str.String()
}

func (exp *Expr) String() string {
	var str strings.Builder
	if exp.Boolean != nil {
		str.WriteString(*exp.Boolean)
	} else if exp.Number != nil {
		str.WriteString(*exp.Number)
		if exp.Unit != "" {
			str.WriteString("#")
			str.WriteString(exp.Unit)
		}
	} else if exp.Hex != nil {
		str.WriteString(*exp.Hex)
	} else if exp.Str != nil {
		str.WriteString(*exp.Str)
	} else if exp.Null {
		return "null"
	} else if exp.VarName != nil {
		return "$" + exp.VarName.Value
	} else if exp.Call != nil {
		return "&(" + exp.Call.String() + ")"
	} else if exp.Seq != nil {
		str.WriteString("[")
		for idx, val := range exp.Seq.Exprs {
			str.WriteString(val.String())
			if idx != len(exp.Seq.Exprs)-1 {
				str.WriteString(" ")
			}
		}
		str.WriteString("]")
	}

	return str.String()
}
