package storeprovider

type InMemoryProvider struct {
	rate float32
}

func NewInMemory() *InMemoryProvider {

	return &InMemoryProvider{rate: 0.22}
}

func (im *InMemoryProvider) Get(amount float32) float32 {
	return im.rate
}
func (im *InMemoryProvider) Set(amount float32, rate float32) {
	im.rate = amount
}
