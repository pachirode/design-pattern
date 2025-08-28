package main

import "time"

func asyncOperation() <-chan int {
	result := make(chan int)

	go func() {
		time.Sleep(time.Second)
		result <- 42
	}()

	return result
}
