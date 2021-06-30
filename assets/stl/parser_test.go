package stl

import (
	"testing"
)

// This test tries to parse "cube.stl"
func TestParser(t *testing.T) {
	triangles := Parse("files/cube.stl", 1, 1, 1, 1)
	var surfaces = triangles.Surfaces
	if surfaces == nil {
		t.Error("couln't read the file with the third party library")
	}
	if len(surfaces) != 12 {
		t.Errorf("cube.stl is supposed to have 12 triangle, but %d were found", len(surfaces))
	}
}
