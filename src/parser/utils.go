package parser

import (
	"token"
	"fmt"
)


func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) currentTokenTypeIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}


func (p *Parser) peekTokenTypeIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}


func (p *Parser) expectPeekType(t token.TokenType) bool {
	if p.peekTokenTypeIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}


func (p *Parser) peekError(tokType token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", tokType, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}