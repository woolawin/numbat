package common

type VariableDeclaration struct {
	Name  Name
	Type  Type
	Value Expression
}

func NewVariableDeclaration(name Name, t Type, val Expression) VariableDeclaration {
	return VariableDeclaration{Name: name, Type: t, Value: val}
}

func (vd *VariableDeclaration) GetName() Name {
	return vd.Name
}

func (vd *VariableDeclaration) GetType() Type {
	return vd.Type
}

type ProcedureCall struct {
	Object    Object
	Arguments []Expression
}

func NewProcedureCall(object Object, arguments []Expression) ProcedureCall {
	return ProcedureCall{Object: object, Arguments: arguments}
}
