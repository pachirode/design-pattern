package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleLogger_Log(t *testing.T) {
	assert := assert.New(t)
	factory := FileLoggerFactory{}
	file_logger := factory.CreateLogger()
	assert.Equal("FileLogger: This is a test log.", file_logger.Log("This is a test log."))
}

func TestFileLogger_Log(t *testing.T) {
	assert := assert.New(t)
	factory := ConsoleLoggerFactory{}
	console_logger := factory.CreateLogger()
	assert.Equal("ConsoleLogger: This is a test log.", console_logger.Log("This is a test log."))
}
