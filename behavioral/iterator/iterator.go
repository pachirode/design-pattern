package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Aggregate interface {
	CreateIterator() Iterator
}

type ConcreteIterator struct {
	items []interface{}
	index int
}

type ConcreteAggregate struct {
	items []interface{}
}

func NewConcreteIterator(items []interface{}) *ConcreteIterator {
	return &ConcreteIterator{items: items, index: 0}
}

func (ci *ConcreteIterator) HasNext() bool {
	return ci.index < len(ci.items)
}

func (ci *ConcreteIterator) Next() interface{} {
	item := ci.items[ci.index]
	ci.index++
	return item
}

func NewConcreteAggregate(items []interface{}) *ConcreteAggregate {
	return &ConcreteAggregate{items: items}
}

func (ca *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(ca.items)
}
