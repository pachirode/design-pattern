package object_pool

type ObjectPool struct {
	objects chan *Object
}

type Object struct {
	id int
}

func NewObjectPool(size int) *ObjectPool {
	pool := &ObjectPool{
		objects: make(chan *Object, size),
	}
	for i := 0; i < size; i++ {
		pool.objects <- &Object{
			id: i,
		}
	}
	return pool
}

func (pool *ObjectPool) Acquire() *Object {
	return <-pool.objects
}

func (pool *ObjectPool) Release(object *Object) {
	pool.objects <- object
}
