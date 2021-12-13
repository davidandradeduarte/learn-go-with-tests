package sync

import "sync"

type Counter struct {
	// Inc   func()
	// Value func() int
	mu sync.Mutex
	v  int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v++
}

func (c *Counter) Value() int {
	return c.v
}
