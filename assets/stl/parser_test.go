package stl

import (
	. "broengine/datatypes"
	"testing"
)

// This test tries to parse "cube.stl"
func TestParser(t *testing.T) {
	// we define the channels r,g,b are equally impactful on the factors
	isoChannelsKa := Col{100, 100, 100}
	isoChannelsKd := Col{100, 100, 100}
	isoChannelsKs := Col{100, 100, 100}

	triangles := Parse("files/cube.stl", 1, isoChannelsKa, isoChannelsKd, isoChannelsKs)
	var surfaces = triangles.Surfaces
	if surfaces == nil {
		t.Error("couln't read the file with the third party library")
	}
	if len(surfaces) != 12 {
		t.Errorf("cube.stl is supposed to have 12 triangle, but %d were found", len(surfaces))
	}
}
