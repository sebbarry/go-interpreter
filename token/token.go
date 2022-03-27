
/*
This package defines the 'token' struct we will 
use to define tokens during our Lexing process.
*/

package token


type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

var keywords = map[string]TokenType {
    "fn":     FUNCTION,
    "let":    LET,
    "true":   TRUE,
    "false":  FALSE,
    "return": RETURN,
    "if":     IF,
    "else":   ELSE,
}

// looks up the token to see if it is in the keywords hashmap
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {   // if in map, return constant value
        return tok
    }
    return IDENT                         // if not, return the TokenType
}


const (
    ILLEGAL   = "ILLEGAL"
    EOF       = "EOF" //this tells our parser at what point it can stop.

    IDENT     = "IDENT"
    INT       = "INT"

    ASSIGN    = "="
    PLUS      = "+"
    MINUS     = "-"
    BANG      = "!"
    ASTERISK  = "*"
    SLASH     = "/"

    COMMA     = ","
    SEMICOLON = ";"

    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    LT        = "<"
    GT        = ">"


    FUNCTION  = "FUNCTION"
    LET       = "LET"
    TRUE      = "TRUE"
    FALSE     = "FALSE"
    IF        = "IF"
    ELSE      = "ELSE"
    RETURN    = "RETURN"
    EQ        = "=="
    NOT_EQ    = "!="


)
