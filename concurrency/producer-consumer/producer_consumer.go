package main

import (
	"fmt"
	"time"
)

func producer(id int, ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Producer %d: %d\n", id, i)
		ch <- i
	}
}

func consumer(id int, ch chan int) {
	for {
		time.Sleep(time.Second * 3)
		i := <-ch
		fmt.Printf("Consumer %d: %d\n", id, i)
	}
}

func main() {
	ch := make(chan int)
	go producer(1, ch)
	go consumer(1, ch)
	go consumer(2, ch)
	time.Sleep(time.Second * 20)
}
