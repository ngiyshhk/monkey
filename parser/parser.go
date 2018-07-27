package parser

import (
	"github.com/ngiyshhk/monkey/ast"
	"github.com/ngiyshhk/monkey/lexer"
	"github.com/ngiyshhk/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// curTokenとpeekTokenのセットのため2回呼ぶ
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
