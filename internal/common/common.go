package common

type Location struct {
	File   string
	Line   int
	Column int
}

type Name struct {
	Value    string
	Location Location
}

type Parameter struct {
	Name Name
	Type InOutType
}

func (param Parameter) GetName() Name {
	return param.Name
}

func (param Parameter) GetType() InOutType {
	return param.Type
}

type Object interface {
	GetName() Name
	GetType() InOutType
}

type Procedure struct {
	Context    *Context
	Name       Name
	Parameters []Parameter
	Type       InOutType

	Statements []Statement
}

func NewProcedure(parent *Context, name Name) Procedure {
	return Procedure{Context: NewContext(parent), Name: name, Type: NoInOutType()}
}

func (p *Procedure) AddStatement(stmt Statement) {
	p.Statements = append(p.Statements, stmt)
}

func (p *Procedure) AddStatements(stmt []Statement) {
	p.Statements = append(p.Statements, stmt...)
}

func (p *Procedure) AddParameter(param Parameter) {
	p.Parameters = append(p.Parameters, param)
}

type Expression interface {
	GetType() InOutType
}

type Source struct {
	Context    Context
	Program    Procedure
	Procedures []Procedure
}

func NewSource() *Source {
	return &Source{
		Context: Context{
			Parent:  nil,
			Types:   GetStandardTypes(),
			Objects: make(map[string]Object),
		},
	}
}

func (p *Source) AddProcedure(name Name, t InOutType) *Procedure {
	p.Procedures = append(p.Procedures, NewProcedure(&p.Context, name))
	procedure := &p.Procedures[len(p.Procedures)-1]
	procedure.Type = t
	procedure.Context.ReturnType = t.Out
	procedure.Context.CurrentObjectName = name.Value
	return procedure
}

type Context struct {
	Parent  *Context
	Types   []AtomicType
	Objects map[string]Object

	InProgram         bool
	CurrentObjectName string
	ReturnType        SuperAtomicType
}

func NewContext(parent *Context) *Context {
	return &Context{
		Parent:            parent,
		Types:             GetStandardTypes(),
		Objects:           make(map[string]Object),
		InProgram:         parent.InProgram,
		CurrentObjectName: parent.CurrentObjectName,
		ReturnType:        parent.ReturnType,
	}
}

func NewProgramContext(sourceContext *Context) *Context {
	return &Context{
		Parent:            sourceContext,
		Types:             GetStandardTypes(),
		Objects:           make(map[string]Object),
		InProgram:         true,
		CurrentObjectName: "program",
		ReturnType:        NewSuperAtomicType(NewNoType()),
	}
}

func (c *Context) AddObject(name string, object Object) {
	c.Objects[name] = object
}

func (c *Context) GetObject(name string) (Object, bool) {
	obj, ok := c.Objects[name]
	if !ok {
		if c.Parent == nil {
			return nil, false
		}
		return c.Parent.GetObject(name)
	}
	return obj, true
}

func (c *Context) GetType(name string) (AtomicType, bool) {
	for _, typ := range c.Types {
		if typ.GetName() == name {
			return typ, true
		}
	}
	if c.Parent != nil {
		return c.Parent.GetType(name)
	}
	return NewCompileErrorType(), false
}
