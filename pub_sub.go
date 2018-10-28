package main

type (
	Suber struct {
	}
)

func (b *Bitq) Pub(subject string, data []byte) error {
	return nil
}

func (b *Bitq) Sub(subject string) (*Suber, error) {
	return nil, nil
}

func (s *Suber) OutputCh() <- chan *Msg {
	return nil
}

func (s *Suber) ExitCh() <- chan struct{} {
	return nil
}

func (s *Suber) SetMaxCache(size int) {
}

func (s *Suber) CacheSize() int {
	return 0
}

func (s *Suber) Close() error {
	return nil
}