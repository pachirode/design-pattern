package adapter

import "fmt"

type IPin interface {
	GetPins() int
}

type TwoPin struct {
}

func (p *TwoPin) GetPins() int {
	return 2
}

type IPowerSocket interface {
	AdapterPlugCharging(p IPin)
}

type IThreePowerSocket interface {
	PlugCharging(p IPin)
}

type ThreePowerSocket struct {
}

func (p *ThreePowerSocket) PlugCharging(pin IPin) {
	if pin.GetPins() != 3 {
		fmt.Println("Only plugged for ThreePin")
	} else {
		fmt.Println("charging ...")
	}
}

type PowerAdapter struct {
	IThreePowerSocket
	pin int
}

func NewPowerAdapter(socket IThreePowerSocket) IPowerSocket {
	return &PowerAdapter{
		IThreePowerSocket: socket,
		pin:               0,
	}
}

func (p *PowerAdapter) GetPins() int {
	return p.pin
}

func (p *PowerAdapter) AdapterPlugCharging(pin IPin) {
	if pin.GetPins() == 2 {
		p.pin = 3
		p.PlugCharging(p)
	}
}
