package monitor

import (
	"sync"
)

type Monitor struct {
	mutex *sync.Mutex
	cond  *sync.Cond
	value int
}

func NewMonitor() *Monitor {
	m := &Monitor{
		mutex: new(sync.Mutex),
	}
	m.cond = sync.NewCond(m.mutex)
	return m
}

func (m *Monitor) SetValue(value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.value = value
	m.cond.Signal()
}

func (m *Monitor) GetValue() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for m.value == 0 {
		m.cond.Wait()
	}

	return m.value
}
