package barrier

import (
	"sync"
	"testing"
)

func TestBarrier(t *testing.T) {
	var barrier sync.WaitGroup = sync.WaitGroup{}

	barrier.Add(3)

	for i := 0; i < 3; i++ {
		go Worker(i, &barrier)
	}

	barrier.Wait()

}
