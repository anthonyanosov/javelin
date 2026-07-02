package lexer

import (
	"javelin/lexer"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []lexer.Token
	}{
		{
			"simple identifier and number",
			"x := 1",
			[]lexer.Token{
				{Type: lexer.IDENT, Literal: "x"},
				{Type: lexer.WALRUS, Literal: ":="},
				{Type: lexer.INT, Literal: "1"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
		{
			"longer identifier",
			"number := 1",
			[]lexer.Token{
				{Type: lexer.IDENT, Literal: "number"},
				{Type: lexer.WALRUS, Literal: ":="},
				{Type: lexer.INT, Literal: "1"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
		{
			"longer number",
			"x := 100",
			[]lexer.Token{
				{Type: lexer.IDENT, Literal: "x"},
				{Type: lexer.WALRUS, Literal: ":="},
				{Type: lexer.INT, Literal: "100"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
		{
			"longer identifier and number",
			"number := 100",
			[]lexer.Token{
				{Type: lexer.IDENT, Literal: "number"},
				{Type: lexer.WALRUS, Literal: ":="},
				{Type: lexer.INT, Literal: "100"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.source)
			l.Tokenize()

			if !reflect.DeepEqual(l.Tokens, tt.want) {
				t.Errorf("got %v, want %v", l.Tokens, tt.want)
			}
		})
	}
}
