package main

type (
	Poper struct {
	}
)

func (b *Bitq) Push(queue string, data []byte) error {
	return nil
}

func (b *Bitq) Pop(queue string) (*Poper, error) {
	return nil, nil
}

func (p *Poper) SetMaxCache(size int) {
}

func (p *Poper) CacheSize() int {
	return 0
}

func (p *Poper) OutputCh() <- chan *Msg {
	return nil
}

func (p *Poper) ExitCh() <- chan struct{} {
	return nil
}

func (p *Poper) Close() error {
	return nil
}
