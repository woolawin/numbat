package c

type CSource struct {
	Main      CFunc
	Functions []CFunc
}

type CFunc struct {
	Type       string
	Name       string
	Parameters []CParameter
	Statements []CStatement
}

type CParameter struct {
	Type string
	Name string
}

type CStatement interface {
	isCStatement()
}

type CExpression interface {
	isCExpression()
}

type CLiteral struct {
	Value string
}

func (CLiteral) isCExpression() {}

type CVariable struct {
	Name string
}

func (CVariable) isCExpression() {}

type CFuncCall struct {
	FunctionName string
	Arguments    []CExpression
}

func (CFuncCall) isCExpression() {}
func (CFuncCall) isCStatement()  {}

type CVariableDeclaration struct {
	Type  string
	Name  string
	Value CExpression
}

func (CVariableDeclaration) isCStatement() {}

type CReturn struct {
	Value CExpression
}

func (CReturn) isCStatement() {}
