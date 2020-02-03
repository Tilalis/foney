package compiler

import "fmt"

// TokenType type
type TokenType int

// Token Types
const (
	NUMBER TokenType = iota
	MONEY  TokenType = iota

	PLUS  TokenType = iota
	MINUS TokenType = iota
	MUL   TokenType = iota
	DIV   TokenType = iota

	LPAREN TokenType = iota
	RPAREN TokenType = iota

	SYMBOL TokenType = iota
	ASSIGN TokenType = iota

	DELIMETER TokenType = iota
	EOF       TokenType = iota
)

// Token represents Token
type Token struct {
	Type  TokenType
	Value interface{}
}

func (t *Token) String() string {
	typeName := map[TokenType]string{
		NUMBER:    "NUMBER",
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
	}[t.Type]

	return fmt.Sprintf("Token(%s, %v)", typeName, t.Value)
}
