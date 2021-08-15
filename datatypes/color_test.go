package datatypes

import (
	"testing"
)

func TestCapfloat64(t *testing.T) {
	g := capfloat64(300)
	if g != 255 {
		t.Error("Error, 300 is supposed to become 255")
	}
	l := capfloat64(-2)
	if l != 0 {
		t.Error("Error, -2 is supposed to become 0")
	}
}

func TestIsEqual(t *testing.T) {
	c1 := Col{12, 12, 12}
	c2 := Col{12, 12, 12}
	if c1 != c2 {
		t.Errorf("Error, should be equal.")
	}
}

func TestDilateColorByChannels(t *testing.T) {
	c := IsoCol(10).DilateColorByChannels(Col{0.1, 0.5, 1})
	if c.R != 1 || c.G != 5 || c.B != 10 {
		t.Errorf("Error, should be equal.")
	}
}
