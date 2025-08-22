package readwritelock

import "sync"

type Data struct {
	rwLock *sync.RWMutex
	data   int
}

func (d *Data) Read() int {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()

	return d.data
}

func (d *Data) Write(value int) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()

	d.data = value
}
