package read

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

func Read(code string) *Source {
	input := antlr.NewInputStream(code)
	lexer := parser.NewNumbatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewNumbatParser(stream)

	errorListener := &ErrorListener{}
	p.AddErrorListener(errorListener)

	tree := p.Prog()

	src := NewSource()
	antlr.ParseTreeWalkerDefault.Walk(src, tree)
	return src
}

type Source struct {
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

func NewSource() *Source {
	return &Source{
		BaseNumbatListener: new(parser.BaseNumbatListener),
	}
}

func (source *Source) SetType(t *Type) {
	source.typeChain = append(source.typeChain, t)
	source.typ = t
}

func (source *Source) UnsetType() {
	if len(source.typeChain) == 0 {
		return
	}

	if len(source.typeChain) == 1 {
		source.typ = source.typeChain[0]
		source.typeChain = nil
		return
	}

	source.typeChain = source.typeChain[:len(source.typeChain)-1]
	source.typ = source.typeChain[len(source.typeChain)-1]
}

func (source *Source) String() string {
	var str strings.Builder
	if source.program != nil {
		str.WriteString("PROGRAM")
		str.WriteString(source.program.String())
	}
	for _, proc := range source.procs {
		str.WriteString("PROC ")
		str.WriteString(proc.String())
	}
	return str.String()
}

func (proc *Proc) String() string {
	var str strings.Builder
	str.WriteString(proc.name)
	str.WriteString("\n")
	if proc.returnType != nil {
		for _, in := range proc.returnType.in {
			str.WriteString("\tIN ")
			str.WriteString(in.name)
			str.WriteString(" ")
			str.WriteString(in.typ.String())

			for _, expr := range in.expr {
				str.WriteString(" ")
				str.WriteString(expr.String())
			}
			str.WriteString("\n")
		}
		if proc.returnType.out.name != "" {
			str.WriteString("\tOUT ")
			str.WriteString(proc.returnType.out.name)
			str.WriteString("\n")
		}
	}
	str.WriteString("\tDO\n")
	for _, stmt := range proc.statements {

		if stmt.call != nil {
			str.WriteString("\tCALL ")
			str.WriteString(stmt.call.String())
			str.WriteString("\n")
		}

		if stmt.let != nil {
			str.WriteString("\tLET ")
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
			str.WriteString("\tRET")
			for _, exp := range stmt.ret.exprs {
				str.WriteString(" ")
				str.WriteString(exp.String())
			}
			str.WriteString("\n")
		}

		if stmt.assignment != nil {
			str.WriteString("\tASN")
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
	var str strings.Builder
	str.WriteString("#")
	str.WriteString(exp.unit)
	str.WriteString(" ")

	if exp.boolean != nil {
		str.WriteString(*exp.boolean)
	} else if exp.number != nil {
		str.WriteString(*exp.number)
	} else if exp.hex != nil {
		str.WriteString(*exp.hex)
	} else if exp.str != nil {
		str.WriteString(*exp.str)
	} else if exp.null {
		return "null"
	} else if exp.varName != nil {
		return "$" + exp.varName.name
	} else if exp.call != nil {
		return "&(" + exp.call.String() + ")"
	}

	return str.String()
}
