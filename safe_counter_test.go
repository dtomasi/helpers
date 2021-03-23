package helpers

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewSafeCounter(t *testing.T) {
	var wg sync.WaitGroup

	cCount := 1000
	c := NewSafeCounter()

	for i := 0; i < cCount; i++ {
		// Add a waitgroup delta ...
		wg.Add(1)
		go func() {
			// This block simulates concurrent writes to SafeCounter
			//
			// Create a seed based on time nano
			rand.Seed(time.Now().UnixNano())
			// Get a random integer between 0 and 10
			n := rand.Intn(10)
			// Sleep for the given integer in nano seconds
			time.Sleep(time.Duration(n)*time.Nanosecond)
			// Increment the counter
			c.Inc()
			// All done
			wg.Done()
		}()
	}

	// Wait for routines to increment
	wg.Wait()

	if c.Value() != cCount {
		t.Errorf("expected to %d actual %d", cCount, c.Value())
	}
}
