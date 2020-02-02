package interpreter

import "fmt"

// TokenType type
type TokenType int

// Token Types
const (
	EOF TokenType = iota

	NUMBER
	MONEY

	PLUS
	MINUS
	MUL
	DIV

	LPAREN
	RPAREN

	SYMBOL
	ASSIGN

	DELIMETER
)

// Token represents Token
type Token struct {
	Type  TokenType
	Value interface{}
}

func (t *Token) String() string {
	typeName := map[TokenType]string{
		MONEY:     "MONEY",
		PLUS:      "PLUS",
		MINUS:     "MINUS",
		MUL:       "MUL",
		DIV:       "DIV",
		LPAREN:    "LPAREN",
		RPAREN:    "RPAREN",
		SYMBOL:    "SYMBOL",
		ASSIGN:    "ASSIGN",
		DELIMETER: "DELIMETER",
		EOF:       "EOF",
		NUMBER:    "NUMBER",
	}[t.Type]

	return fmt.Sprintf("Token(%s, %v)", typeName, t.Value)
}
