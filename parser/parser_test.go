package parser

import (
    "testing"
    "interpreter/ast"
    "interpreter/lexer"
)



func TestReturnStatements(t *testing.T) {
    input := `
    return 5;
    return 10; 
    return 123091309;
    `

    l := lexer.New(input)           //(1) Lex the input
    p := New(l)                     //(2) Make a parser
    program := p.ParseProgram()     //(3) Parser the program.

    checkParserErrors(t, p)     // test function

    //erorr handling
    if program == nil {
        t.Fatal("ParseProgram() return nil")
    }
    //checking we have three return staements in the inputs above.
    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    for _, stmt := range program.Statements {
        returnStmt, ok := stmt.(*ast.ReturnStatement);
        if !ok {
            t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
            continue
        }
        if returnStmt.TokenLiteral() != "return" {
            t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
        }
    }
}

func TestLetStatements(t *testing.T) {
    input := `
    let x = 5;
    let y = 10; 
    let foobar = 8383838;
    `

    l := lexer.New(input)           //(1) Lex the input
    p := New(l)                     //(2) Make a parser
    program := p.ParseProgram()     //(3) Parser the program.

    checkParserErrors(t, p)     // test function

    //erorr handling
    if program == nil {
        t.Fatal("ParseProgram() return nil")
    }

    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    tests := []struct {
        expectedIdenfitier string
    }{
        {"x"},
        {"y"},
        {"foobar"},
    }
    for i, tt := range tests {
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdenfitier) {
            return
        }
    }

}


//function to test for errors in our parser.
func checkParserErrors(t *testing.T, p *Parser)  {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }
    t.Errorf("Parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}




// function to test the let statement.
func  testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("s.TokenLiteral not 'let'. got=%q, s.TokenLiteral()", s.TokenLiteral())
        return false
    }
    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s not *ast.LetStatement. got=%T", s)
        return false
    }
    if letStmt.Name.Value != name {
        t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
        return false
    }
    if letStmt.Name.TokenLiteral() != name {
        t.Errorf("letStmt.Name.TokenLiteral() not '%s', got=%s", name, letStmt.Name.TokenLiteral())
        return false
    }
    return true
}

