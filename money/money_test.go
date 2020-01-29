package money

import (
	"fmt"
	"testing"
)

var usd, _ = GetCurrency("USD", "$")
var byn, _ = GetCurrency("BYN", "Br")

var _ = SetExchangeRate(usd, byn, 2)

func TestNewPrecision(t *testing.T) {
	money := NewMoney(
		12.345457547547,
		usd,
	)

	checkPrecision := fmt.Sprintf("%.3f", money.value)
	if checkPrecision != "12.350" {
		t.Errorf("TestNew: Expected Money.value to be 12.350, got %s", checkPrecision)
	}
}

func TestAdd(t *testing.T) {
	tenUsd := NewMoney(10.0, usd)
	twentyByn := NewMoney(20.0, byn)

	result, err := tenUsd.Add(twentyByn)

	if err != nil {
		t.Errorf("TestAdd: Got error: %v", err)
	}

	if result.value != 20.0 {
		t.Errorf("Got %s, expected 20.00USD", result)
	}

}

func TestSub(t *testing.T) {
	twentyByn := NewMoney(23.446, byn)
	tenUsd := NewMoney(10.0, usd)

	result, err := twentyByn.Sub(tenUsd)

	if err != nil {
		t.Errorf("TestSub: Got error: %v", err)
	}

	if result.value != 3.45 {
		t.Errorf("Got %s, expected 3.45BYN", result)
	}
}
