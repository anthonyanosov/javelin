package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	text   string
	Tokens []Token
}

func NewLexer(input string) *Lexer {
	l := Lexer{
		text:   input,
		Tokens: []Token{},
	}

	return &l
}

func (l *Lexer) Tokenize() {
	i := 0
	text := l.text
	for i < len(text) {
		c := rune(text[i])

		// Whitespace
		if unicode.IsSpace(c) {
			i++
			continue
		}

		// Identifier
		if unicode.IsLetter(c) {
			start := i
			for i < len(text) && (unicode.IsDigit(rune(text[i])) || unicode.IsLetter(rune(text[i]))) {
				i++
			}

			sample := text[start:i]

			l.Tokens = append(l.Tokens, Token{Type: IDENT, Literal: sample})
			continue
		}

		// Integer
		if unicode.IsDigit(c) {
			start := i
			for i < len(text) && unicode.IsDigit(rune(text[i])) {
				i++
			}

			integer := text[start:i]
			l.Tokens = append(l.Tokens, Token{Type: INT, Literal: integer})
			continue
		}

		// Walrus operator
		if c == ':' {
			if i+1 < len(text) && text[i+1] == '=' {
				l.Tokens = append(l.Tokens, Token{Type: WALRUS, Literal: ":="})
				// Extra increment here to advance past the '=' in ':='
				i += 2
			}
			continue
		}

		i++
	}
	l.Tokens = append(l.Tokens, Token{Type: EOF})
}

func (l *Lexer) PrintTokens() {
	fmt.Println(l.Tokens)
}
