package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	FUNCTION  = "FUNCTION"
	VAR       = "VAR"
	BANG      = "BANG"
	STAR      = "STAR"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	RETURN    = "RETURN"
	IF        = "IF"
	ELSE      = "ELSE"
	DIFF      = "DIFF"
	EQUAL     = "EQUAL"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupID(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return IDENT
}
