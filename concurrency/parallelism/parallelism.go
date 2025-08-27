package main

func fibonacci(n int, ch chan int) {
	defer close(ch)

	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
}

func main() {
	ch := make(chan int)

	go fibonacci(10, ch)

	for num := range ch {
		println(num)
	}
}
