package c

import (
	. "numbat/internal/common"
	"strings"
)

type CTranspiler struct {
	indentation int
	builder     strings.Builder
}

func NewCTranspiler() CTranspiler {
	return CTranspiler{indentation: 0, builder: strings.Builder{}}
}

func (transpiler *CTranspiler) Transpile(src *Source) {
	for _, procedure := range src.Procedures {
		transpiler.procedure(procedure)
	}
}

func (transpiler *CTranspiler) String() string {
	return transpiler.builder.String()
}

func (transpiler *CTranspiler) indent() {
	transpiler.indentation++
}

func (transpiler *CTranspiler) unindent() {
	transpiler.indentation--
}

func (transpiler *CTranspiler) newline() {
	transpiler.builder.WriteString("\n")
}

func (transpiler *CTranspiler) write(values ...string) {
	transpiler.builder.WriteString(strings.Repeat(" ", transpiler.indentation))
	for _, value := range values {
		transpiler.builder.WriteString(value)
		transpiler.builder.WriteRune(' ')
	}
}

func (transpiler *CTranspiler) writenl(values ...string) {
	transpiler.builder.WriteString(strings.Repeat(" ", transpiler.indentation))
	for idx, value := range values {
		transpiler.builder.WriteString(value)
		if idx != len(values)-1 {
			transpiler.builder.WriteRune(' ')
		} else {
			transpiler.newline()
		}
	}
}

func (transpiler *CTranspiler) procedure(procedure Procedure) {
	transpiler.newline()
	transpiler.atomicType(procedure.Type.Out.Type)
	transpiler.write(MangleProcedureName(procedure.Name.Value), "(")
	for idx, param := range procedure.Parameters {
		transpiler.write(param.Name.Value)
		transpiler.atomicType(param.Type.Out.Type)
		if idx < len(procedure.Parameters)-1 {
			transpiler.write(",")
		}
	}
	transpiler.writenl(")", "{")
	//transpiler.indent()
	//transpiler.unindent()
	transpiler.write("}")
}

func (transpiler *CTranspiler) atomicType(atomic AtomicType) {
	if atomic.IsNoType() {
		transpiler.write("void")
	}

	if atomic.GetName() == TypeBool {
		transpiler.write("int")
	}

	if atomic.GetName() == TypeInt32 {
		transpiler.write("int")
	}

	if atomic.GetName() == TypeInt64 {
		transpiler.write("long int")
	}

	if atomic.GetName() == TypeFloat64 {
		transpiler.write("double")
	}

	if atomic.GetName() == TypeFloat32 {
		transpiler.write("float")
	}

	if atomic.GetName() == TypeAscii {
		transpiler.write("char")
	}

	if atomic.GetName() == TypeByte {
		transpiler.write("char")
	}
}
