package fakegen

import (
	"github.com/jfbramlett/faker"
	"reflect"
	"strings"
	"testing"
)

func TestSetPrice(t *testing.T) {
	SetPrice(Price{})
}

func TestCurrency(t *testing.T) {
	p, err := GetPrice().Currency(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !faker.Contains(currencies, p.(string)) {
		t.Error("Expected a currency code from currencies")
	}
}

func TestAmountWithCurrency(t *testing.T) {
	p, err := GetPrice().AmountWithCurrency(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}

	if !strings.Contains(p.(string), " ") {
		t.Error("Expected Price currency followed by a space and it's ammount")
	}
}

func TestFakeCurrency(t *testing.T) {
	p := Currency()
	if !faker.Contains(currencies, p) {
		t.Error("Expected a currency code from currencies")
	}
}

func TestFakeAmountWithCurrency(t *testing.T) {
	p := AmountWithCurrency()

	if !strings.Contains(p, " ") {
		t.Error("Expected Price currency followed by a space and it's ammount")
	}

	arrCurrency := strings.Split(p, " ")

	if !faker.Contains(currencies, arrCurrency[0]) {
		t.Error("Expected a currency code from currencies")
	}
}
