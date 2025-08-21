package proxy

type IBuy interface {
	Buy()
}

type User struct {
}

type BuyerProxy struct {
	User *User
}

func NewUser() *User {
	return &User{}
}

func (u *User) Buy() {
	println("User buy")
}

func NewBuyerProxy(user *User) *BuyerProxy {
	return &BuyerProxy{
		User: user,
	}
}

func (p *BuyerProxy) Buy() {
	println("BuyerProxy buy")
	p.User.Buy()
	println("BuyerProxy buy end")
}
