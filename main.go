package main

import (
	. "broengine/config"
	"broengine/render"
	. "broengine/util"
	"image/color"
)

func main() {
	// load the sphere
	var cube []Surface = ConvertTriangleListIntoSurfaceList(
		ParseStl("assets/cube.stl", 1, 1, 1, 1),
	)

	// p1 := Vector{10, 10, 200}
	// p2 := Vector{-30, 10, 200}
	// p3 := Vector{-10, -40, 200}
	// n := Vector{0, 0, -1}
	// t := NewTriangle(p1, p2, p3, n, 5, 2, 3, 1)

	scene := SurfaceFromSurfaces(cube)

	scene = scene.Translate(Vector{0, 0, 40})

	var screen = new(render.Screen)

	for i := Lx; i < Hx; i++ {
		for j := Ly; j < Hy; j++ {
			point := scene.Intersect(NewRay(Eye, Pxy(i, j)))
			var pixelColor = color.Black
			if point.HasIntersection {
				pixelColor = color.White
			}
			screen.FillPixel(i, j, pixelColor)
		}
	}
	render.Rendering(screen)
}
