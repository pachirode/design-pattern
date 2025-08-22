package semaphore

import (
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(2, time.Second*1)

	for i := 0; i < 5; i++ {
		go func() {
			err := s.Acquire()
			if err != nil {
				println(i, "信号量已经超时")
			} else {
				println(i, "信号量已经获取成功")
			}
			defer s.Release()
			time.Sleep(time.Second * 2)
		}()
	}

	time.Sleep(time.Second * 5)
}
