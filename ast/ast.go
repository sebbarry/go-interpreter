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
Some ndoes produe a statement and some produce an expression.
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
type LetStatement struct {
    Token token.Token  //the token.LET token.
    Name *Identifier  // holdes the identifier of the binding.
    Value Expression  // holds the value for the expression that produces teh value.
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }


/*
This struct holds the identifier in the LetStatement node.
*/
type Identifier struct {
    Token token.Token // the token.IDENT token 
    Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

