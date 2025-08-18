package visitor

import "fmt"

type Visitor interface {
	VisitorMath(m *MathLesson)
	VisitorEnglish(e *EnglishLesson)
}

type Lesson interface {
	Accept(visitor Visitor)
	Price() int
}

type School struct {
	Lessons []Lesson
}

func NewSchool() *School {
	return &School{
		Lessons: []Lesson{},
	}
}

func (s *School) AddLesson(lesson Lesson) {
	s.Lessons = append(s.Lessons, lesson)
}

func (s *School) Accept(visitor Visitor) {
	for _, lesson := range s.Lessons {
		lesson.Accept(visitor)
	}
}

type MathLesson struct {
}

func NewMathLesson() *MathLesson {
	return &MathLesson{}
}

func (m *MathLesson) Accept(visitor Visitor) {
	visitor.VisitorMath(m)
}

func (m *MathLesson) Price() int {
	return 100
}

type EnglishLesson struct {
}

func NewEnglishLesson() *EnglishLesson {
	return &EnglishLesson{}
}
func (e *EnglishLesson) Accept(visitor Visitor) {
	visitor.VisitorEnglish(e)
}

func (e *EnglishLesson) Price() int {
	return 50
}

type MemberVisitor struct {
}

func NewMemberVisitor() *MemberVisitor {
	return &MemberVisitor{}
}

func (m *MemberVisitor) VisitorMath(m1 *MathLesson) {
	fmt.Printf("会员学习打八折：%d\n", m1.Price()*8/10)
}

func (m *MemberVisitor) VisitorEnglish(e1 *EnglishLesson) {
	fmt.Printf("会员学习打五折：%d\n", e1.Price()*5/10)
}
