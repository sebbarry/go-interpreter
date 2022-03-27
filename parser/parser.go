package parser

import (
    "interpreter/ast"
    "interpreter/token"
    "interpreter/lexer"
)


type Parser struct {
    l *lexer.Lexer  // l is a pionter to an instance of a lexer (on which we repeatedly call NextToken() to get the next token in the input)

    //these act as the lexer does when pointing to the current and next character in the input. here they point to tokens.
    curToken token.Token    // curToken is a pointer.
    peekToken token.Token   // peekToken is a pointer.
}

// make a new parser
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // read the two tokens, so curtoken and peektoken are both set
    p.nextToken()
    p.nextToken()

    return p
}

// helper to advance curToken and peekToken
func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    return nil
}

