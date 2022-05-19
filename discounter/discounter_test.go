package discounter_test

import (
	"testing"

	"github.com/felixbecker/hexadiscountexample/discounter"
)

type fakeStore struct {
	GetFunc func(amount float32) float32
	Setfunc func(amount float32, rate float32)
}

func (f *fakeStore) Get(amount float32) float32 {
	return f.GetFunc(amount)
}
func (f *fakeStore) Set(amount float32, rate float32) {
	f.Setfunc(amount, rate)
}

func Test_Discounter_should_return_the_given_rate(t *testing.T) {

	var setFuncWasNotCalled bool

	f := fakeStore{}
	f.Setfunc = func(amount, rate float32) {
		setFuncWasNotCalled = true
	}
	f.GetFunc = func(amount float32) float32 {

		return 0.2
	}

	disc := discounter.New(&f)
	result := disc.Rate(200)

	if setFuncWasNotCalled == true {
		t.Errorf("This should not happen")
	}
	if result != 0.2 {
		t.Errorf("Expected the result to be 0.2; got: %f", result)
	}
}

func Test_Discounter_should_return_1_when_the_rate_is_zero(t *testing.T) {
	var setFuncWasNotCalled bool

	f := fakeStore{}
	f.Setfunc = func(amount, rate float32) {
		setFuncWasNotCalled = true
	}
	f.GetFunc = func(amount float32) float32 {

		return 0
	}

	disc := discounter.New(&f)
	result := disc.Rate(200)

	if setFuncWasNotCalled == true {
		t.Errorf("This should not happen")
	}
	if result != 1 {
		t.Errorf("Expected the result to be 1; got: %f", result)
	}

}

func Test_Discounter_should_return_1_when_the_rate_lt_zero(t *testing.T) {
	var setFuncWasNotCalled bool

	f := fakeStore{}
	f.Setfunc = func(amount, rate float32) {
		setFuncWasNotCalled = true
	}
	f.GetFunc = func(amount float32) float32 {

		return -2.0
	}

	disc := discounter.New(&f)
	result := disc.Rate(200)

	if setFuncWasNotCalled == true {
		t.Errorf("This should not happen")
	}
	if result != 1 {
		t.Errorf("Expected the result to be 1; got: %f", result)
	}

}

func Test_Discounter_should_return_1_when_the_rate_gt_1(t *testing.T) {

	var setFuncWasNotCalled bool

	f := fakeStore{}
	f.Setfunc = func(amount, rate float32) {
		setFuncWasNotCalled = true
	}
	f.GetFunc = func(amount float32) float32 {

		return 2.0
	}

	disc := discounter.New(&f)
	result := disc.Rate(200)

	if setFuncWasNotCalled == true {
		t.Errorf("This should not happen")
	}
	if result != 1 {
		t.Errorf("Expected the result to be 1; got: %f", result)
	}
}
