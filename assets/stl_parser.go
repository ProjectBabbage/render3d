package assets

import (
	. "broengine/datatypes"
	"fmt"
	"github.com/hschendel/stl"
	"log"
)

// Get the Triangles from a .stl file.
func ParseStl(filepath string, ka, kd, ks, a float64) Object {

	s, err := stl.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error when parsing the stl file %s", filepath)
	}
	var triangles []Surface = make([]Surface, len(s.Triangles))
	for i := 0; i < len(s.Triangles); i++ {
		tri := convertTriangle(s.Triangles[i], ka, kd, ks, a)
		triangles[i] = &tri
	}
	return Object{triangles}
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
