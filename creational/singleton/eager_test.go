package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingletonEager(t *testing.T) {
	assert := assert.New(t)
	initEager(1)
	eager := getEager()
	assert.Equal(1, eager.count)
}
