package broadcast

import "testing"

func TestBroadcast(t *testing.T) {
	receivers := make([]chan string, 3)

	for i := range receivers {
		receivers[i] = make(chan string)
	}

	for i := range receivers {
		go func() {
			message := <-receivers[i]
			println("Receiver %d received message: %s", i, message)
		}()
	}

	msg := "Hello, World!"
	Broadcast(msg, receivers)
}
