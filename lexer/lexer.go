package lexer

import "interpreter/token"

type Lexer struct {
    input        string // input defined as we want.
    position     int    // current position in input. (current char)
    readPosition int    // next reading position in input. (next char in queue)
    ch           byte   // actual char under examination.
}


func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}


//reads the current character of the lexer
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {    // check if reached the end of the input
        l.ch = 0                           // set to ascii char to NULL
    } else {
        l.ch = l.input[l.readPosition]     // otherwise read char from cur. pos.
    }
    l.position = l.readPosition            // set cur pos. to next pos.
    l.readPosition += 1                    // set next pos as next.next
}


// this determines what the token is at the current index of the input value
func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    l.skipWhiteSpace()                      // skip the whitespace of the character (if it is)
    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:                                // default value - reads the char. and determiens if we have a char.
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()          // look up the literal value
            tok.Type = token.LookupIdent(tok.Literal) //look up the tokens type
            return tok
        } else if isDigit(l.ch) {                   // check if the token is a digit value (not a character)
            tok.Type = token.INT                    // assign type to INT
            tok.Literal = l.readNumber()            // assign literal to the readNumber() output
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }
    l.readChar()
    return tok
}


func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

// determines if the value is a Latin digit of 0-9. We could also add another function to parser
// octal, floats, hex etc.
func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9' //return true if the ch is ch is 0 <= ch <= 9
}

// read the current identifier of the lexer
func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}


// function to determine if the character at the lexer position is a letter
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' // _ allows for snake case of vars.
                                                                         // we can also allow ? or ! .etc here.
}

// function to skip the whitespace of the current character if it is
// this is sometimes called 'eatwhitespace'
func (l *Lexer) skipWhiteSpace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}
