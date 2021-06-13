package util

import (
	"testing"
)

// This test tries to parse "../assets/cube.stl"
func TestStlParser(t *testing.T) {
	filepath := "../assets/cube.stl"

	triangles := ParseStl(filepath, 1, 1, 1, 1)
	if triangles == nil {
		t.Error("couln't read the file with the third party library")
	}
	if len(triangles) != 12 {
		t.Errorf("cube.stl is supposed to have 12 triangle, but %d were found", len(triangles))
	}
}
