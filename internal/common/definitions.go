package common

type ProcedureDefinition struct {
	Name Name
	Type *Type
}

func NewProcedureDefinition(name Name, t *Type) ProcedureDefinition {
	return ProcedureDefinition{Name: name, Type: t}
}

func (p *ProcedureDefinition) GetName() Name {
	return p.Name
}

func (p *ProcedureDefinition) GetType() *Type {
	return p.Type
}

func (p *ProcedureDefinition) IsCallable() bool {
	return true
}
