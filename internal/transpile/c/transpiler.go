package c

import (
	. "numbat/internal/common"
)

type CTranspiler struct {
}

func NewCTranspiler() CTranspiler {
	return CTranspiler{}
}

func (transpiler *CTranspiler) Transpile(src *Source) *CSource {
	main := transpiler.program(src.Program)
	var funcs []CFunc
	for _, procedure := range src.Procedures {
		funcs = append(funcs, transpiler.procedure(procedure))
	}
	return &CSource{Main: main, Functions: funcs}
}

func (transpiler *CTranspiler) program(procedure Procedure) CFunc {
	cfunc := CFunc{}
	cfunc.Name = "main"
	cfunc.Type = "int"
	cfunc.Parameters = []CParameter{
		{
			Name: "argc",
			Type: "int",
		},
		{
			Name: "argv",
			Type: "char**",
		},
	}
	cfunc.Statements = transpiler.statements(procedure.Statements)
	return cfunc
}

func (transpiler *CTranspiler) procedure(procedure Procedure) CFunc {
	cfunc := CFunc{}
	cfunc.Type = transpiler.atomicType(procedure.Type.Out.Type)
	cfunc.Name = MangleProcedureName(procedure.Name.Value)
	for _, param := range procedure.Parameters {
		cparam := CParameter{
			Name: MangleVariableName(param.Name.Value),
			Type: transpiler.atomicType(param.Type.Out.Type),
		}
		cfunc.Parameters = append(cfunc.Parameters, cparam)
	}
	cfunc.Statements = transpiler.statements(procedure.Statements)
	return cfunc
}

func (transpiler *CTranspiler) statements(statements []Statement) []CStatement {
	var stmts []CStatement
	for _, statement := range statements {
		stmts = append(stmts, transpiler.statement(statement))
	}
	return stmts
}

func (transpiler *CTranspiler) statement(statement Statement) CStatement {
	switch statement.(type) {
	case VariableDeclaration:
		return transpiler.variableDeclaration(statement.(VariableDeclaration))
	case ProcedureCall:
		return transpiler.procedureCall(statement.(ProcedureCall))
	}
	return nil // TODO ADD ERROR
}

func (transpiler *CTranspiler) variableDeclaration(statement VariableDeclaration) CVariableDeclaration {
	var value CExpression
	if statement.Value != nil {
		value = transpiler.expression(statement.Value)
	}
	return CVariableDeclaration{
		Name:  MangleVariableName(statement.Name.Value),
		Type:  transpiler.atomicType(statement.Type.Out.Type),
		Value: value,
	}
}

func (transpiler *CTranspiler) procedureCall(call ProcedureCall) CFuncCall {
	var args []CExpression
	for _, arg := range call.Arguments {
		args = append(args, transpiler.expression(arg))
	}
	return CFuncCall{
		FunctionName: MangleProcedureName(call.Object.GetName().Value),
		Arguments:    args,
	}
}

func (transpiler *CTranspiler) expression(expression Expression) CExpression {
	switch expression.(type) {
	case *LiteralExpression:
		return literal(expression.(*LiteralExpression))
	case *ProceduresExpression:
		return transpiler.procedureCall(expression.(*ProceduresExpression).Call)
	}
	return nil
}

func literal(value *LiteralExpression) CLiteral {
	atomic := value.GetType().Out.Type
	if atomic.GetName() == TypeBool {
		if value.Value == "true" {
			return CLiteral{"1"}
		}
		return CLiteral{"0"}
	}
	if atomic.GetName() == TypeAscii {
		return CLiteral{"'" + value.Value + "'"}
	}

	if atomic.GetName() == TypeInt32 {
		return CLiteral{value.Value}
	}

	if atomic.GetName() == TypeInt64 {
		return CLiteral{value.Value}
	}

	if atomic.GetName() == TypeFloat32 {
		return CLiteral{value.Value}
	}

	if atomic.GetName() == TypeFloat64 {
		return CLiteral{value.Value}
	}
	return CLiteral{} // record ERROR
}

func (transpiler *CTranspiler) atomicType(atomic AtomicType) string {
	if atomic.IsNoType() {
		return "void"
	}

	if atomic.GetName() == TypeBool {
		return "int"
	}

	if atomic.GetName() == TypeInt32 {
		return "int"
	}

	if atomic.GetName() == TypeInt64 {
		return "long int"
	}

	if atomic.GetName() == TypeFloat64 {
		return "double"
	}

	if atomic.GetName() == TypeFloat32 {
		return "float"
	}

	if atomic.GetName() == TypeAscii {
		return "char"
	}

	if atomic.GetName() == TypeByte {
		return "char"
	}
	return "" // TODO record error
}
