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
	Type Type
}

func (param Parameter) GetName() Name {
	return param.Name
}

func (param Parameter) GetType() Type {
	return param.Type
}

type InType struct {
	Type         Type
	Name         Name
	DefaultValue Expression
}

func NewInType(t Type, name Name, defaultValue Expression) InType {
	return InType{Type: t, Name: name, DefaultValue: defaultValue}
}

type Object interface {
	GetName() Name
	GetType() Type
}

type Procedure struct {
	Context    *Context
	Name       Name
	Parameters []Parameter
	Type       Type

	Statements []Statement
}

func NewProcedure(parent *Context, name Name) Procedure {
	return Procedure{Context: NewContext(parent), Name: name, Type: nil}
}

func (p *Procedure) AddStatement(stmt Statement) {
	p.Statements = append(p.Statements, stmt)
}

func (p *Procedure) AddStatements(stmt []Statement) {
	p.Statements = append(p.Statements, stmt...)
}

type Expression interface {
	GetType() Type
}

type Statement interface {
	GetContext() *Context
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

func (p *Source) AddProcedure(name Name, t Type) *Procedure {
	p.Procedures = append(p.Procedures, NewProcedure(&p.Context, name))
	param := &p.Procedures[len(p.Procedures)-1]
	param.Type = t
	return param
}

type Context struct {
	Parent  *Context
	Types   []Type
	Objects map[string]Object
}

func NewContext(parent *Context) *Context {
	return &Context{Parent: parent, Types: GetStandardTypes(), Objects: make(map[string]Object)}
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

func (c *Context) GetType(name string) (Type, bool) {
	for _, typ := range c.Types {
		if typ.GetOut().Value == name {
			return typ, true
		}
	}
	if c.Parent != nil {
		return c.Parent.GetType(name)
	}
	return NewCompileErrorType(), false
}
