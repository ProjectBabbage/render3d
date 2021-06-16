package assets

import (
	. "broengine/datatypes"
	"fmt"
)

var L0 = Light{Vector{0, -50, -50}, 30, 250, 3}
var L1 = Light{Vector{0, -50, -50}, 30, 200, 0}
var L2 = Light{Vector{0, 0, -50}, 15, 100, 3}
var L3 = Light{Vector{0, 0, 10}, 25, 200, 0}

var L4 = Light{Vector{-2, -2, 5}, 30, 100, 150}

var Path = "assets/"

func SSphere() Scene {
	sphere := ParseStl(Path+"sphere.stl", 1, 1, 1, 1)

	sphere.Translate(Vector{0, 0, 12})

	scene := NewEmptyScene()
	scene.AddObjects(sphere)
	scene.AddLights(L1)

	return scene
}

func SSpherePlane() Scene {
	sphere := ParseStl(Path+"sphere_high_definition.stl", 1, 1, 1, 1)
	sphere2 := ParseStl(Path+"sphere.stl", 1, 1, 1, 1)
	plane := ParseStl(Path+"plane.stl", 1, 1, 5, 1)

	sphere.Translate(Vector{1, 0, 8})
	sphere2.Translate(Vector{-0.5, -1, 9})
	plane.Translate(Vector{0, 1, 7})

	scene := NewEmptyScene()
	scene.AddObjects(sphere, sphere2, plane)
	scene.AddLights(L4)

	return scene
}

func SSpherePlaneShadow() Scene {
	sphere := ParseStl(Path+"sphere.stl", 1, 1, 1, 1)
	plane := ParseStl(Path+"plane.stl", 1, 1, 1, 1)

	sphere.Translate(Vector{0, 0, 12})
	plane.Translate(Vector{0, 1, 7})

	scene := NewEmptyScene()
	scene.AddObjects(sphere, plane)
	scene.AddLights(L3)

	return scene
}

func SCubeRotated() Scene {
	cube_rotated := ParseStl(Path+"cube_rotated.stl", 1, 1, 1, 1)

	cube_rotated.Translate(Vector{2, 0, 15})
	scene := NewEmptyScene()
	scene.AddObjects(cube_rotated)
	scene.AddLights(L1)

	return scene
}

func SCubeManuallyRotated() Scene {
	cube := ParseStl(Path+"cube.stl", 1, 1, 1, 1)
	cube.Rotate(XAxis, 20)
	cube.Rotate(YAxis, 20)
	cube.Translate(Vector{2, 0, 15})
	scene := NewEmptyScene()
	scene.AddObjects(cube)
	scene.AddLights(L1)

	return scene
}

func SFaces(listIndex ...string) Scene {
	var objects = []Object{}

	for _, face := range listIndex {
		filename := fmt.Sprintf(Path+"faces/%s.stl", face)
		o := ParseStl(filename, 1, 1, 1, 1)
		objects = append(objects, o)
	}
	scene := NewEmptyScene()
	scene.AddObjects(objects...)
	scene.AddLights(L1)
	scene.TranslateObjects(Vector{0, 0, 20})

	return scene
}

func SSimpleTriangle() Scene {
	var distance float64 = 100
	p1 := Vector{0, 0, distance}
	p2 := Vector{0, 25, distance}
	p3 := Vector{25, 0, distance}
	newTriangle := NewTriangle(p1, p2, p3, Vector{0, 0, 0}, 1, 1, 1, 1)
	newTriangle.RecomputeNormal()
	o := Object{[]Surface{&newTriangle}}

	scene := NewEmptyScene()
	scene.AddObjects(o)
	scene.AddLights(L2)
	scene.TranslateObjects(Vector{4, -4, 40})

	return scene
}
