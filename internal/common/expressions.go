package common

type Expression interface {
	GetType() InOutType
}

type VariableExpression struct {
	Name Name
	Type InOutType
}

func NewVariableExpression(name Name, t InOutType) VariableExpression {
	return VariableExpression{Name: name, Type: t}
}

func (ve *VariableExpression) GetType() InOutType {
	return ve.Type
}

type LiteralExpression struct {
	Value string
	Unit  string
	Type  InOutType
}

func NewLiteralExpression(value string, unit string, t InOutType) LiteralExpression {
	return LiteralExpression{Value: value, Unit: unit, Type: t}
}

func (le LiteralExpression) GetType() InOutType {
	return le.Type
}

type ProceduresExpression struct {
	Call       ProcedureCall
	ReturnType InOutType
}

func NewProcedureExpression(call ProcedureCall) ProceduresExpression {

	return ProceduresExpression{Call: call, ReturnType: NewInOutType(nil, call.Object.GetType().Out)}
}

func (pe *ProceduresExpression) GetType() InOutType {
	return pe.ReturnType
}

type SequenceExpression struct {
	Size   int
	Values []Expression
	Type   InOutType
}

func NewSequenceExpression(size int, values []Expression, typ InOutType) SequenceExpression {
	return SequenceExpression{
		Size:   size,
		Values: values,
		Type:   typ,
	}
}

func (se SequenceExpression) GetType() InOutType {
	return se.Type
}
