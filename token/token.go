
/*
This package defines the token struct we will use to define tokens during our Lexing process.
*/

package token


type TokenType string

type Token struct {
    Token TokenType
    Literal string
}


const (
    ILLEGAL   = "ILLEGAL"
    EOF       = "EOF" //this tells our parser at what point it can stop.

    IDENT     = "IDENT"
    INT       = "INT"

    ASSIGN    = "="
    PLUS      = "+"

    COMMA     = ","
    SEMICOLON = ";"

    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"

    FUNCTION  = "FUNCTION"
    LET       = "LET"
)
