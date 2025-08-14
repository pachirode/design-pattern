package new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductGetName(t *testing.T) {
	assert := assert.New(t)

	product := NewProduct("test")
	assert.Equal("test", product.GetName())
}
