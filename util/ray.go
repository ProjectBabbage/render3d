package util

// Assume direction to be normalized
type Ray struct {
	origin    Vector
	direction Vector
}

func NewRay(origin, direction Vector) Ray {
	return Ray{origin, direction.Normalize()}
}

func (r Ray) Origin() Vector {
	return r.origin
}

func (r Ray) Direction() Vector {
	return r.direction
}
