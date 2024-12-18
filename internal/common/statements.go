package common

type VariableDeclaration struct {
	Name    Name
	Type    Type
	Value   Expression
	Context *Context
}

func NewVariableDeclaration(context *Context, name Name, t Type, val Expression) VariableDeclaration {
	return VariableDeclaration{Context: context, Name: name, Type: t, Value: val}
}

func (vd VariableDeclaration) GetContext() *Context {
	return vd.Context
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
	Context   *Context
}

func NewProcedureCall(context *Context, object Object, arguments []Expression) ProcedureCall {
	return ProcedureCall{Context: context, Object: object, Arguments: arguments}
}

func (pc ProcedureCall) GetContext() *Context {
	return pc.Context
}
