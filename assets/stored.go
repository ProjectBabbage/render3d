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
var L5 = Light{Vector{0, 0, 0}, 30, 150, 0}
var L6 = Light{Vector{-8, 0, 0}, 30, 150, 70}
var L7 = Light{Vector{0, -50, -50}, 30, 150, 70}
var L8 = Light{Vector{-5, -5, -10}, 30, 150, 70}

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

func STwoTrianglesPlane() Scene {
	triangles := ParseStl(Path+"two_triangles.stl", 1, 1, 1, 1)
	plane := ParseStl(Path+"plane.stl", 1, 1, 1, 1)
	// plane.Rotate(XAxis, 90)
	triangles.Rotate(YAxis, -45)

	triangles.Translate(Vector{0, 0, 4})
	plane.Translate(Vector{0, 2, 10})

	scene := NewEmptyScene()
	scene.AddObjects(triangles, plane)
	scene.AddLights(L5)
	scene.Print()

	return scene
}

func STwoTrianglesPlane2() Scene {

	q1 := Vector{0.22975452523155737, 0, 3.026750959038924}
	q2 := Vector{-0.3162252123460546, 0, 3.051315849298438}
	x := -0.11487165708492686
	y := 0.49999698996543884
	z := 3.141625397815853
	triangle1 := NewTriangle(q1, q2, Vector{x, y, z}, Vector{}, 1, 1, 1, 1)
	triangle1.RecomputeNormal()
	triangle2 := NewTriangle(q1, Vector{x, -y, z}, q2, Vector{}, 1, 1, 1, 1)
	triangle2.RecomputeNormal()
	triangles := Object{[]Surface{&triangle1, &triangle2}}
	plane := ParseStl(Path+"plane.stl", 1, 1, 1, 1)
	// plane.Rotate(XAxis, 90)
	// triangles.Rotate(YAxis, -45)

	triangles.Translate(Vector{0, 0, 0})
	plane.Translate(Vector{0, 2, 10})

	scene := NewEmptyScene()
	scene.AddObjects(triangles, plane)
	scene.AddLights(L5)
	scene.Print()

	return scene
}

func STrueSphere() Scene {
	s := NewSphere(Vector{}, 1, 1, 1, 1, 30)
	sphere := Object{[]Surface{&s}}

	sphere.Translate(Vector{0, 0, 10})

	scene := NewEmptyScene()
	scene.AddObjects(sphere)
	scene.AddLights(L7)

	return scene
}

func STrueSpherePlane() Scene {
	s := NewSphere(Vector{}, 1, 1, 1, 1, 30)
	sphere := Object{[]Surface{&s}}
	plane := ParseStl(Path+"plane.stl", 1, 1, 1, 1)
	plane.Rotate(XAxis, 90)

	sphere.Translate(Vector{-1, 0, 10})
	plane.Translate(Vector{3, 0, 20})

	scene := NewEmptyScene()
	scene.AddObjects(sphere, plane)
	scene.AddLights(L6)

	return scene
}

func STrueSphereInside() Scene {
	s1 := NewSphere(Vector{}, 25, 1, 0.2, 1, 30)
	s2 := NewSphere(Vector{}, 1, 1, 1, 1, 30)

	sphere1 := Object{[]Surface{&s1}}
	sphere2 := Object{[]Surface{&s2}}
	sphere2.Translate(Vector{0, 0, 10})

	scene := NewEmptyScene()
	scene.AddObjects(sphere1, sphere2)
	scene.AddLights(L8)

	return scene
}
