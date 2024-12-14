package read

import "numbat/internal/common"

type TypeOut struct {
	Name     string
	Location common.Location
}

type ParamDefaultValue struct {
	VarName *Name
	Str     *string
}

type Name struct {
	Value    string
	Location common.Location
}

type Param struct {
	Name Name
	Typ  *Type
	Expr []Expr
}

type Type struct {
	Out TypeOut
	In  []Param

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

	Location common.Location
}

type Call struct {
	Primary   Name
	Secondary Name
	Exprs     []Expr
}

type Var struct {
	Name    Name
	VarType *Type
	Exprs   []Expr
}

type Return struct {
	Exprs []Expr
}

type Assignment struct {
	Vars  []string
	Exprs []Expr
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
