package fanout

import (
	"sync"
	"testing"
)

func TestFanOut(t *testing.T) {
	dataStreams := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputChan := gen(dataStreams...)

	ch := square(inputChan)

	outArray := SpiltToChanList(ch, 3)

	var wg sync.WaitGroup
	for i := 0; i < len(outArray); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 读取并处理每个通道中的值，直到通道关闭
			for v := range outArray[i] {
				t.Logf("worker %d got %d\n", i, v)
			}
		}(i)
	}

	wg.Wait()
}
