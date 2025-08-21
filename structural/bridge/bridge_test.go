package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	circleBulb := NewCircleSwitch(&Bulb{})
	circleBulb.TurnOn()
	circleBulb.TurnOff()

	circleFan := NewCircleSwitch(&Fan{})
	circleFan.TurnOn()
	circleFan.TurnOff()

	squareFan := NewSquareSwitch(&Fan{})
	squareFan.TurnOn()
	squareFan.TurnOff()

	squareBulb := NewSquareSwitch(&Bulb{})
	squareBulb.TurnOn()
	squareBulb.TurnOff()
}
