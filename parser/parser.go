package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
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

// parse the incoming stream
func (p *Parser) ParseProgram() *ast.Program {

    program := &ast.Program{}
    program.Statements = []ast.Statement{} // slice of statements.

    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            // add to statement slice
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken() // loop to the next token 
        return nil
    }
    return program
}


/*
switch between different types of statements.
*/
func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
    case token.LET:
        return p.parseLetStatement()
    default:
        return nil
    }
}



// let statement parser function
func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken}

    if !p.expectPeek(token.IDENT) {
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeek(token.ASSIGN) {
        return nil
    }

    // while the current token is not a ';'
    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }

    return stmt

}



func (p *Parser) curTokenIs(t token.TokenType) bool {
    return p.curToken.Type == t
}



func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

// assertion function common to parsers.
func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        return false
    }
}
