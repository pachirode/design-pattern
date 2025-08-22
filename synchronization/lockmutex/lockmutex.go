package lockmutex

import "sync"

type Counter struct {
	mutex *sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{
		mutex: new(sync.Mutex),
		count: 0,
	}
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
	println("当前计数为：", c.count)
}
