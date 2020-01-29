package interpreter

import "fmt"

const (
	NUMBER = iota
	MONEY  = iota

	PLUS  = iota
	MINUS = iota
	MUL   = iota
	DIV   = iota

	LPAREN = iota
	RPAREN = iota

	SYMBOL = iota
	ASSIGN = iota

	DELIMETER = iota
	EOF       = iota
)

type Token struct {
	Type  int
	Value interface{}
}

func (t *Token) String() string {
	typeName := map[int]string{
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
