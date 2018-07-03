package lexer

import "github.com/ngiyshhk/monkey/token"

type Lexer struct {
}

func New(input string) *Lexer {
	l := &Lexer{}
	return l
}

func (*Lexer) NextToken() token.Token {
	return token.Token{}
}
