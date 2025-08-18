package visitor

import "testing"

func TestVisitor(t *testing.T) {
	school := NewSchool()
	school.AddLesson(NewMathLesson())
	school.AddLesson(NewEnglishLesson())
	school.Accept(NewMemberVisitor())
}
