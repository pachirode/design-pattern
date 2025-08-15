package mediator

import "fmt"

type Mediator interface {
	SendMessage(message string, colleague Colleague)
}

type ConcreteMediator struct {
	colleagues []Colleague
}

type Colleague interface {
	SendMessage(message string)
	ReceiveMessage(message string)
	SetMediator(mediator Mediator)
}

type User struct {
	name     string
	mediator Mediator
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{colleagues: make([]Colleague, 0)}
}

func (cm *ConcreteMediator) AddColleague(colleague Colleague) {
	cm.colleagues = append(cm.colleagues, colleague)
}

func (cm *ConcreteMediator) SendMessage(message string, colleague Colleague) {
	for _, c := range cm.colleagues {
		if c != colleague {
			c.ReceiveMessage(message)
		}
	}
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) SendMessage(message string) {
	u.mediator.SendMessage(message, u)
}
func (u *User) ReceiveMessage(message string) {
	fmt.Printf("User %s received message: %s\n", u.name, message)
}
func (u *User) SetMediator(mediator Mediator) {
	u.mediator = mediator
}
