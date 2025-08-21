package decorator

type IShow interface {
	Show()
}

type Girl struct {
}

func NewGirl() *Girl {
	return &Girl{}
}

func (g *Girl) Show() {
	println("girl")
}

type Decorator struct {
	showFn IShow
}

func NewDecorator(showFn IShow) *Decorator {
	return &Decorator{
		showFn: showFn,
	}
}

func (d *Decorator) Show() {
	d.showFn.Show()
}

type Lipstick struct {
	Decorator
}

func NewLipstick(showFn IShow) *Lipstick {
	return &Lipstick{
		Decorator: *NewDecorator(showFn),
	}
}

func (l *Lipstick) Show() {
	println("lipstick")
	l.Decorator.Show()
	println("lipstick end")
}
