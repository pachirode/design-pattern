package main

import (
	"fmt"
	"time"
)

func printA() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("A")
	}
}

func printB() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("B")
	}
}

func main() {
	go printA()
	go printB()

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}
