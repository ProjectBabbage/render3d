package render

import (
	. "broengine/assets"
	. "broengine/config"
	. "broengine/datatypes"
	"testing"
)

func BenchmarkCastSphere(b *testing.B) {
	sphere := ParseStl("../assets/sphere.stl", 1, 1, 1, 1)
	sphere.Translate(Vector{0, 0, 10})
	scene := NewEmptyScene()
	scene.AddObjects(sphere)
	scene.AddLights(Lights...)

	r := NewRay(Vector{0, 0, 0}, Vector{0, 0, 0})
	for i := 0; i < b.N; i++ {
		Cast(r, scene)
	}
}
