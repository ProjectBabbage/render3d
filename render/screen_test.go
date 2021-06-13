package render

import (
	"testing"
)

func TestConvertIndexToScreenIndex(t *testing.T) {
	PixelsWidth := 200
	PixelsHeight := 100

	var i, j = 0, 0
	var expected_I, expected_J = -100, 50
	var I, J int
	I, J = convertIndexToScreenIndex(i, j, PixelsWidth, PixelsHeight)
	t.Log(i, j)
	t.Log(I, J)

	if I != expected_I && J != expected_J {
		t.Errorf("Indexes conversion to Screen Indexes error.")
	}

}
