package main

import (
	"fmt"
	"time"
)

const batchSize = 5

func processData(data []string) {
	fmt.Printf("Processing batch: %v\n", data)
	time.Sleep(2 * time.Second)
}

func main() {
	data := []string{"Data 1", "Data 2", "Data 3", "Data 4", "Data 5", "Data 6", "Data 7", "Data 8", "Data 9", "Data 10"}
	for i := 0; i < len(data); i += batchSize {
		batch := data[i:min(i+batchSize, len(data))]
		go processData(batch)
	}

	time.Sleep(10 * time.Second)
	println("All batches processed")
}
