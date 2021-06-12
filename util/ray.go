package util

// Assume direction to be normalized
type Ray struct {
	origin    Vector
	direction Vector
}

func NewRay(o, d Vector) Ray {
	return Ray{o, d.Normalize()}
}

func (r Ray) Origin() Vector {
	return r.origin
}

func (r Ray) Direction() Vector {
	return r.direction
}
