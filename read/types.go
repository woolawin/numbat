package read

type TypeOut struct {
	name string
}

type ParamDefaultValue struct {
	varName *VarName
	str     *string
}

type Param struct {
	name string
	typ  *Type
	expr []Expr
}

type Type struct {
	out TypeOut
	in  []Param

	param *Param
}

type VarName struct {
	name string
}

type Expr struct {
	varName *VarName
	boolean *string
	number  *string
	hex     *string
	str     *string
	null    bool
	call    *Call
}

type Call struct {
	primary   string
	secondary string
	exprs     []Expr
}

type Let struct {
	varName string
	varType *Type
	exprs   []Expr
}

type Return struct {
	exprs []Expr
}

type Assignment struct {
	vars  []string
	exprs []Expr
}

type Statement struct {
	call       *Call
	let        *Let
	ret        *Return
	assignment *Assignment
}

type Proc struct {
	name       string
	returnType *Type
	statements []Statement
}
