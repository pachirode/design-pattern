package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {

	scrambled := NewScrambled()
	DoMake(scrambled)
}
