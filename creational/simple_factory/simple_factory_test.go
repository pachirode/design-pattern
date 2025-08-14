package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudent_Say(t *testing.T) {
	assert := assert.New(t)
	student := NewPerson(PositionStudent)
	words := student.Say()

	assert.Equal("I'm a student", words)
}

func TestTeacher_Say(t *testing.T) {
	assert := assert.New(t)
	teacher := NewPerson(PositionTeacher)
	words := teacher.Say()

	assert.Equal("I'm a teacher", words)
}
