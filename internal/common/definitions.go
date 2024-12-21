package common

type ProcedureForwardDeclaration struct {
	Name Name
	Type InOutType
}

func NewProcedureForwardDeclaration(name Name, t InOutType) ProcedureForwardDeclaration {
	return ProcedureForwardDeclaration{Name: name, Type: t}
}

func (p *ProcedureForwardDeclaration) GetName() Name {
	return p.Name
}

func (p *ProcedureForwardDeclaration) GetType() InOutType {
	return p.Type
}

func (p *ProcedureForwardDeclaration) IsCallable() bool {
	return true
}
