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
var aliases = map[string]string{}

// GetCurrency -- currency getter
func GetCurrency(name, alias string) (*Currency, error) {
	// Check name to be valid
	// Valid name is of length 3
	if name != "" {
		if len(name) != 3 {
			return nil, fmt.Errorf("%w: '%s'", ErrBadCurrencyName, name)
		}

		// Valid name contains only uppercase latin characters
		for _, r := range name {
			if !strings.ContainsRune(allowedChars, r) {
				return nil, fmt.Errorf("%w: '%s'", ErrBadCurrencyName, name)
			}
		}
	}

	if name == "" {
		var ok bool
		name, ok = aliases[alias]

		if !ok {
			return nil, fmt.Errorf("%w: '%s'", ErrBadCurrencyName, alias)
		}
	}

	if currency, ok := currencies[name]; ok {
		return currency, nil
	}

	return newCurrency(name, alias), nil
}

// GetCurrencyByName Returns Currency By Name
func GetCurrencyByName(name string) (*Currency, error) {
	return GetCurrency(name, "")
}

// GetCurrencyByAlias Returns Currency By Alias
func GetCurrencyByAlias(alias string) (*Currency, error) {
	return GetCurrency("", alias)
}

func newCurrency(name, alias string) *Currency {
	if alias == "" {
		alias = name
	} else {
		aliases[alias] = name
	}

	currency := &Currency{
		name:  name,
		alias: alias,
	}

	currencies[name] = currency

	return currency
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
	return 0, fmt.Errorf("%w: %s to %s", ErrNoExchangeRate, from.name, to.name)
}

// SetExchangeRate sets exchange rate
func SetExchangeRate(from *Currency, to *Currency, value float64) error {
	fromTo := fmt.Sprintf("%s_%s", from.name, to.name)
	toFrom := fmt.Sprintf("%s_%s", to.name, from.name)

	exchangeRates[fromTo] = value
	exchangeRates[toFrom] = 1 / value

	return nil
}

// Currencies
var (
	USD = newCurrency("USD", "$")
	EUR = newCurrency("EUR", "€")
	RUB = newCurrency("RUB", "")
	BYN = newCurrency("BYN", "Br")
)
