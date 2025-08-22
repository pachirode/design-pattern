package semaphore

import (
	"errors"
	"time"
)

var Error = errors.New("信号量已经超时")

type Semaphore struct {
	sem     chan struct{}
	timeout time.Duration
}

func NewSemaphore(n int, timeout time.Duration) *Semaphore {
	return &Semaphore{
		sem:     make(chan struct{}, n),
		timeout: timeout,
	}
}

func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return Error
	}
}

func (s *Semaphore) Release() {
	<-s.sem
}

func (s *Semaphore) IsEmpty() bool {
	return len(s.sem) == 0
}
