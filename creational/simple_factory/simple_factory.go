package simple_factory

type Position string

const (
	PositionStudent Position = "student"
	PositionTeacher Position = "teacher"
)

type Person interface {
	Say() string
}

type Student struct {
}

type Teacher struct {
}

func (s *Student) Say() string {
	return "I'm a student"
}

func (t *Teacher) Say() string {
	return "I'm a teacher"
}

func NewPerson(position Position) Person {
	switch position {
	case PositionStudent:
		return &Student{}
	case PositionTeacher:
		return &Teacher{}
	default:
		return nil
	}
}
