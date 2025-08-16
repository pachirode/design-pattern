package strategy

type Strategy interface {
	Do(interface{}, interface{}) interface{}
}

type Add struct {
}

func (a *Add) Do(x, y interface{}) interface{} {
	return x.(int) + y.(int)
}

type Sub struct {
}

func (s *Sub) Do(x, y interface{}) interface{} {
	return x.(int) - y.(int)
}

type Operator struct {
	strategy Strategy
}

func NewOperator() *Operator {
	return &Operator{}
}

func (o *Operator) SetStrategy(strategy Strategy) {
	o.strategy = strategy
}

func (o *Operator) Calculate(x, y interface{}) interface{} {
	return o.strategy.Do(x, y)
}
