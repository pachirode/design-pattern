package interpreter

type Expression interface {
	Interpreter() int
}

type Number struct {
	value int
}

func (n *Number) Interpreter() int {
	return n.value
}

type Add struct {
	left, right Expression
}

func (a *Add) Interpreter() int {
	return a.left.Interpreter() + a.right.Interpreter()
}
