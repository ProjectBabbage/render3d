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
	StlPath = "../assets/"
	m.Run()
}

func TestDisplay1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}
	conf := NewConfig(Config{})
	var screen = NewScreen(conf.PixelsX, conf.PixelsY)
	size := 100
	for i := -size / 2; i < size/2; i++ {
		for j := 0; j < i; j++ {
			screen.FillPixel(i, j, WHITE)
		}
	}
	RenderScreen(&screen, NewConfig(Config{}))
}

func TestDisplay2(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}

	conf := NewConfig(Config{})
	var screen = NewScreen(conf.PixelsX, conf.PixelsY)
	for i := conf.Lx(); i <= conf.Hx(); i++ {
		for j := conf.Ly(); j <= conf.Hy(); j++ {
			if i >= 0 && i < j && j < 2*i {
				screen.FillPixel(i, j, WHITE)
			}
		}
	}
	RenderScreen(&screen, conf)
}

func TestDisplay3(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Display tests in short mode")
	}

	conf := NewConfig(Config{})
	var screen = NewScreen(conf.PixelsX, conf.PixelsY)
	for i := conf.Lx(); i <= conf.Hx(); i++ {
		for j := conf.Ly(); j <= conf.Hy(); j++ {
			if j > i*i {
				screen.FillPixel(i, j, WHITE)
			}
		}
	}
	RenderScreen(&screen, conf)
}

// ~ 6 ms/op
func BenchmarkCastSphere(b *testing.B) {
	scene, conf := SSphere()
	r := NewRay(Vector{}, Vector{})

	for i := 0; i < b.N; i++ {
		Cast(r, scene, conf.Eps)
	}
}

// ~ 0.3 ms/op
func BenchmarkCastTrueSphere(b *testing.B) {
	scene, conf := STrueSphere()
	r := NewRay(Vector{}, Vector{})

	for i := 0; i < b.N; i++ {
		Cast(r, scene, conf.Eps)
	}
}

// ~ 16 s for 500x500 px
func BenchmarkCastAllSpherePlane(b *testing.B) {
	scene, conf := SSpherePlane()
	for i := 0; i < b.N; i++ {
		CastAll(scene, conf)
	}
}
