package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilderCar(t *testing.T) {
	assert := assert.New(t)
	builder := NewCarStudio()
	builder.Brand("1").Speed(100).Engine("bmw").Wheel(1)
	car := builder.Build()

	assert.Equal(car.Speed(), 100)
	assert.Equal(car.Brand(), "1")
}
