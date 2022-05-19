package application_test

import (
	"testing"

	"github.com/felixbecker/hexadiscountexample/application"
)

type fakeDiscounter struct {
	RateFunc func(amount float32) float32
}

func (f *fakeDiscounter) Rate(amount float32) float32 {
	return f.RateFunc(amount)
}

func Test_Calculate_a_discount_with_a_rate_of_zero_dot_five(t *testing.T) {

	fake := fakeDiscounter{
		RateFunc: func(amount float32) float32 {
			return 0.5
		},
	}
	app := application.New(&fake)

	result := app.Discount(200)
	if result != 100 {
		t.Errorf("Expected the result to be 100; got: %f", result)
	}
}
