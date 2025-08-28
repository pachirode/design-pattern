package pub_sub

import "testing"

func TestPubSub(t *testing.T) {
	p := NewPublisher()
	ch1 := make(chan string)
	ch2 := make(chan string)
	p.Subscribe("1", ch1)
	p.Subscribe("2", ch2)

	go func() {
		for {
			select {
			case msg := <-ch1:
				t.Logf("1: %s", msg)
			case msg := <-ch2:
				t.Logf("2: %s", msg)
			}
		}
	}()

	p.Publish("1", "hello")
	p.Publish("2", "world")

}
