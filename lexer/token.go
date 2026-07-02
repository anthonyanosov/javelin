package lexer

type TokenType int

const (
	EOF TokenType = iota
	IDENT
	INT
	WALRUS
)

var OPERATORS = map[string]TokenType{
	":=": WALRUS,
}

type Token struct {
	Type    TokenType
	Literal string
}
