package application

import (
	"log"
	"math"
)

type Application struct {
	discounter Discounter
}
type Discounter interface {
	Rate(amount float32) float32
}

func New(discounter Discounter) *Application {

	if discounter == nil {
		panic("Discounter cannot be nil")
	}
	return &Application{
		discounter: discounter,
	}
}

func (a *Application) Discount(amount float32) float32 {

	rate := a.discounter.Rate(amount)
	log.Printf("Rate from discounter: %f", rate)
	return float32(math.Round(float64(amount*rate*100) / 100))
}
