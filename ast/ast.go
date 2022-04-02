package ast 

import (
    "interpreter/token"
)

/*
these nodes are used to construct the AST (abstract syntax tree)
*/
type Node interface {
    TokenLiteral() string     // returns a token literal value the node is associated with.
}


/*

-- Some nodes produe a statement and some produce an expression.

    [1] Statements: declares variables without returning anything. no stack frame associated with it. 
    [2] Expressions: creates a new value by comibning a function and other statements. returns something.

*/
type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

/*
This program node is used as the root node of every AST
*/
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}



/////

// let x = 5
// x  is the identifier (identifiers do not produce values)
// 5 is the expression (expressions produce values)
//let statement struct
type LetStatement struct {
    Token token.Token  //the token.LET token.
    Name *Identifier  // holdes the identifier of the binding.
    Value Expression  // holds the value for the expression that produces teh value.
}
type Identifier struct {   //This struct holds the identifier in the LetStatement node.
    Token token.Token // the token.IDENT token 
    Value string
}
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

//////


/////
//return statement struct.
type ReturnStatement struct {
    Token token.Token         // the 'return' token
    ReturnValue Expression    // return statements return an expression.
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }





