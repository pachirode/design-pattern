package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	superMarket := NewMarket()
	colorFactory := GetColorFactory()
	red, _ := colorFactory.GetColor("red")
	blue, _ := colorFactory.GetColor("blue")
	superMarket.AddSeller(red)
	superMarket.AddSeller(red)
	superMarket.AddBuyer(blue)
	superMarket.AddBuyer(blue)

	for color, buyer := range superMarket.buyers {
		fmt.Println(color, buyer.name)
	}
	for color, seller := range superMarket.sellers {
		fmt.Println(color, seller.name)
	}
}
