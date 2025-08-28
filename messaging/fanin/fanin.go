package fanin

import "sync"

// Merge 将不同通道组合成一个
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int, 3)

	wg.Add(len(cs))

	// 从输入通道中赋值值到输出通道
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}

		wg.Done()
	}

	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateNumbersPipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

	return out
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
