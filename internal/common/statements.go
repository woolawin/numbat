package common

type Statement interface {
	GetContext() *Context

	GuaranteedReturn() bool
}

// VariableDeclaration ----------------------------------------
//
// ------------------------------------------------------------

type VariableDeclaration struct {
	Name    Name
	Type    InOutType
	Value   Expression
	Context *Context
}

func NewVariableDeclaration(context *Context, name Name, t InOutType, val Expression) VariableDeclaration {
	return VariableDeclaration{Context: context, Name: name, Type: t, Value: val}
}

func (vd VariableDeclaration) GetContext() *Context {
	return vd.Context
}

func (vd *VariableDeclaration) GetName() Name {
	return vd.Name
}

func (vd *VariableDeclaration) GetType() InOutType {
	return vd.Type
}

func (vd VariableDeclaration) GuaranteedReturn() bool {
	return false
}

// ProcedureCall ----------------------------------------------
//
// ------------------------------------------------------------

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

func (pc ProcedureCall) GuaranteedReturn() bool {
	return false
}

// ReturnStatement --------------------------------------------
//
// ------------------------------------------------------------

type ReturnStatement struct {
	Context *Context
	Value   Expression
}

func NewReturnStatement(context *Context, value Expression) ReturnStatement {
	return ReturnStatement{Context: context, Value: value}
}

func (rs ReturnStatement) GetContext() *Context {
	return rs.Context
}

func (rs ReturnStatement) GuaranteedReturn() bool {
	return true
}
