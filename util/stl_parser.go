package util

import (
	"github.com/hschendel/stl"
	"log"
)

// Get the Triangles from a .stl file.
func ParseStl(filepath string, ka, kd, ks, a float64) []Triangle {

	s, err := stl.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error when parsing the stl file %s", filepath)
		return nil
	}
	triangles := make([]Triangle, len(s.Triangles))
	for i := 0; i < len(s.Triangles); i++ {
		triangles[i] = convertTriangle(s.Triangles[i], ka, kd, ks, a)
	}
	return triangles
}

func convertTriangle(t stl.Triangle, ka, kd, ks, a float64) Triangle {
	return NewTriangle(
		convertVector(t.Vertices[0]),
		convertVector(t.Vertices[1]),
		convertVector(t.Vertices[2]),
		convertVector(t.Normal), ka, kd, ks, a)
}

func convertVector(v stl.Vec3) Vector {
	return Vector{X: float64(v[0]), Y: float64(v[1]), Z: float64(v[2])}
}
