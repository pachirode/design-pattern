package object_pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectPool(t *testing.T) {
	assert := assert.New(t)
	pool := NewObjectPool(5)
	for i := 0; i < 5; i++ {
		object := pool.Acquire()
		assert.Equal(object.id, i)
	}
}
