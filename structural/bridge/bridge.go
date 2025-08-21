package bridge

import "fmt"

type Application interface {
	Run()
	Name() string
}

type Bulb struct {
}

func (b *Bulb) Run() {
	fmt.Println("Bulb is running")
}

func (b *Bulb) Name() string {
	return "Bulb"
}

type Fan struct {
}

func (f *Fan) Run() {
	fmt.Println("Fan is running")
}

func (f *Fan) Name() string {
	return "Fan"
}

type Switch interface {
	TurnOn()
	TurnOff()
}

type CircleSwitch struct {
	application Application
}

func NewCircleSwitch(app Application) *CircleSwitch {
	return &CircleSwitch{
		application: app,
	}
}

func (c *CircleSwitch) TurnOn() {
	fmt.Printf("打开圆形开关，打开 %s, 开始运行", c.application.Name())
	c.application.Run()
}

func (c *CircleSwitch) TurnOff() {
	fmt.Printf("关闭圆形开关，关闭 %s\n", c.application.Name())
}

type SquareSwitch struct {
	application Application
}

func NewSquareSwitch(app Application) *SquareSwitch {
	return &SquareSwitch{
		application: app,
	}
}

func (s *SquareSwitch) TurnOn() {
	fmt.Printf("打开正方形开关，打开 %s, 启动成功\n", s.application.Name())
	s.application.Run()
}

func (s *SquareSwitch) TurnOff() {
	fmt.Printf("关闭正方形开关，关闭 %s\n", s.application.Name())
}
