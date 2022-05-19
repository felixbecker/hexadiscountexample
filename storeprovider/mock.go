package storeprovider

type MockProvider struct {
	GetFunc func(amount float32) float32
	SetFunc func(amount float32, rate float32)
}

func NewMockProvider() {}

func (mp *MockProvider) Set(amount float32, rate float32) {
	if mp.SetFunc == nil {
		panic("SetFunc on Mock not provided")
	}
	mp.SetFunc(amount, rate)
}
func (mp *MockProvider) Get(amount float32) float32 {
	if mp.GetFunc == nil {
		panic("SetFunc on Mock not provided")
	}
	return mp.GetFunc(amount)
}
