package builder

type ICar interface {
	Speed() int
	Brand() string
	Brief()
}

type ICarBuilder interface {
	Wheel(wheel int) ICarBuilder
	Engine(engine string) ICarBuilder
	Speed(max int) ICarBuilder
	Brand(brand string) ICarBuilder
	Build() ICar
}

type Car struct {
	Wheel     int
	Engine    string
	MaxSpeed  int
	BrandName string
}

func (c *Car) Speed() int {
	return c.MaxSpeed
}

func (c *Car) Brand() string {
	return c.BrandName
}

func (c *Car) Brief() {
	println("wheel:", c.Wheel, "engine:", c.Engine, "max speed:", c.MaxSpeed, "brand:", c.BrandName)
}

type CarStudio struct {
	protoType *Car
}

func NewCarStudio() ICarBuilder {
	return &CarStudio{protoType: &Car{}}
}

func (c *CarStudio) Wheel(wheel int) ICarBuilder {
	c.protoType.Wheel = wheel
	return c
}

func (c *CarStudio) Engine(engine string) ICarBuilder {
	c.protoType.Engine = engine
	return c
}

func (c *CarStudio) Speed(max int) ICarBuilder {
	c.protoType.MaxSpeed = max
	return c
}

func (c *CarStudio) Brand(brand string) ICarBuilder {
	c.protoType.BrandName = brand
	return c
}

func (c *CarStudio) Build() ICar {
	car := &Car{
		Wheel:     c.protoType.Wheel,
		Engine:    c.protoType.Engine,
		MaxSpeed:  c.protoType.MaxSpeed,
		BrandName: c.protoType.BrandName,
	}

	return car
}
