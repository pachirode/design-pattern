package main

func main() {
	pushCh := make(chan int)
	pullCh := make(chan int)

	go func(dataCh chan<- int) {
		for i := 0; i < 10; i++ {
			dataCh <- i
		}
		close(dataCh)
	}(pushCh)

	go func(dataCh <-chan int) {
		for n := range dataCh {
			println(n)
		}
	}(pullCh)

	for i := 0; i < 10; i++ {
		pullCh <- <-pushCh
	}
}
