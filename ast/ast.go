/*

-- terminology: 

prefix operator: --5 

postfix operator: foobar++ 

infix operator - infix operators appear in binary expressions - where the operator has two operands : 5 * 5  
--infix operators fall under the scope of operator precendence. An altrenative term for this is operator precedence. (5*5+10)


*/


/*

We ahve three types of statements: 
1. let
2. return 
3. expression statements: 

- expression statements: let x = 10; x + 10;
                                     |-> this is the shorthand way of updating our let statement (through an expression.)

*/


package ast

import (
    "interpreter/token"
    "bytes"
)

/*
these nodes are used to construct the AST (abstract syntax tree)
*/
type Node interface {
    TokenLiteral() string     // returns a token literal value the node is associated with.
    String() string
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
func (p *Program) String() string { // write out a buffer of strings.
    var out bytes.Buffer
    for _, s := range p.Statements { //get statements from the []slice
        out.WriteString(s.String())
    }
    return out.String()
}



///// Nodes for the ast.

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
func (ls *LetStatement) String() string {
    var out bytes.Buffer
    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")
    if ls.Value != nil {
        out.WriteString(ls.Value.String())
    }
    out.WriteString(";")
    return out.String()
}


func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
    return i.Value
}


//////


/////
//return statement struct.
type ReturnStatement struct {
    Token token.Token         // the 'return' token
    ReturnValue Expression    // return statements return an expression.
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
    var out bytes.Buffer
    out.WriteString(rs.TokenLiteral() + " ")
    if rs.ReturnValue != nil {
        out.WriteString(rs.ReturnValue.String())
    }
    out.WriteString(";")
    return out.String()
    // return <value>;
}
//////



//// expression statement struct. 
type ExpressionStatement struct {
    Token token.Token // teh frist token fo the expression.
    Expression Expression // the expression we have tied to this expression statement.
}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}



