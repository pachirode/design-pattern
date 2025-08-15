package observer

import "fmt"

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type Observer interface {
	Update(subject Subject)
}

type ConcreteSubject struct {
	observers []Observer
}

type ConcreteObserver struct {
	name string
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (cs *ConcreteSubject) Attach(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) Detach(observer Observer) {
	for i, o := range cs.observers {
		if o == observer {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			break
		}
	}
}

func (cs *ConcreteSubject) Notify() {
	fmt.Println("开始发送通知...")
	for _, observer := range cs.observers {
		observer.Update(cs)
	}
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (co *ConcreteObserver) Update(subject Subject) {
	fmt.Println("收到通知：", subject)
}
