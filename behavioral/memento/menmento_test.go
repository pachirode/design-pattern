package memento

import "testing"

func TestMemento(t *testing.T) {
	editor := NewEditor(NewDoc("hello.txt", ""))
	editor.Append("hello")
	editor.Append("world")
	editor.Append("!")
	editor.Undo()
}
