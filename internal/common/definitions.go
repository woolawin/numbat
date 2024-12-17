package common

type ProcedureForwardDeclaration struct {
	Name Name
	Type Type
}

func NewProcedureForwardDeclaration(name Name, t Type) ProcedureForwardDeclaration {
	return ProcedureForwardDeclaration{Name: name, Type: t}
}

func (p *ProcedureForwardDeclaration) GetName() Name {
	return p.Name
}

func (p *ProcedureForwardDeclaration) GetType() Type {
	return p.Type
}

func (p *ProcedureForwardDeclaration) IsCallable() bool {
	return true
}
