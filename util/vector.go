package util

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Normalize() Vector {
	n := v.Norm()
	x := v.X / n
	y := v.Y / n
	z := v.Z / n
	return Vector{x, y, z}
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector) Minus(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v Vector) Dilate(t float64) Vector {
	return Vector{t * v.X, t * v.Y, t * v.Z}
}

func (v1 Vector) ProdScal(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Returns v1 without its v2 component
func (v1 Vector) Orthogonalize(v2 Vector) Vector {
	v2 = v2.Normalize()
	return v1.Minus(v2.Dilate(v2.ProdScal(v1)))
}

func (v Vector) Print() {
	fmt.Println(v)
}
