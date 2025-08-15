package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	subject := NewConcreteSubject()
	observer1 := NewConcreteObserver("Observer1")
	observer2 := NewConcreteObserver("Observer2")
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Notify()
}
