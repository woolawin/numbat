package read

import "numbat/internal/common"

type TypeOut struct {
	Name             string
	Sequence         bool
	SequenceSize     string
	Location         common.Location
	SequenceLocation common.Location
}

func (to *TypeOut) ToName() common.Name {
	return common.Name{Value: to.Name, Location: to.Location}
}

type ParamDefaultValue struct {
	VarName *Name
	Str     *string
}

type Name struct {
	Value    string
	Location common.Location
}

func (n *Name) ToName() common.Name {
	return common.Name{Value: n.Value, Location: n.Location}
}

type Param struct {
	Name  Name
	Typ   *Type
	Value ExprGroup
}

type Type struct {
	Out TypeOut
	In  []Param

	Location common.Location

	Param *Param
}

type Expr struct {
	Unit    string
	VarName *Name
	Boolean *string
	Number  *string
	Hex     *string
	Str     *string
	Null    bool
	Call    *Call
	Seq     *ExprGroup

	Type     *Type
	Location common.Location
}

func (exp *Expr) ToExpression() common.Expression {
	if exp.VarName != nil {

	}
	return nil // TODO ADD more
}

type ExprGroup struct {
	Exprs []Expr
}

func (grp *ExprGroup) Add(expr Expr) *Expr {
	grp.Exprs = append(grp.Exprs, expr)
	return &grp.Exprs[len(grp.Exprs)-1]
}

func (grp *ExprGroup) IsEmpty() bool {
	return len(grp.Exprs) == 0
}

func (grp *ExprGroup) IsNotEmpty() bool {
	return len(grp.Exprs) != 0
}

func (grp *ExprGroup) First() *Expr {
	return &grp.Exprs[0]
}

func (grp *ExprGroup) Count() int {
	return len(grp.Exprs)
}

type Call struct {
	Primary   Name
	Secondary Name
	Arguments ExprGroup
}

type Seq struct {
	Exprs ExprGroup
}

type Var struct {
	Name    Name
	VarType *Type
	Value   ExprGroup
}

type Return struct {
	Value ExprGroup
}

type Assignment struct {
	Vars   []string
	Values ExprGroup
}

type Statement struct {
	Call       *Call
	Var        *Var
	Ret        *Return
	Assignment *Assignment

	Location common.Location
}

type Proc struct {
	Name       Name
	Type       *Type
	Statements []Statement
	Location   common.Location
}
