package parser

import (
	"luffy/src/ast"
	"luffy/src/lexer"
	"testing"
)

func TestVarStatements(t *testing.T) {
	input := `
	var x = 5;
	var y = 10;
	var foo = 23423;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Parsing return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("wrong number of statement. got %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVar(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVar(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not var. got %q", s.TokenLiteral())
		return false
	}
	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not varStatement. got %T", s)
		return false
	}
	if varStmt.Name.Value != name {
		t.Errorf("varStmt.name.value not %s, got %s", name, varStmt.Name.Value)
		return false
	}
	if varStmt.Name.TokenLiteral() != name {
		return false
	}
	return true
}
