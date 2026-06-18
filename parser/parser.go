package parser

import (
	"fmt"
	"javelin/ast"
	"javelin/lexer"
	"strconv"
)

type Parser struct {
	Tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	p := Parser{
		Tokens: tokens,
		pos:    0,
	}

	return &p
}

func (p *Parser) current() lexer.Token {
	return p.Tokens[p.pos]
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) Parse() ast.Program {
	var program ast.Program
	fmt.Println(p.Tokens)
	for p.current().Type != lexer.EOF {
		stmt := p.parseStatement()
		program = append(program, stmt)
	}
	return program
}

func (p *Parser) parseStatement() ast.Stmt {
	token := p.current()
	if token.Type == lexer.VAR {
		p.advance()
		return p.parseVarStatement()
	}
	return nil
}

func (p *Parser) parseVarStatement() *ast.VarStmt {
	var ident string
	if p.current().Type != lexer.IDENT {
		panic("Expected an identifier after 'var'")
	}
	ident = p.current().Literal
	p.advance()
	// TODO: Wrap things like INT, and later types like strings in a wrapper
	if p.current().Type != lexer.WALRUS && p.current().Type != lexer.INT {
		panic("Expected assignment or type after identifier")
	}
	p.advance()
	value, err := strconv.ParseInt(p.current().Literal, 10, 64)
	if err != nil {
		panic("Integer variable was assigned a non-integer value")
	}
	p.advance()
	// TODO: We probably need to store variable type as well, especially if we want to transpile to C for the codegen
	return &ast.VarStmt{Ident: ident, Integer: value}
}
