package fanout

import (
	"sync"
)

func SpiltToChanList(ch <-chan int, n int) []chan int {
	chanList := []chan int{}
	for i := 0; i < n; i++ {
		chanList = append(chanList, make(chan int))
	}

	go func() {
		defer func() {
			for _, c := range chanList {
				close(c)
			}
		}()

		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				for _, c := range chanList {
					c <- v
				}
			}
		}
	}()

	return chanList
}

func SplitToWorkers(ch <-chan int, n int) []chan int {
	chanList := []chan int{}
	for i := 0; i < n; i++ {
		chanList = append(chanList, make(chan int))
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			val, ok := <-ch
			if !ok {
				return
			}

			for _, c := range chanList {
				c <- val
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		for _, c := range chanList {
			close(c)
		}
	}()

	return chanList
}

func SplitChanListRandom(ch <-chan int, n int) []chan int {
	chanList := []chan int{}
	for i := 0; i < n; i++ {
		chanList = append(chanList, make(chan int))
	}

	go func() {
		defer func() {
			for _, c := range chanList {
				close(c)
			}
		}()

		for {
			for _, c := range chanList {
				select {
				case v, ok := <-ch:
					if !ok {
						return
					}
					c <- v
				}
			}

		}
	}()

	return chanList
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}
