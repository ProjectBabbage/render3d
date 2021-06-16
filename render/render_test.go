package render

import (
	. "broengine/assets"
	. "broengine/config"
	. "broengine/datatypes"
	"image/color"
	"testing"
)

func BenchmarkCastSphere(b *testing.B) {
	sphere := Sphere80
	sphere.Translate(Vector{0, 0, 10})
	scene := NewEmptyScene()
	scene.AddObjects(sphere)
	scene.AddLights(L1)

	r := NewRay(Vector{0, 0, 0}, Vector{0, 0, 0})
	for i := 0; i < b.N; i++ {
		Cast(r, scene)
	}
}

func TestDisplay1(t *testing.T) {
	var screen = new(Screen)
	screen.Init() // set to black every pixel
	var pixelColor = color.White
	size := 100
	for i := -size / 2; i < size/2; i++ {
		for j := 0; j < i; j++ {
			screen.FillPixel(i, j, pixelColor)
		}
	}
	Rendering(screen)
}

func TestDisplay2(t *testing.T) {
	var screen = new(Screen)
	screen.Init() // set to black every pixel
	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			if i >= 0 && i < j && j < 2*i {
				screen.FillPixel(i, j, color.White)
			}
		}
	}
	Rendering(screen)
}

func TestDisplay3(t *testing.T) {
	var screen = new(Screen)
	screen.Init() // set to black every pixel
	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			if j > i*i {
				screen.FillPixel(i, j, color.White)
			}
		}
	}
	Rendering(screen)
}
