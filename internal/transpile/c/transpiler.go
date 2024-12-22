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
	transpiler.program(src.Program)
	for _, procedure := range src.Procedures {
		transpiler.procedure(procedure)
	}
}

func (transpiler *CTranspiler) String() string {
	return transpiler.builder.String()
}

func (transpiler *CTranspiler) program(procedure Procedure) {
	transpiler.newline()
	transpiler.write("int main(int argc, char** argv) {")
	transpiler.endline()
	transpiler.indent()
	transpiler.statements(procedure.Statements)
	transpiler.unindent()
	transpiler.write("}")
	transpiler.endline()
}

func (transpiler *CTranspiler) procedure(procedure Procedure) {
	transpiler.newline()
	transpiler.atomicType(procedure.Type.Out.Type)
	transpiler.write(MangleProcedureName(procedure.Name.Value), "(")
	for idx, param := range procedure.Parameters {
		transpiler.write(param.Name.Value, " ")
		transpiler.atomicType(param.Type.Out.Type)
		if idx < len(procedure.Parameters)-1 {
			transpiler.write(",")
		}
	}
	transpiler.write(")", " {")
	transpiler.indent()
	transpiler.statements(procedure.Statements)
	transpiler.unindent()
	transpiler.write("}")
	transpiler.endline()
}

func (transpiler *CTranspiler) statements(statements []Statement) {
	for _, statement := range statements {
		transpiler.statement(statement)
	}
}

func (transpiler *CTranspiler) statement(statement Statement) {

	switch statement.(type) {
	case VariableDeclaration:
		transpiler.variableDeclaration(statement.(VariableDeclaration))
	case ProcedureCall:
		transpiler.procedureCall(statement.(ProcedureCall))
	}
}

func (transpiler *CTranspiler) variableDeclaration(statement VariableDeclaration) {
	transpiler.startline()
	transpiler.atomicType(statement.Type.Out.Type)
	transpiler.write(MangleVariableName(statement.Name.Value))
	if statement.Value != nil {
		transpiler.write(" ", "=", " ")
		transpiler.expression(statement.Value)
	}
	transpiler.write(";")
	transpiler.endline()
}

func (transpiler *CTranspiler) procedureCall(call ProcedureCall) {
	transpiler.write(MangleProcedureName(call.Object.GetName().Value))
	transpiler.write("(")
	for idx, arg := range call.Arguments {
		transpiler.expression(arg)
		if idx < len(call.Arguments)-1 {
			transpiler.write(",")
		}
	}
	transpiler.write(")")
}

func (transpiler *CTranspiler) expression(expression Expression) {
	switch expression.(type) {
	case *LiteralExpression:
		transpiler.write(literal(expression.(*LiteralExpression)))
	case *ProceduresExpression:
		transpiler.procedureCall(expression.(*ProceduresExpression).Call)
	}
}

func literal(value *LiteralExpression) string {
	atomic := value.GetType().Out.Type
	if atomic.GetName() == TypeBool {
		if value.Value == "true" {
			return "1"
		}
		return "0"
	}
	if atomic.GetName() == TypeAscii {
		return "'" + value.Value + "'"
	}

	if atomic.GetName() == TypeInt32 {
		return value.Value
	}

	if atomic.GetName() == TypeInt64 {
		return value.Value
	}

	if atomic.GetName() == TypeFloat32 {
		return value.Value
	}

	if atomic.GetName() == TypeFloat64 {
		return value.Value
	}
	return ""
}

func (transpiler *CTranspiler) atomicType(atomic AtomicType) {
	if atomic.IsNoType() {
		transpiler.write("void", " ")
	}

	if atomic.GetName() == TypeBool {
		transpiler.write("int", " ")
	}

	if atomic.GetName() == TypeInt32 {
		transpiler.write("int", " ")
	}

	if atomic.GetName() == TypeInt64 {
		transpiler.write("long int", " ")
	}

	if atomic.GetName() == TypeFloat64 {
		transpiler.write("double", " ")
	}

	if atomic.GetName() == TypeFloat32 {
		transpiler.write("float", " ")
	}

	if atomic.GetName() == TypeAscii {
		transpiler.write("char", " ")
	}

	if atomic.GetName() == TypeByte {
		transpiler.write("char", " ")
	}
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
	for _, value := range values {
		transpiler.builder.WriteString(value)
	}
}

func (transpiler *CTranspiler) startline() {
	transpiler.builder.WriteString(strings.Repeat(" ", transpiler.indentation*4))
}

func (transpiler *CTranspiler) endline() {
	transpiler.builder.WriteString("\n")
}
