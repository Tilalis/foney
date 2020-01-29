package interpreter

import (
	"io"
	"strconv"
	"strings"
	"unicode"
)

type runePredicate func(rune) bool

// Lexer type
type Lexer struct {
	input    *strings.Reader
	finished bool
	current  rune
}

// NewLexer -- constructor for Lexer type
func NewLexer(input string) (*Lexer, error) {
	reader := strings.NewReader(input)
	current, _, err := reader.ReadRune()

	if err == io.EOF {
		return nil, ErrEmptyInput
	}

	return &Lexer{
		input:    reader,
		finished: false,
		current:  current,
	}, nil
}

// Next -- method of Lexer, returns next Token
// Returns nil after EOF
func (lexer *Lexer) Next() *Token {
	if lexer.finished {
		return nil
	}

	if lexer.current == 0 {
		lexer.finished = true
		return &Token{
			Type:  EOF,
			Value: nil,
		}
	}

	lexer.skipSpaces()

	symbol := lexer.symbol()

	if symbol != nil {
		return symbol
	}

	var tokenType int

	switch lexer.current {
	case '+':
		tokenType = PLUS
	case '-':
		tokenType = MINUS
	case '*':
		tokenType = MUL
	case '/':
		tokenType = DIV
	case '(':
		tokenType = LPAREN
	case ')':
		tokenType = RPAREN
	case ';':
		tokenType = DELIMETER
	case '\n':
		tokenType = DELIMETER
	case '=':
		tokenType = ASSIGN
	default:
		tokenType = -1
	}

	if tokenType != -1 {
		value := lexer.current
		lexer.read()
		return &Token{
			Type:  tokenType,
			Value: strconv.QuoteRune(value),
		}
	}

	return nil
}

func (lexer *Lexer) skipSpaces() {
	for lexer.current != 0 && unicode.IsSpace(lexer.current) {
		lexer.read()
	}
}

func (lexer *Lexer) read() {
	ch, _, err := lexer.input.ReadRune()
	if err == io.EOF {
		ch = 0
	}
	lexer.current = ch
}

func (lexer *Lexer) accumulateWhile(predicate runePredicate) string {
	var builder strings.Builder

	for lexer.current != 0 {
		if predicate(lexer.current) {
			builder.WriteRune(lexer.current)
			lexer.read()
		} else {
			break
		}
	}

	return builder.String()
}

func (lexer *Lexer) symbol() *Token {
	digitsPredicate := func(r rune) bool {
		return unicode.IsDigit(r) || r == '.'
	}

	digits := lexer.accumulateWhile(digitsPredicate)

	alphanumeric := lexer.accumulateWhile(func(r rune) bool {
		// unicode.Sc is Currency Symbols
		return unicode.IsLetter(r) || unicode.Is(unicode.Sc, r) || r == '_'
	})

	if digits == "" {
		digits = lexer.accumulateWhile(digitsPredicate)
	}

	if digits != "" && alphanumeric != "" {
		return &Token{
			Type:  MONEY,
			Value: digits + alphanumeric,
		}
	}

	if digits != "" {
		return &Token{
			Type:  NUMBER,
			Value: digits,
		}
	}

	if alphanumeric != "" {
		return &Token{
			Type:  SYMBOL,
			Value: alphanumeric,
		}
	}

	return nil
}
