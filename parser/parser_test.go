package parser

import (
    "testing"
    "interpreter/ast"
    "interpreter/lexer"
)


func TestLetStatements(t *Testing.T) {
    input := `
    let x = 5;
    let y = 10; 
    let foobar = 8383838;
    `
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    if program != nil {
        t.Fatal("ParseProgram() return nil")
    }
    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    test := []struct {
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


// function to test the let statement.
func  testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("s.TokenLiteral not 'let'. got=%q, s.TokenLiteral()")
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
        t.Errorf("letStmt.Name.TokenLiteral() not '%s', got=%s", name, letStmt.Name.TokenLiteral)
        return false
    }
    return true
}

