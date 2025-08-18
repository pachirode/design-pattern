package registry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry(t *testing.T) {
	assert := assert.New(t)
	reg := Registry{
		registry: map[string]interface{}{
			"math":    10,
			"english": 20,
		},
	}
	assert.Equal(10, reg.Get("math"))
	assert.Equal(20, reg.Get("english"))
}
