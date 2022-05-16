package application

import "math"

type Application struct {
	discounter Discounter
}
type Discounter interface {
	Rate(amount float32) float32
}

func New(discounter Discounter) *Application {

	return &Application{
		discounter: discounter,
	}
}

func (a *Application) Discount(amount float32) float32 {

	rate := a.discounter.Rate(amount)

	return float32(math.Round(float64(amount*rate*100)/100)) * rate
}
