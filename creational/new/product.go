package new

type Product struct {
	Name string
}

func (p *Product) GetName() string {
	return p.Name
}

func NewProduct(name string) *Product {
	return &Product{
		Name: name,
	}
}
