package main

import (
	. "broengine/config"
	"broengine/render"
	. "broengine/util"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	render.Rendering()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}

func testIntesection() {
	p1 := Vector{20, 20, 200}
	p2 := Vector{-15, -30, 200}
	p3 := Vector{40, -10, 200}
	n := Vector{0, 0, 1}
	t := Triangle{p1, p2, p3, n}

	scene := SurfaceFromSurfaces([]Surface{t})

	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			point := scene.Intersect(Ray{Eye, Pxy(i, j)})
			fmt.Println(point)
		}
	}
}
