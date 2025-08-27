package main

func fibonacci(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		for i := 0; i < n; i++ {
			ch <- a
			a, b = b, a+b
		}
	}()

	return ch
}

func main() {
	for i := range fibonacci(10) {
		println(i)
	}
}
