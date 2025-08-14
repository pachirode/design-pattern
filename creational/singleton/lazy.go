package singleton

import "sync"

var (
	lazy *Lazy
	once sync.Once
)

type Lazy struct {
}

func GetLazy() *Lazy {
	once.Do(func() {
		lazy = &Lazy{}
	})

	return lazy
}
