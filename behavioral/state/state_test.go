package state

import "testing"

func TestState(t *testing.T) {
	runner := NewRunner()
	runner.Walk()
	runner.Run()
	runner.Walk()
	runner.Stop()
}
