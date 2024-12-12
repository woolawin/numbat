package common

type Location struct {
	File   string
	Line   int
	Column int
}

func (token *Location) Nullify() {
	token.File = ""
	token.Line = 0
	token.Column = 0
}
