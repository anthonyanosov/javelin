package lexer

type TokenType int

const (
	EOF TokenType = iota
	IDENT
	INT
	WALRUS
	VAR
)

var OPERATORS = map[string]TokenType{
	":=": WALRUS,
}

var KEYWORDS = map[string]TokenType{
	"var": VAR,
}

type Token struct {
	Type    TokenType
	Literal string
}
