

package parser

import (
	"fmt"

	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)


const (
    _ int = iota
    /* 
    _ int = iota means that the following declarations have a precedence value from 1-7
    the order and the relation to each value have meaning ie. 1 + 1 is different than 1 * 1
    the values allow us to ocmpare expressions and their precedence to each other.
    */
    LOWEST
    EQUALS // ==
    LESSGREATER // > or < 
    SUM // +
    PRODUCT // *
    PREFIX // -X or !X
    CALL  // myFunction(X)
)

type (
    prefixParseFn func() ast.Expression // gets called when we encounter a token in the prefix position. 
    infixParseFn func(ast.Expression) ast.Expression // gets called when we encounter a token in the infix position 
)


// main parser struct.
type Parser struct {
    errors []string // maintaining a slice of errors to refer to.
    l *lexer.Lexer  // l is a pionter to an instance of a lexer (on which we repeatedly call NextToken() to get the next token in the input)

    //these act as the lexer does when pointing to the current and next character in the input. here they point to tokens.
    curToken token.Token    // curToken is a pointer.
    peekToken token.Token   // peekToken is a pointer.

    // these maps in place we can just check if the appropriate map (infix or prefix) has a parsing functoin associated with curToken.Type
    prefixParseFns map[token.TokenType]prefixParseFn
    infixParseFns map[token.TokenType]infixParseFn
}


// make a new parser
func New(l *lexer.Lexer) *Parser {
    p := &Parser{
        l: l,
        errors: []string{},
    }

    // read the two tokens, so curtoken and peektoken are both set
    p.nextToken()
    p.nextToken()


    p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
    p.registerPrefix(token.IDENT, p.parseIdentifier)

    return p
}

func (p *Parser) parseIdentifier() ast.Expression {
    return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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
    case token.RETURN:
        return p.parseReturnStatement()
    default:
        return p.parseExpressionStatement()
    }
}


// ************ Parser Handling Functions ************


//expression statement handler
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    stmt := &ast.ExpressionStatement{Token: p.curToken}
    stmt.Expression = p.parseExpression(LOWEST)
    if p.peekTokenIs(token.SEMICOLON) {
        p.nextToken()
    }
    return stmt
}


func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefix := p.prefixParseFns[p.curToken.Type]
    if prefix == nil {
        return nil
    }
    leftExp := prefix()
    return leftExp
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


// we always return the node that we can add to the ast.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.curToken}

    p.nextToken()

    //TODO we are skipping the expression until we encounter a semicolon

    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }
    return stmt
}


// ************ Token Handling Functions ************

func (p *Parser) curTokenIs(t token.TokenType) bool {
    return p.curToken.Type == t
}



func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

// assertion function common to parsers.
func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken() //loop to the next token
        return true
    } else {
        p.peekError(t) //add the token to the error
        return false
    }
}


// ************ Error Handling Functions ************


// getter method for the parser to return the errors in the errors slice
func (p *Parser) Errors() []string {
    return p.errors
}

//add an error to hte slice of error strings in the parser
func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token tok be %s, got %s instead", t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}


// ************ Semantic Code functions ********************** 

// These 'semantic parsing functions' are helper functoins to populate 
// the infix maps or prefix maps depending on what we discover in our
// parsing.
// ---->
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
    p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix (tokenType token.TokenType, fn infixParseFn) {
    p.infixParseFns[tokenType] = fn
}
// <----





