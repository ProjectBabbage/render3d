package render

import (
	"testing"
)

func TestConvertIndexToScreenIndex(t *testing.T) {
	PixelsX := 200
	PixelsY := 100

	var i, j = 10, 20
	var expected_I, expected_J = 110, 30
	var I, J int
	I, J = convertIndexToScreenIndex(i, j, PixelsX, PixelsY)
	t.Log(i, j)
	t.Log(I, J)

	if I != expected_I || J != expected_J {
		t.Errorf("Indexes conversion to Screen Indexes error.")
	}

}
