package stl

import (
	. "broengine/datatypes"
	"fmt"
	"log"

	"github.com/hschendel/stl"
)

// Get the Triangles from a .stl file.

func Parse(filepath string, a float64, ka, kd, ks Col) Object {

	s, err := stl.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error when parsing the stl file %s", filepath)
	}
	var triangles []Surface = make([]Surface, len(s.Triangles))
	for i := 0; i < len(s.Triangles); i++ {
		tri := convertTriangle(s.Triangles[i], a, ka, kd, ks)
		triangles[i] = &tri
	}
	return Object{triangles}
}

func convertTriangle(t stl.Triangle, a float64, ka, kd, ks Col) Triangle {
	return NewTriangle(
		convertVector(t.Vertices[0]),
		convertVector(t.Vertices[1]),
		convertVector(t.Vertices[2]),
		convertVector(t.Normal), a, ka, kd, ks)
}

func convertVector(v stl.Vec3) Vector {
	return Vector{X: float64(v[0]), Y: float64(v[1]), Z: float64(v[2])}
}
