package render

import (
	. "broengine/assets"
	. "broengine/config"
	. "broengine/datatypes"
	"image/color"
	"testing"
)

var (
	WHITE = color.Gray{255}
	BLACK = color.Gray{0}
)

func TestMain(m *testing.M) {
	Path = "../assets/"
	m.Run()
}

func TestDisplay1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}
	var screen = new(Screen)
	screen.Init() // set every pixel to black
	size := 100
	for i := -size / 2; i < size/2; i++ {
		for j := 0; j < i; j++ {
			screen.FillPixel(i, j, WHITE)
		}
	}
	RenderScreen(*screen)
}

func TestDisplay2(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}
	var screen = new(Screen)
	screen.Init() // set every pixel to black
	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			if i >= 0 && i < j && j < 2*i {
				screen.FillPixel(i, j, WHITE)
			}
		}
	}
	RenderScreen(*screen)
}

func TestDisplay3(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}
	var screen = new(Screen)
	screen.Init() // set every pixel to black
	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			if j > i*i {
				screen.FillPixel(i, j, WHITE)
			}
		}
	}
	RenderScreen(*screen)
}

// ~ 5ms/op
func BenchmarkCastSphere(b *testing.B) {
	scene := SSphere()
	r := NewRay(Vector{}, Vector{})

	for i := 0; i < b.N; i++ {
		Cast(r, scene)
	}
}

// ~ 8 seconds
func BenchmarkCastAllSpherePlane(b *testing.B) {
	scene := SSpherePlane()
	for i := 0; i < b.N; i++ {
		CastAll(scene)
	}
}
