package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	text      string
	character byte
	tokens    []Token
}

func NewLexer(input string) *Lexer {
	l := Lexer{
		text:   input,
		tokens: []Token{},
	}

	return &l
}

func (l *Lexer) Tokenize() {
	fmt.Println("Beginning tokenization...")
	fmt.Printf("Text: %s\n", l.text)
	i := 0
	text := l.text
	fmt.Printf("Lenght: %d\n", len(text))
	for i < len(text) {
		fmt.Printf("i: %d\n", i)
		c := rune(text[i])

		// Check whitespace
		if unicode.IsSpace(c) {
			i++
			continue
		}

		// Check keyword or identifier
		if unicode.IsLetter(c) {
			start := i
			for unicode.IsLetter(rune(text[i])) || unicode.IsDigit(rune(text[i])) {
				i++
			}
			sample := text[start:i]

			if value, ok := KEYWORDS[sample]; ok {
				l.tokens = append(l.tokens, Token{Type: value, Literal: sample})
			} else {
				l.tokens = append(l.tokens, Token{Type: IDENT, Literal: sample})
			}
			continue
		}

		// Check number/digit
		if unicode.IsDigit(c) {
			start := i
			for unicode.IsDigit(rune(text[i])) {
				i++
			}

			number := text[start:i]
			l.tokens = append(l.tokens, Token{Type: NUMBER, Literal: number})
		}

		// Walrus operator (not safe)
		if c == ':' {
			if text[i+1] == '=' {
				l.tokens = append(l.tokens, Token{Type: WALRUS, Literal: ":="})
			}
		}

		// Safety
		i++
	}
	l.tokens = append(l.tokens, Token{Type: EOF})
	fmt.Println("Tokenizing finished!")
}

func (l *Lexer) PrintTokens() {
	fmt.Println(l.tokens)
}
