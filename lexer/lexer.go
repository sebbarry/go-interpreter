package lexer


type Lexer struct {
    input        string // input defined as we want.
    position     int    // current position in input. (current char)
    readPosition int    // next reading position in input. (next char in queue)
    ch           byte   // actual char under examination.
}


func New(input string) *Lexer {
    l := &Lexer{input: input}
    return l
}


func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {    // check if reached the end of the input
        l.ch = 0                           // set to ascii char to NUL
    } else {
        l.ch = l.input[l.readPosition]     // otherwise read char from cur. pos.
    }
    l.position = l.readPosition            // set cur pos. to next pos.
    l.readPosition += 1                    // set next pos as next.next

}
