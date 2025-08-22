package condition

import "sync"

type Data struct {
	ready bool
	cond  *sync.Cond
}

func NewData() *Data {
	d := &Data{
		ready: false,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	return d
}

func (d *Data) Wait() {
	d.cond.L.Lock()
	defer d.cond.L.Unlock()
	for !d.ready {
		d.cond.Wait()
	}

	println("数据准备完毕")
}

func (d *Data) SignalReady() {
	d.cond.L.Lock()
	defer d.cond.L.Unlock()
	d.ready = true
	d.cond.Signal()
}
