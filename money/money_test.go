package money

import (
	"fmt"
	"testing"
)

var nonExistentCurrency, _ = GetCurrency("NEC", "")

var _ = SetExchangeRate(USD, BYN, 2)

func TestString(t *testing.T) {
	money := NewMoney(12.45, USD)
	value := money.String()
	if value != "12.45USD" {
		t.Errorf("Got %s, expected 12.45USD", value)
		return
	}

	t.Logf("[%v].String() == %s", money, value)
}

func TestNew(t *testing.T) {
	money := NewMoney(12.345457547547, USD)

	checkPrecision := fmt.Sprintf("%.3f", money.value)
	if checkPrecision != "12.350" {
		t.Errorf("TestNew: Expected Money.value to be 12.350, got %s", checkPrecision)
		return
	}

	t.Logf("NewMoney(12.345457547546, usd) == %v", money)
}

// Test Mul too
func TestConvert(t *testing.T) {
	money := NewMoney(10, USD)
	converted, err := money.Convert(BYN)

	if err != nil {
		t.Errorf("Got error %v", err)
		return
	}

	if converted.value != 20 && converted.currency != BYN {
		t.Errorf("Got %v, expected value to be 20 and currency to be BYN", converted)
		return
	}

	t.Logf("[%v].Convert(byn) == %v", money, converted)
}

func TestConvertError(t *testing.T) {
	money := NewMoney(10, USD)
	converted, err := money.Convert(nonExistentCurrency)

	if converted != nil && err == nil {
		t.Errorf("Got return value instead of error: %v", converted)
		return
	}

	if err != ErrNoExchangeRate {
		t.Errorf("Got error other than ErrNoExchangeRate: %v", err)
		return
	}

	t.Logf("Success. Got ErrNoExchangeRate: '%v'", err)
}

func TestDiv(t *testing.T) {
	money := NewMoney(10, USD)
	divided, err := money.Div(2.5)

	if err != nil {
		t.Errorf("Got error: '%v'", err)
		return
	}

	if divided.value != 4 {
		t.Errorf("Got %v, expected 4.00USD", divided)
		return
	}

	t.Logf("10.00USD / 2.5 == %v", divided)
}

func TestDivError(t *testing.T) {
	money := NewMoney(10, USD)
	divided, err := money.Div(0)

	if divided != nil && err == nil {
		t.Errorf("Got non-nil result and no Error. Result: %v, Error: %v", divided, err)
		return
	}

	if err != ErrDivisionOnZero {
		t.Errorf("Got error other than ErrDivisionOnZero: %v", err)
		return
	}

	t.Logf("Success. Got ErrDivisionOnZero: '%v'", err)
}

func TestAdd(t *testing.T) {
	tenUsd := NewMoney(10.0, USD)
	twentyByn := NewMoney(20.0, BYN)

	result, err := tenUsd.Add(twentyByn)

	if err != nil {
		t.Errorf("TestAdd: Got error: %v", err)
		return
	}

	if result.value != 20.0 {
		t.Errorf("Got %s, expected 20.00USD", result)
		return
	}

	t.Logf("%v + %v == %v", tenUsd, twentyByn, result)

}

func TestAddError(t *testing.T) {
	tenUsd := NewMoney(10.0, USD)
	twentySomething := NewMoney(20.0, nonExistentCurrency)

	result, err := tenUsd.Add(twentySomething)

	if result != nil && err == nil {
		t.Errorf("Got non-nil result and no Error. Result: %v, Error: %v", result, err)
		return
	}

	if err != ErrNoExchangeRate {
		t.Errorf("Got error other than ErrNoExchangeRate: %v", err)
		return
	}

	t.Logf("Success. Got ErrNoExchangeRate: '%v'", err)
}

func TestSub(t *testing.T) {
	twentyByn := NewMoney(23.446, BYN)
	tenUsd := NewMoney(10.0, USD)

	result, err := twentyByn.Sub(tenUsd)

	if err != nil {
		t.Errorf("TestSub: Got error: %v", err)
		return
	}

	if result.value != 3.45 {
		t.Errorf("Got %s, expected 3.45BYN", result)
		return
	}

	t.Logf("%v - %v == %v", twentyByn, tenUsd, result)
}

func TestSubError(t *testing.T) {
	tenUsd := NewMoney(10.0, USD)
	twentySomething := NewMoney(20.0, nonExistentCurrency)

	result, err := tenUsd.Sub(twentySomething)

	if result != nil && err == nil {
		t.Errorf("Got non-nil result and no Error. Result: %v, Error: %v", result, err)
		return
	}

	if err != ErrNoExchangeRate {
		t.Errorf("Got error other than ErrNoExchangeRate: %v", err)
		return
	}

	t.Logf("Success. Got ErrNoExchangeRate: '%v'", err)
}

func TestConvertSame(t *testing.T) {
	tenUsd := NewMoney(10, USD)
	converted, err := tenUsd.Convert(USD)

	if converted != tenUsd {
		t.Errorf("Converting to same currency did not return same value! %v != %v", tenUsd, converted)
		return
	}

	if err != nil {
		t.Errorf("Got error '%v'", err)
		return
	}

	t.Logf("Success. addr %v == addr %v", tenUsd, converted)
}

func TestGetCurrencyByName(t *testing.T) {
	usd, err := GetCurrencyByName("USD")

	if err != nil {
		t.Errorf("Got error '%v'", err)
		return
	}

	if usd != USD {
		t.Errorf("Got currency with different address in memory: %v != %v", usd, USD)
		return
	}

	t.Logf("Success. %v == %v", usd, USD)
}

func TestGetCurrencyByAlias(t *testing.T) {
	usd, err := GetCurrencyByAlias("$")

	if err != nil {
		t.Errorf("Got error '%v'", err)
		return
	}

	if usd != USD {
		t.Errorf("Got currency with different address in memory: %v != %v", usd, USD)
		return
	}

	t.Logf("Success. %v == %v", usd, USD)
}

func TestCurrencyTooLongNameError(t *testing.T) {
	badCurrency, err := GetCurrency("BADCUR", "alias")

	if badCurrency != nil && err == nil {
		t.Errorf("Got non-nil result and no Error. Result: %v, Error: %v", badCurrency, err)
		return
	}

	if err != ErrBadCurrencyName {
		t.Errorf("Got error other than ErrBadCurrencyName: %v", err)
		return
	}

	t.Logf("Success. Got ErrBadCurrencyName: '%v'", err)
}

func TestCurrencyBadRunesError(t *testing.T) {
	badCurrency, err := GetCurrency("B1D", "alias")

	if badCurrency != nil && err == nil {
		t.Errorf("Got non-nil result and no Error. Result: %v, Error: %v", badCurrency, err)
		return
	}

	if err != ErrBadCurrencyName {
		t.Errorf("Got error other than ErrBadCurrencyName: %v", err)
		return
	}

	t.Logf("Success. Got ErrBadCurrencyName: '%v'", err)
}

func TestGetNameAndAlias(t *testing.T) {
	name := USD.GetName()
	alias := USD.GetAlias()

	if name != "USD" {
		t.Errorf("Got %s, expected 'USD'", name)
		return
	}

	if alias != "$" {
		t.Errorf("Got %s, expected '$'", alias)
		return
	}

	t.Logf("Succes. (%s, %s)", name, alias)
}
