package c

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

type CFuncCall struct {
	FunctionName string
	Arguments    []CExpression
}

func (CFuncCall) isCExpression() {}

type CVariableDeclaration struct {
	Type  string
	Name  string
	Value CExpression
}

func (CVariableDeclaration) isCStatement() {}

type CSource struct {
	Main      CFunc
	Functions []CFunc
}
