package datatypes

import (
	"testing"
)

func TestCapuint(t *testing.T) {
	g := capuint(300)
	if g != 255 {
		t.Error("Error, 300 is supposed to become 255")
	}
	l := capuint(-2)
	if l != 0 {
		t.Error("Error, -2 is supposed to become 0")
	}
}

func TestIsEqual(t *testing.T) {
	c1 := Col{12, 12, 12}
	c2 := Col{12, 12, 12}
	if !c1.Equal(c2) {
		t.Errorf("Error, should be equal.")
	}
}
