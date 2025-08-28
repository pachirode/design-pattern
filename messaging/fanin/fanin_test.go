package fanin

import (
	"fmt"
	"testing"
)

func TestFanIn(t *testing.T) {
	dataStream1 := generateNumbersPipeline([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	dataStream2 := generateNumbersPipeline([]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	squareStream1 := squareNumber(dataStream1)
	squareStream2 := squareNumber(dataStream2)
	sum := Merge(squareStream1, squareStream2)

	for n := range sum {
		fmt.Println(n)
	}
}
