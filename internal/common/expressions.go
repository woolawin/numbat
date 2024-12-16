package common

type VariableExpression struct {
	Name Name
	Type Type
}

func NewVariableExpression(name Name, t Type) VariableExpression {
	return VariableExpression{Name: name, Type: t}
}

func (ve *VariableExpression) GetType() Type {
	return ve.Type
}

type LiteralExpression struct {
	Value string
	Unit  string
	Type  Type
}

func NewLiteralExpression(value string, unit string, t Type) LiteralExpression {
	return LiteralExpression{Value: value, Unit: unit, Type: t}
}

func (le *LiteralExpression) GetType() Type {
	return le.Type
}

type ProceduresExpression struct {
	Call       ProcedureCall
	ReturnType *Type
}

func NewProcedureExpression(call ProcedureCall) ProceduresExpression {
	var rt *Type
	if call.Object.GetType() != nil {
		rt = &Type{Out: call.Object.GetType().Out}
	}
	return ProceduresExpression{Call: call, ReturnType: rt}
}

func (pe *ProceduresExpression) GetType() Type {
	return *pe.ReturnType
}
