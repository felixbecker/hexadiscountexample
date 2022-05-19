package discounter

import (
	"log"
)

type Store interface {
	Get(amout float32) float32
	Set(amount float32, rate float32)
}

type Discounter struct {
	store Store
}

func New(store Store) *Discounter {

	if store == nil {
		panic("Store cannot be nil")
	}
	d := Discounter{}
	d.store = store

	return &d
}

func (d *Discounter) Rate(amount float32) float32 {
	// sure amount is ignored ... implementation more advance... this is just for simplicity
	rate := d.store.Get(amount)
	if rate <= 0 || rate >= 1.0 {
		return 1
	}
	log.Printf("Rate from store: %f", rate)
	return rate
}
