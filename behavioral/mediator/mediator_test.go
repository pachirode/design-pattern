package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := NewConcreteMediator()

	user1 := NewUser("user1")
	user2 := NewUser("user2")
	user3 := NewUser("user3")
	mediator.AddColleague(user1)
	mediator.AddColleague(user2)
	mediator.AddColleague(user3)
	user1.SetMediator(mediator)
	user2.SetMediator(mediator)
	user3.SetMediator(mediator)

	user1.SendMessage("hello")
	user2.SendMessage("world")
	user3.SendMessage("!")
}
