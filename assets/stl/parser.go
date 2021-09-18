package stl

import (
	. "broengine/datatypes"
	"fmt"
	"log"

	"github.com/hschendel/stl"
)

// Get the Triangles from a .stl file, and give them the same Material mat
func Parse(filepath string, mat Material) Object {
	fmt.Println("Using", filepath)
	s, err := stl.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("(From us) Error when parsing the stl file %s", filepath)
	}
	var triangles []Surface = make([]Surface, len(s.Triangles))
	for i := 0; i < len(s.Triangles); i++ {
		tri := convertTriangle(s.Triangles[i], mat)
		triangles[i] = &tri
	}
	return Object{triangles}
}

// Converts the stl Triangle to our triangle type (float64 vectors and with a material)
func convertTriangle(t stl.Triangle, mat Material) Triangle {
	return NewTriangle(
		convertVector(t.Vertices[0]),
		convertVector(t.Vertices[1]),
		convertVector(t.Vertices[2]),
		convertVector(t.Normal),
		mat,
	)
}

// Converts the stl package Vector point from float32 to float64
func convertVector(v stl.Vec3) Vector {
	return Vector{X: float64(v[0]), Y: float64(v[1]), Z: float64(v[2])}
}
