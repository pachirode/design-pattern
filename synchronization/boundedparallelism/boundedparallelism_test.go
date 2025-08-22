package boundedparallelism

import (
	"sync"
	"testing"
)

func TestBoundedParallelism(t *testing.T) {
	numWorkers := 5
	boundedParallelism := 2
	semaphore := make(chan struct{}, boundedParallelism)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			Worker(i, semaphore)
		}()
	}

	wg.Wait()
}
