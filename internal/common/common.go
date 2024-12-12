package common

type Token struct {
	File   string
	Line   int
	Column int
}

func (token *Token) Nullify() {
	token.File = ""
	token.Line = 0
	token.Column = 0
}
