package monitor

import "testing"

func TestMonitor(t *testing.T) {

	monitor := NewMonitor()

	go func() {
		monitor.SetValue(1)
	}()

	data := monitor.GetValue()
	println("已获取值为: ", data)
}
