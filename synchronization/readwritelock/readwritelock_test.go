package readwritelock

import (
	"sync"
	"testing"
)

func TestReadWriteLock(t *testing.T) {
	data := Data{data: 0, rwLock: &sync.RWMutex{}}

	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println("Read: ", data.Read())
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		data.Write(1000000000)
	}()

	wg.Wait()
}
