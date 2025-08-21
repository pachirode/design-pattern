package flyweight

import (
	"errors"
	"fmt"
)

type Color interface {
	Color() string
}

type RedFlower struct {
	color string
}

func NewRedFlower() *RedFlower {
	return &RedFlower{color: "red"}
}

func (r *RedFlower) Color() string {
	return r.color
}

type BlueFlower struct {
	color string
}

func NewBlueFlower() *BlueFlower {
	return &BlueFlower{color: "blue"}
}

func (b *BlueFlower) Color() string {
	return b.color
}

const (
	RED  = "red"
	BLUE = "blue"
)

type ColorFactory struct {
	colors map[string]Color
}

func (c *ColorFactory) GetColor(color string) (Color, error) {
	if c.colors[color] != nil {
		return c.colors[color], nil
	}

	if color == RED {
		c.colors[RED] = NewRedFlower()
	} else if color == BLUE {
		c.colors[BLUE] = NewBlueFlower()
	} else {
		return nil, errors.New("color not found")
	}

	return c.colors[color], nil
}

var colorFactory = &ColorFactory{
	colors: make(map[string]Color),
}

func GetColorFactory() *ColorFactory {
	return colorFactory
}

type Buyer struct {
	color Color
	name  string
}

func NewBuyer(color Color, name string) *Buyer {
	return &Buyer{
		color: color,
		name:  name,
	}
}

type Market struct {
	buyers  []Buyer
	sellers []Buyer
}

func NewMarket() *Market {
	return &Market{
		buyers:  make([]Buyer, 0),
		sellers: make([]Buyer, 0),
	}
}

func (m *Market) AddBuyer(color Color) {
	buyer := Buyer{
		color: color,
		name:  fmt.Sprintf("Buyer-%s", color.Color()),
	}
	m.buyers = append(m.buyers, buyer)
}

func (m *Market) AddSeller(color Color) {
	seller := Buyer{
		color: color,
		name:  fmt.Sprintf("Seller-%s", color.Color()),
	}
	m.sellers = append(m.sellers, seller)
}
