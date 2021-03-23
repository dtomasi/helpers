package safe_counter

import "sync"

// A concurrency safe counter
type Counter struct {
	mu sync.Mutex
	v  int
}

// Create a new counter
func NewSafeCounter() *Counter {
	return &Counter{
		v: 0,
	}
}

// Increment the counter
func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

// Get the counter value
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

