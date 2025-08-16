package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrategy(t *testing.T) {
	assert := assert.New(t)
	operator := NewOperator()
	operator.SetStrategy(&Add{})
	assert.Equal(5, operator.Calculate(2, 3), "2 + 3 should be 5")
}
