package singleton

var eager *Eager

type Eager struct {
	count int
}

func initEager(count int) {
	eager = &Eager{count: count}
}

func getEager() *Eager {
	return eager
}
