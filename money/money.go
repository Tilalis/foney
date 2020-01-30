package money

import "math"

import "fmt"

// Money represents money
type Money struct {
	// TODO: Change that to something better
	value    float64
	currency *Currency
}

// New constructs Money
func New(value float64, currency *Currency) *Money {
	return &Money{
		value:    round(value, 2),
		currency: currency,
	}
}

// Add adds
func (m *Money) Add(other *Money) (*Money, error) {
	converted, err := other.Convert(m.currency)

	if err != nil {
		return nil, err
	}

	return &Money{
		value:    round(m.value+converted.value, 2),
		currency: m.currency,
	}, nil
}

// Sub substitutes
func (m *Money) Sub(other *Money) (*Money, error) {
	converted, err := other.Convert(m.currency)

	if err != nil {
		return nil, err
	}

	return &Money{
		value:    round(m.value-converted.value, 2),
		currency: m.currency,
	}, nil
}

// Mul multiplies
func (m *Money) Mul(value float64) (*Money, error) {
	return &Money{
		value:    round(m.value*value, 2),
		currency: m.currency,
	}, nil
}

// Div divides
func (m *Money) Div(value float64) (*Money, error) {
	if value == 0 {
		return nil, ErrDivisionByZero
	}

	return &Money{
		value:    round(m.value/value, 2),
		currency: m.currency,
	}, nil
}

// Convert returns new Money converted to currency
func (m *Money) Convert(currency *Currency) (*Money, error) {
	if m.currency == currency {
		return m, nil
	}

	exchangeRate, err := GetExchangeRate(m.currency, currency)
	if err != nil {
		return nil, err
	}

	multiplied, _ := m.Mul(exchangeRate)
	return multiplied, nil
}

func (m *Money) String() string {
	return fmt.Sprintf("%.2f%s", m.value, m.currency.String())
}

func round(x float64, points int) float64 {
	helper := math.Pow(10.0, float64(points))
	return math.Round(x*helper) / helper
}
