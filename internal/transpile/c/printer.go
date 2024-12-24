package c

import "strings"

type CSourcePrinter struct {
	indentation int
	builder     strings.Builder
}

func NewCSourcePrinter() CSourcePrinter {
	return CSourcePrinter{indentation: 0, builder: strings.Builder{}}
}

func (printer *CSourcePrinter) reset() {
	printer.indentation = 0
	printer.builder.Reset()
}

func (printer *CSourcePrinter) Print(src *CSource) string {
	printer.newline()
	for _, cfunc := range src.Functions {
		printer.forwardFunctionDeclaration(&cfunc)
	}
	printer.newline()
	printer.function(&src.Main)
	for _, cfunc := range src.Functions {
		printer.function(&cfunc)
	}
	return printer.builder.String()
}

func (printer *CSourcePrinter) function(cfunc *CFunc) {
	printer.newline()
	printer.write(cfunc.Type, " ")
	printer.write(cfunc.Name, "(")
	for idx, param := range cfunc.Parameters {
		printer.write(param.Type, " ")
		printer.write(param.Name)

		if idx < len(cfunc.Parameters)-1 {
			printer.write(",")
		}
	}
	printer.write(")", " {")
	printer.indent()
	printer.statements(cfunc.Statements)
	printer.unindent()
	printer.write("}")
	printer.endline()
}

func (printer *CSourcePrinter) forwardFunctionDeclaration(cfunc *CFunc) {
	printer.newline()
	printer.write(cfunc.Type, " ")
	printer.write(cfunc.Name, "(")
	for idx, param := range cfunc.Parameters {
		printer.write(param.Type, " ")
		printer.write(param.Name)

		if idx < len(cfunc.Parameters)-1 {
			printer.write(",")
		}
	}
	printer.write(");")
}

func (printer *CSourcePrinter) statements(statements []CStatement) {
	printer.newline()
	for _, statement := range statements {
		printer.statement(statement)
	}
}

func (printer *CSourcePrinter) statement(statement CStatement) {
	ok := false
	switch statement.(type) {
	case CVariableDeclaration:
		printer.startline()
		ok = true
		printer.variableDeclaration(statement.(CVariableDeclaration))
	case CFuncCall:
		printer.startline()
		ok = true
		printer.functionCall(statement.(CFuncCall))
	case CReturn:
		printer.startline()
		ok = true
		printer.returnStatement(statement.(CReturn))
	}
	if ok {
		printer.write(";")
		printer.newline()
	}
}

func (printer *CSourcePrinter) variableDeclaration(statement CVariableDeclaration) {
	printer.write(statement.Type, " ")
	printer.write(statement.Name)
	if statement.Value != nil {
		printer.write(" = ")
		printer.expression(statement.Value)
	}

}

func (printer *CSourcePrinter) functionCall(call CFuncCall) {
	printer.write(call.FunctionName, "(")
	for idx, arg := range call.Arguments {
		printer.expression(arg)
		if idx < len(call.Arguments)-1 {
			printer.write(",")
		}
	}
	printer.write(")")
}

func (printer *CSourcePrinter) returnStatement(stmt CReturn) {
	printer.write("return")
	if stmt.Value != nil {
		printer.write(" ")
		printer.expression(stmt.Value)
	}
}

func (printer *CSourcePrinter) expression(expression CExpression) {
	switch expression.(type) {
	case CLiteral:
		printer.write(expression.(CLiteral).Value)
	case CFuncCall:
		printer.functionCall(expression.(CFuncCall))
	case CVariable:
		printer.write(expression.(CVariable).Name)
	}
}

func (printer *CSourcePrinter) indent() {
	printer.indentation++
}

func (printer *CSourcePrinter) unindent() {
	printer.indentation--
}

func (printer *CSourcePrinter) newline() {
	printer.builder.WriteString("\n")
}

func (printer *CSourcePrinter) write(values ...string) {
	for _, value := range values {
		printer.builder.WriteString(value)
	}
}

func (printer *CSourcePrinter) startline() {
	printer.builder.WriteString(strings.Repeat(" ", printer.indentation*4))
}

func (printer *CSourcePrinter) endline() {
	printer.builder.WriteString("\n")
}
