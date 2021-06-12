package util

type Inter struct {
	hasIntersection bool
	position        Vector
}

type Reflector interface {
	Intersect(Vector) Inter
	Translate(Vector)
	Print()
}
