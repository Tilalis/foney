package money

import "fmt"

import "strings"

// Currency represents currency
type Currency struct {
	name  string
	alias string
}

const allowedChars = "ABCDEFGHIJKLMNOPQRSTUVDXYZ"

var currencies = map[string]*Currency{}

// GetCurrency -- currency constructor
func GetCurrency(name, alias string) (*Currency, error) {
	// Check name to be valid
	// Valid name is of length 3
	if len(name) != 3 {
		return nil, ErrWrongCurrencyName
	}

	// Valid name contains only uppercase latin characters
	for _, r := range name {
		if !strings.ContainsRune(allowedChars, r) {
			return nil, ErrWrongCurrencyName
		}
	}

	if currency, ok := currencies[name]; ok {
		return currency, nil
	}

	if currency, ok := currencies[alias]; ok {
		return currency, nil
	}

	currency := &Currency{
		name:  name,
		alias: alias,
	}

	if name == "" {
		name = alias
	}

	currencies[name] = currency

	return currency, nil
}

// GetName returns currency name
func (currency *Currency) GetName() string {
	return currency.name
}

// GetAlias returns currency alias
func (currency *Currency) GetAlias() string {
	return currency.alias
}

func (currency *Currency) String() string {
	return currency.name
}

var exchangeRates = map[string]float64{
	"USD_BYN": 2.13,
	"BYN_USD": 0.469,
}

// GetExchangeRate returns exchange rate between two currencies
func GetExchangeRate(from *Currency, to *Currency) (float64, error) {
	fromTo := fmt.Sprintf("%s_%s", from.name, to.name)

	if value, ok := exchangeRates[fromTo]; ok {
		return value, nil
	}

	// TODO: Add getting exchangerate from https://www.currconv.com/
	return 0, ErrNoExchangeRate
}

// SetExchangeRate sets exchange rate
func SetExchangeRate(from *Currency, to *Currency, value float64) error {
	fromTo := fmt.Sprintf("%s_%s", from.name, to.name)
	toFrom := fmt.Sprintf("%s_%s", to.name, from.name)

	exchangeRates[fromTo] = value
	exchangeRates[toFrom] = 1 / value

	return nil
}
