package render

import (
	. "broengine/datatypes"
	"testing"
)

func TestConvertIndexToScreenIndex(t *testing.T) {
	PixelsX := 200
	PixelsY := 100

	var i, j = 10, 20
	var expected_I, expected_J = 110, 70
	var I, J int
	I, J = convertIndexToScreenIndex(i, j, PixelsX, PixelsY)
	t.Log(i, j)
	t.Log(I, J)

	if I != expected_I || J != expected_J {
		t.Errorf("Indexes conversion to Screen Indexes error.")
	}

}

func TestNewScreen(t *testing.T) {
	ns := NewScreen(100, 200)
	if ns.PixelsX != 100 || ns.PixelsY != 200 {
		t.Error("Error, new screen size is wrong")
	}
	if len(ns.Pixels) != 101 || len(ns.Pixels[0]) != 201 {
		t.Error("Error, ns.Pixels length is wrong ")
	}
	if ns.Pixels[0][0] != (Col{0, 0, 0}) {
		t.Error("Error, the initial color of the screen is not black.")
	}
}

func TestMeanScreenSize(t *testing.T) {
	ns := NewScreen(100, 200)
	ms := ns.MeanScreen(2)

	if ms.PixelsX != 50 || ms.PixelsY != 100 {
		t.Error("Error, mean screen size is wrong")
	}
	// real screen is one more pixels bigger
	// I think we should change that
	if len(ms.Pixels) != 51 || len(ms.Pixels[0]) != 101 {
		t.Error("Error, ms.Pixels double array length is wrong (", len(ms.Pixels), len(ms.Pixels[0]), ")")
		t.Error("Expected (101 201)")
	}
}

// intensity is a Col
func TestMeanScreenIntensity(t *testing.T) {
	ns := NewScreen(2, 2)
	ns.Pixels = [][]Col{
		{Col{1, 1, 1}, Col{2, 2, 2}},
		{Col{1, 1, 1}, Col{1, 1, 1}},
	}

	ms := ns.MeanScreen(2)
	expected_col1 := Col{1.25, 1.25, 1.25}

	if ms.Pixels[0][0] != expected_col1 {
		t.Error("Color mean should be ", expected_col1, "instead it was:", ms.Pixels[0][0])
	}

	ns2 := NewScreen(2, 2)

	ns2.Pixels = [][]Col{
		{Col{1, 1, 1}, Col{3, 3, 3}},
		{Col{1, 1, 1}, Col{3, 3, 3}},
	}

	ms2 := ns2.MeanScreen(2)
	expected_col2 := Col{2, 2, 2}

	if ms2.Pixels[0][0] != expected_col2 {
		t.Error("Color mean should be", expected_col2, "instead it was:", ms2.Pixels[0][0])
	}
}
