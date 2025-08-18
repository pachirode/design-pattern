package interpreter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpreter(t *testing.T) {
	assert := assert.New(t)
	expression := Add{
		left: &Number{
			value: 1,
		},
		right: &Number{
			value: 2,
		},
	}
	assert.Equal(3, expression.Interpreter())
}
