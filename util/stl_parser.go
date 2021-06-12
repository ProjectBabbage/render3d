package util

import (
	"log"

	"github.com/hschendel/stl"
)

// Get the Triangles from a .stl file.
func ParseStl(filepath string) []Triangle {

	s, err := stl.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error when parsing the stl file %s", filepath)
		return nil
	}
	triangles := make([]Triangle, len(s.Triangles))
	for i := 0; i < len(s.Triangles); i++ {
		triangles[i] = convertTriangle(s.Triangles[i])
	}
	return triangles
}

func convertTriangle(t stl.Triangle) Triangle {
	return Triangle{
		P1: convertVector(t.Vertices[0]),
		P2: convertVector(t.Vertices[1]),
		P3: convertVector(t.Vertices[2]),
		N:  convertVector(t.Normal),
	}
}

func convertVector(v stl.Vec3) Vector {
	return Vector{X: float64(v[0]), Y: float64(v[1]), Z: float64(v[2])}
}
