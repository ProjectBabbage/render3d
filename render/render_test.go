package render

import (
	. "broengine/assets"
	. "broengine/config"
	. "broengine/datatypes"
	"image/color"
	"testing"
)

func TestConvertIndexToScreenIndex(t *testing.T) {
	PixelsX := 200
	PixelsY := 100

	var i, j = 10, 20
	var expected_I, expected_J = 110, 70
	var I, J int
	I, J = convertIndexToScreenIndex(i, j, PixelsX, PixelsY)
	t.Log(i, j)
	t.Log(I, J)

	if I != expected_I || J != expected_J {
		t.Errorf("Indexes conversion to Screen Indexes error.")
	}

}

func TestDisplay1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}
	var screen = new(Screen)
	screen.Init() // set every pixel to black
	var pixelColor = color.White
	size := 100
	for i := -size / 2; i < size/2; i++ {
		for j := 0; j < i; j++ {
			screen.FillPixel(i, j, pixelColor)
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
				screen.FillPixel(i, j, color.White)
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
				screen.FillPixel(i, j, color.White)
			}
		}
	}
	RenderScreen(*screen)
}

// ~ 5ms/op
func BenchmarkCastSphere(b *testing.B) {
	sphere := ParseStl("../assets/sphere.stl", 1, 1, 1, 1)
	sphere.Translate(Vector{0, 0, 10})
	scene := NewEmptyScene()
	scene.AddObjects(sphere)
	scene.AddLights(L1)

	r := NewRay(Vector{0, 0, 0}, Vector{0, 0, 0})
	for i := 0; i < b.N; i++ {
		Cast(r, scene)
	}
}

// ~ 8 seconds
func BenchmarkCastAllCubePlane(b *testing.B) {
	Path = "../assets/"
	scene := SCubePlane()
	for i := 0; i < b.N; i++ {
		CastAll(scene)
	}
}
