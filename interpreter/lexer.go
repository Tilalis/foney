package interpreter

import (
	"errors"
	"fmt"
	money "foney/money"

	"io"
	"strconv"
	"strings"
	"unicode"
)

type runePredicate func(rune) bool

// Lexer type
type Lexer struct {
	reader   io.RuneReader
	finished bool
	current  rune
	previous rune
}

// NewLexer -- constructor for Lexer type
func NewLexer(reader io.RuneReader) (*Lexer, error) {
	current, _, err := reader.ReadRune()

	if errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("Error: %w", ErrEmptyInput)
	}

	return &Lexer{
		reader:   reader,
		finished: false,
		current:  current,
	}, nil
}

// Next -- method of Lexer, returns next Token
// Returns nil after EOF
func (lexer *Lexer) Next() (*Token, error) {
	if lexer.finished {
		return nil, nil
	}

	if lexer.current == 0 {
		lexer.finished = true
		return &Token{
			Type:  EOF,
			Value: nil,
		}, nil
	}

	lexer.skipSpaces()

	symbol, err := lexer.symbol()

	if err != nil {
		return nil, err
	}

	if symbol != nil {
		return symbol, nil
	}

	var tokenType TokenType

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
		}, nil
	}

	return nil, nil
}

func (lexer *Lexer) skipSpaces() {
	for lexer.current != 0 && unicode.IsSpace(lexer.current) {
		lexer.read()
	}
}

func (lexer *Lexer) read() {
	ch, _, err := lexer.reader.ReadRune()
	if errors.Is(err, io.EOF) {
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

func (lexer *Lexer) symbol() (*Token, error) {
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
		amount, currencyName := digits, alphanumeric

		value, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return nil, err
		}

		currency, err := money.GetCurrencyByName(currencyName)
		if errors.Is(err, money.ErrBadCurrencyName) {
			currency, err = money.GetCurrencyByAlias(currencyName)
			if err != nil {
				return nil, err
			}
		}

		return &Token{
			Type:  MONEY,
			Value: money.New(value, currency),
		}, nil
	}

	if digits != "" {
		value, err := strconv.ParseFloat(digits, 64)
		if err != nil {
			return nil, err
		}

		return &Token{
			Type:  NUMBER,
			Value: value,
		}, nil
	}

	if alphanumeric != "" {
		return &Token{
			Type:  SYMBOL,
			Value: alphanumeric,
		}, nil
	}

	return nil, nil
}
