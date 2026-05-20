package lexer

type TokenType int

const (
	EOF TokenType = iota
	IDENT
	INT
	WALRUS
	VAR
	NUMBER
	EQUALS
	PLUS
	MINUS
	PRODUCT
	QUOTIENT
)

var OPERATORS = map[string]TokenType{
	":=": WALRUS,
	"=":  EQUALS,
	"+":  PLUS,
	"-":  MINUS,
	"*":  PRODUCT,
	"/":  QUOTIENT,
}

var KEYWORDS = map[string]TokenType{
	"var": VAR,
}

type Token struct {
	Type    TokenType
	Literal string
}
