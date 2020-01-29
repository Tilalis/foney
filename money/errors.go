package money

import "errors"

// ErrWrongCurrencyName -- Wrong Currency Name
var ErrWrongCurrencyName = errors.New("Error: wrong currency name")

// ErrNoExchangeRate -- No Exchange Rate
var ErrNoExchangeRate = errors.New("Error: no exchange rate")

// ErrCantConvertToSameCurrency -- Can't convery to same currency
var ErrCantConvertToSameCurrency = errors.New("Error: can't convert to same currency")

// ErrDivisionOnZero -- Division on zero
var ErrDivisionOnZero = errors.New("Error: division on zezo")
