package util

type IntersectRes struct {
	hasIntersection bool
	position        Vector
}

type Surface interface {
	Intersect(Vector) IntersectRes
	Translate(Vector)
	Print()
}
