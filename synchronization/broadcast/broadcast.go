package broadcast

func Broadcast(message string, receivers []chan string) {
	for _, receiver := range receivers {
		receiver <- message
	}
}
