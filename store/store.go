package store

type Provider interface {
	Set(amount float32, rate float32)
	Get(amount float32) float32
}

type Store struct {
	provider Provider
}

func New(adapter Provider) *Store {

	return &Store{
		provider: adapter,
	}
}

func (s *Store) Set(amount float32, rate float32) {

	s.provider.Set(amount, rate)
}
func (s *Store) Get(amount float32) float32 {
	return s.provider.Get(amount)

}
