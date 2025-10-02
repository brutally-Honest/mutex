package counter

import "sync"

type Counter struct {
	CountSafe   int
	CountUnsafe int
	mu          sync.Mutex
}

func (ctr *Counter) IncrementSafe() {
	ctr.mu.Lock()
	defer ctr.mu.Unlock()
	ctr.CountSafe++
}

func (ctr *Counter) GetCountSafe() int {
	ctr.mu.Lock()
	defer ctr.mu.Unlock()
	return ctr.CountSafe
}

func (ctr *Counter) IncrementUnsafe() {
	ctr.CountUnsafe++
}

func (ctr *Counter) GetCountUnsafe() int {
	return ctr.CountUnsafe
}

func GetCounter() *Counter {
	return &Counter{CountSafe: 0, CountUnsafe: 0}
}
