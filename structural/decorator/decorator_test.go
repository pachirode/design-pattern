package decorator

import "testing"

func TestDecorator(t *testing.T) {
	girl := NewGirl()
	lipstick := NewLipstick(girl)
	lipstick.Show()
}
