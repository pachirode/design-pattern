package lockmutex

import (
	"sync"
	"testing"
)

func TestLockMutex(t *testing.T) {
	count := NewCounter()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Increment()
		}()
	}

	wg.Wait()
}
