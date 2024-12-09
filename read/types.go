package read

type TypeOut struct {
	Name string
}

type ParamDefaultValue struct {
	VarName *VarName
	Str     *string
}

type Param struct {
	Name string
	Typ  *Type
	Expr []Expr
}

type Type struct {
	Out TypeOut
	In  []Param

	Param *Param
}

type VarName struct {
	Name string
}

type Expr struct {
	Unit    string
	VarName *VarName
	Boolean *string
	Number  *string
	Hex     *string
	Str     *string
	Null    bool
	Call    *Call
}

type Call struct {
	Primary   string
	Secondary string
	Exprs     []Expr
}

type Var struct {
	Name    string
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
}

type Proc struct {
	Name       string
	ReturnType *Type
	Statements []Statement
}
