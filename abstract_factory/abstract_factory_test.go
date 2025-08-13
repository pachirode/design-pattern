package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {

	assert := assert.New(t)
	factory := WoodenCarFactory{}
	frame := factory.CreateFrame()
	wheel := factory.CreateWheel()

	assert.Equal("Wooden Frame Lock", frame.Lock())
	assert.Equal("Wooden Frame Unlock", frame.Unlock())
	assert.Equal("Wooden Wheel Run", wheel.Run())
	assert.Equal("Wooden Wheel Stop", wheel.Stop())
}
