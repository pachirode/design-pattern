package main

import "testing"

func TestFuturePromise(t *testing.T) {
	future := asyncOperation()

	value := <-future
	println("Async operation result: ", value)
}
