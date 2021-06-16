package main

import (
	. "broengine/assets"
	. "broengine/config"
	. "broengine/datatypes"
	"broengine/render"
	"fmt"
	"image/color"
)

func main() {
	// testCubeRotated()
	// testCube()
	testSphere()
	// testFaces()

	// triangleTest()
	// displayTest()
	// displayTest1()
	// displayTest2()
}

func testSphere() {
	sphere := ParseStl("assets/sphere.stl", 1, 1, 1, 1)
	plane := ParseStl("assets/plane.stl", 1, 1, 1, 1)

	sphere.Translate(Vector{0, 0, 12})
	plane.Translate(Vector{0, 1, 5})

	scene := NewEmptyScene()
	scene.AddObjects(sphere, plane)
	scene.AddLights(Lights...)

	render.Render(scene)
}

func testCubeRotated() {
	cube_rotated := ParseStl("assets/cube_rotated.stl", 1, 1, 1, 1)

	cube_rotated.Translate(Vector{2, 0, 15})
	scene := NewEmptyScene()
	scene.AddObjects(cube_rotated)
	scene.AddLights(Lights...)

	render.Render(scene)
}

func testCube() {
	cube := ParseStl("assets/cube.stl", 1, 1, 1, 1)

	cube.Rotate(XAxis, 20)
	cube.Rotate(YAxis, 20)
	cube.Translate(Vector{2, 0, 15})
	scene := NewEmptyScene()
	scene.AddObjects(cube)
	scene.AddLights(Lights...)

	render.Render(scene)
}

func testFaces() {
	var objects = []Object{}

	listIndex := []string{
		"top",
		"front",
		"back",
		"right",
		"left",
		"bottom",
	}
	for _, face := range listIndex {
		filename := fmt.Sprintf("assets/faces/%s.stl", face)
		o := ParseStl(filename, 1, 1, 1, 1)
		objects = append(objects, o)
	}

	scene := NewEmptyScene()
	scene.AddObjects(objects...)
	scene.AddLights(Lights...)
	scene.TranslateObjects(Vector{0, 0, 10})

	render.Render(scene)
}

func triangleTest() {
	var distance float64 = 100
	p1 := Vector{0, 0, distance}
	p2 := Vector{0, 25, distance}
	p3 := Vector{25, 0, distance}
	newTriangle := NewTriangle(p1, p2, p3, Vector{0, 0, 0}, 1, 1, 1, 1)
	newTriangle.RecomputeNormal()
	o := Object{[]Surface{&newTriangle}}

	scene := NewEmptyScene()
	scene.AddObjects(o)
	scene.AddLights(Lights...)
	// scene = scene.Translate(Vector{4, -4, 40})
	var screen = new(render.Screen)
	screen.Init() // set to black every pixel

	ray := NewRay(Vector{2, 2, 0}, Vector{0, 0, 1})
	intensity := render.Cast(ray, scene)
	c := color.Gray16{uint16(intensity)}
	fmt.Println(intensity)
	screen.FillPixel(0, 0, c)

	render.Rendering(screen)
}

func displayTest() {
	var screen = new(render.Screen)
	screen.Init() // set to black every pixel
	var pixelColor = color.White
	size := 100
	for i := -size / 2; i < size/2; i++ {
		for j := 0; j < i; j++ {
			screen.FillPixel(i, j, pixelColor)
		}
	}
	render.Rendering(screen)
}

func displayTest1() {
	var screen = new(render.Screen)
	screen.Init() // set to black every pixel
	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			if i >= 0 && i < j && j < 2*i {
				screen.FillPixel(i, j, color.White)
			}
		}
	}
	render.Rendering(screen)
}

func displayTest2() {
	var screen = new(render.Screen)
	screen.Init() // set to black every pixel
	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			if j > i*i {
				screen.FillPixel(i, j, color.White)
			}
		}
	}
	render.Rendering(screen)
}
