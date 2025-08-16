package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {
	assert := assert.New(t)
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	interfaceSlice := make([]interface{}, len(strSlice))
	for i, v := range strSlice {
		interfaceSlice[i] = v
	}
	aggregate := NewConcreteAggregate(interfaceSlice)
	iterator := aggregate.CreateIterator()
	for _, v := range interfaceSlice {
		assert.Equal(v, iterator.Next())
	}
}
