package main

import (
	. "broengine/config"
	"broengine/render"
	. "broengine/util"
	"image/color"
)

func main() {
	p1 := Vector{20, 20, 200}
	p2 := Vector{-15, -30, 200}
	p3 := Vector{40, -10, 200}
	n := Vector{0, 0, 1}
	t := Triangle{p1, p2, p3, n}

	scene := SurfaceFromSurfaces([]Surface{t})

	var screen = new(render.Screen)

	for i := Lx; i < Hx; i++ {
		for j := Ly; j < Hy; j++ {
			point := scene.Intersect(Ray{Eye, Pxy(i, j)})
			var pixelColor = color.Black
			if point.HasIntesection {
				pixelColor = color.White
			}
			screen.FillPixel(i, j, pixelColor)
		}
	}
	render.Rendering(screen)
	for i := 0; i < len(screen.Pixels); i++ {

	}
}
