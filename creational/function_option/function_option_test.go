package function_option

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestServer_FunctionOption(t *testing.T) {
	assert := assert.New(t)
	server := NewServer(WithTimeout(10 * time.Second))

	assert.Equal(server.timeout, 10*time.Second)
}
