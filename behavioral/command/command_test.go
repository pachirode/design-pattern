package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	assert := assert.New(t)
	runner := Runner{}

	runCommand := RunCommand{
		runner: &runner,
	}
	stopCommand := StopCommand{
		runner: &runner,
	}
	invoker := Invoker{
		command: &runCommand,
	}
	assert.Equal("run", invoker.Execute())
	invoker = Invoker{
		command: &stopCommand,
	}
	assert.Equal("stop", invoker.Execute())
}
