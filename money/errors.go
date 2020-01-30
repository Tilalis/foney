package money

import (
	"errors"
)

// Base Errors
var (
	ErrBadCurrencyName = errors.New("Bad currency name")
	ErrNoExchangeRate  = errors.New("No exchange rate")
	ErrDivisionByZero  = errors.New("Division by zezo")
)
