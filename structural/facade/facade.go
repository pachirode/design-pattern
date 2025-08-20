package facade

type FreshFood struct {
}

func (f *FreshFood) purchase() {
	println("purchase fresh food")
}

type DailyNecessary struct {
}

func (d *DailyNecessary) purchase() {
	println("purchase daily necessary")
}

type Shop struct {
	freshFood      FreshFood
	dailyNecessary DailyNecessary
}

func NewShop() *Shop {
	return &Shop{
		freshFood:      FreshFood{},
		dailyNecessary: DailyNecessary{},
	}
}

func (s *Shop) Shopping() {
	s.freshFood.purchase()
	s.dailyNecessary.purchase()
}
