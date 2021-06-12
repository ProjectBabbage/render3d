package util

type IntersectRes struct {
	hasIntersection bool
	distance        float64
	position        Vector
}

type Surface interface {
	Intersect(Ray) IntersectRes
	Translate(Vector)
	Print()
}

type Scene = []Surface

// func (s Scene) Intersect(r Ray) IntersectRes {

// }
