package parser

import (
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
	for p.current().Type != lexer.EOF {
		stmt := p.parseStatement()
		program = append(program, stmt)
	}
	return program
}

func (p *Parser) parseStatement() ast.Stmt {
	token := p.current()
	if token.Type == lexer.IDENT {
		// TODO: We will later need to differentiate between declaration and assignment. For now, we're doing just declaration
		return p.parseDeclStatement()
	}
	return nil
}

func (p *Parser) parseDeclStatement() *ast.DeclStmt {
	// First grab the name of the identifier to use in the AST node
	ident := p.current().Literal
	p.advance()

	// Check that the next token is a walrus operator
	if p.current().Type != lexer.WALRUS {
		panic("Expected an assignment operator after identifier")
	}
	p.advance()

	// Grab the value
	if p.current().Type != lexer.INT {
		panic("Expected an integer value after assignment")
	}
	value, err := strconv.ParseInt(p.current().Literal, 10, 64)
	if err != nil {
		panic("Integer variable was assigned a non-integer value")
	}
	p.advance()

	// Return an AST node with the identifier and the value
	return &ast.DeclStmt{Ident: ident, Integer: value}

	// TODO: Next task is to write some tests before moving on
}
