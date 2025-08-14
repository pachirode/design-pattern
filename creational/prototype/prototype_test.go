package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrototype(t *testing.T) {
	assert := assert.New(t)

	circle := &Circle{Type: "Circle"}
	cloneCircle := circle.Clone()

	assert.Equal(circle.GetType(), cloneCircle.GetType(), "The type of the cloned shape should be the same as the original.")
}
