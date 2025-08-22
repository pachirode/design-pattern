package boundedparallelism

import "time"

func Worker(id int, semaphore chan struct{}) {
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	println("Worker ", id)
	time.Sleep(2 * time.Second)
}
