package lexer

import (
	"luffy/src/token"
	"testing"
)

type expected_token struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func do_test(t *testing.T, tests []expected_token, l *Lexer) {
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func Test1(t *testing.T) {
	input := `var five = 5;
	var ten = 10;
	var add = fn(x, y) {
	x + y;
	};
	var result = add(five, ten);
	`

	tests := []expected_token{
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
	do_test(t, tests, l)
}

func Test2(t *testing.T) {
	input := `
	!*/<> true false return if else
	!= ==
	`
	tests := []expected_token{
		{token.BANG, "!"},
		{token.STAR, "*"},
		{token.SLASH, "/"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.RETURN, "return"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.DIFF, "!="},
		{token.EQUAL, "=="},
	}
	l := New(input)
	do_test(t, tests, l)
}
