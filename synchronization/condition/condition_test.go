package condition

import "testing"

func TestCondition(t *testing.T) {
	data := NewData()

	go func() {
		println("等待开始标志")
		data.Wait()
	}()

	println("设置开始标志")
	data.SignalReady()
}
