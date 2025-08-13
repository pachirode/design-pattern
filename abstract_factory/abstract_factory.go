package abstract_factory

type Car interface {
	CreateFrame() Frame
	CreateWheel() Wheel
}

type Frame interface {
	Lock() string
	Unlock() string
}

type Wheel interface {
	Run() string
	Stop() string
}

type WoodenCarFactory struct {
}

type WoodenFrame struct {
}

type WoodenWheel struct {
}

func (car *WoodenCarFactory) CreateFrame() Frame {
	return &WoodenFrame{}
}

func (car *WoodenCarFactory) CreateWheel() Wheel {
	return &WoodenWheel{}
}

func (frame *WoodenFrame) Lock() string {
	return "Wooden Frame Lock"
}

func (frame *WoodenFrame) Unlock() string {
	return "Wooden Frame Unlock"
}

func (wheel *WoodenWheel) Run() string {
	return "Wooden Wheel Run"
}

func (wheel *WoodenWheel) Stop() string {
	return "Wooden Wheel Stop"
}
