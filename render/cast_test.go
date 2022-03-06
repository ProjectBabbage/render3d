// https://en.wikipedia.org/wiki/Phong_reflection_model
package render

import (
	. "render3d/datatypes"
	"fmt"
	"testing"
)

var ray = NewRay(Vector{}, Vector{0, 0, 1})
var mat = Material{
	A:  1,
	Ka: IsoCol(50),
	Kd: IsoCol(100),
	Ks: IsoCol(50),
	Kr: IsoCol(20),
}
var triangle = NewTriangle(
	Vector{0, 1, 1},
	Vector{1, -1, 1},
	Vector{-1, -1, 1},
	Vector{0, 0, -1},
	mat,
)
var light = Light{
	Vector: Vector{-1, -1, 0},
	Ia:     IsoCol(50),
	Id:     IsoCol(50),
	Is:     IsoCol(50),
}

func TestCast(t *testing.T) {
	scene := NewEmptyScene()

	scene.AddObjects(Object{[]Surface{&triangle}})
	scene.AddLights(light)

	c := Cast(ray, scene, 2)
	fmt.Println(c)
	if c == (Col{}) {
		t.Error("Cast should return a non zero Col.")
	}
}

func TestComputeIntensity(t *testing.T) {

	iR := IntersectRes{
		Vector:           Vector{0, 0, 1},
		HasIntersection:  true,
		DistanceToOrigin: 1,
		Normale:          Vector{0, 0, -1},
		Material:         mat,
	}
	scene := NewEmptyScene()
	scene.AddObjects(Object{[]Surface{&triangle}})
	scene.AddLights(light)

	col := compute_intensity(iR, ray, scene)
	if col.R < 10 || col.G < 10 || col.B < 10 {
		t.Errorf("Color col (%v) is incorrect", col)
	}
}
