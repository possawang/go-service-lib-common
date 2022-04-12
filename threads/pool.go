package threads

import "sync"

type Pool struct {
	sync.Mutex
	max    uint64
	active uint64
}

func (p *Pool) Block() int {
	i := 0
	for p.active > 0 {
		i++
		if i == 10 {
			i = 0
		}
	}
	return i
}

func (p *Pool) OnMax() int {
	i := 0
	for p.active >= p.max {
		i++
		if i == 10 {
			i = 0
		}
	}
	return i
}

func (p *Pool) Exec(execution func()) {
	go func() {
		p.Lock()
		p.active++
		p.Unlock()
		execution()
		p.Lock()
		p.active--
		p.Unlock()
	}()
}

func NewPool(max uint64) Pool {
	return Pool{
		max:    max,
		active: 0,
	}
}
