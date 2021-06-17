package assets

import (
	. "broengine/datatypes"
	"fmt"
)

var L0 = Light{Vector{0, -50, -50}, 30, 250, 3}
var L1 = Light{Vector{0, -50, -50}, 30, 200, 0}
var L2 = Light{Vector{0, 0, -50}, 15, 100, 3}
var L3 = Light{Vector{0, 0, 10}, 25, 200, 0}

var Path = "assets/"

func SCubePlane() Scene {
	sphere := ParseStl(Path+"sphere.stl", 1, 1, 1, 1)
	plane := ParseStl(Path+"plane.stl", 1, 1, 1, 1)

	sphere.Translate(Vector{0, 0, 12})
	plane.Translate(Vector{0, 1, 5})

	scene := NewEmptyScene()
	scene.AddObjects(sphere, plane)
	scene.AddLights(L1)

	return scene
}

func SCubePlaneShadow() Scene {
	sphere := ParseStl(Path+"sphere.stl", 1, 1, 1, 1)
	plane := ParseStl(Path+"plane.stl", 1, 1, 1, 1)

	sphere.Translate(Vector{0, 0, 12})
	plane.Translate(Vector{0, 1, 5})

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
	// 	listIndex := []string{
	// 	"top",
	// 	"front",
	// 	"back",
	// 	"right",
	// 	"left",
	// 	"bottom",
	// }
	var objects = []Object{}

	for _, face := range listIndex {
		filename := fmt.Sprintf(Path+"faces/%s.stl", face)
		o := ParseStl(filename, 1, 1, 1, 1)
		objects = append(objects, o)
	}
	scene := NewEmptyScene()
	scene.AddObjects(objects...)
	scene.AddLights(L1)
	scene.TranslateObjects(Vector{0, 0, 10})

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

	// scene = scene.Translate(Vector{4, -4, 40})
	// var screen = new(render.Screen)
	// screen.Init() // set to black every pixel

	// ray := NewRay(Vector{2, 2, 0}, Vector{0, 0, 1})
	// intensity := render.Cast(ray, scene)
	// c := color.Gray16{uint16(intensity)}
	// fmt.Println(intensity)
	// screen.FillPixel(0, 0, c)

	return scene
}
