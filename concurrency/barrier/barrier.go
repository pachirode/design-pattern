package barrier

import "sync"

func Worker(id int, barrier *sync.WaitGroup) {
	defer barrier.Done()

	println("Worker ", id)

	for i := 0; i < 5; i++ {
		println("Worker ", id, " is processing task ", i)
	}
}
