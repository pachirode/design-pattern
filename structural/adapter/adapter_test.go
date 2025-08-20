package adapter

import "testing"

func TestAdapter(t *testing.T) {
	plug := &TwoPin{}
	threePowerSocket := ThreePowerSocket{}

	threePowerSocket.PlugCharging(plug)
	powerAdapter := NewPowerAdapter(&threePowerSocket)
	powerAdapter.AdapterPlugCharging(plug)
}
