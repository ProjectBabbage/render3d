package datatypes

import (
	"fmt"
	"math"
)

type Sphere struct {
	C   Vector
	R   float64
	mat Material
}

func NewSphere(c Vector, r float64, mat Material) Sphere {
	return Sphere{c, r, mat}
}

func (s Sphere) Print() {
	fmt.Println(s)
}

func (s *Sphere) Translate(v Vector) {
	s.C = s.C.Add(v)
}

func (s *Sphere) Dilate(x float64) {
	s.C = s.C.Dilate(x)
	s.R = s.R * x
}

func (s *Sphere) Rotate(axis int, d float64) {
	s.C = s.C.Rotate(axis, d)
}

func (s Sphere) Intersect(r Ray) IntersectRes {
	x := r.Origin()
	u := r.Direction()
	v := s.C.Minus(x)
	ps := u.ProdScal(v)
	delta := ps*ps + s.R*s.R - v.SquareNorm()
	if delta < 0 {
		return NoIntersection
	}
	sd := math.Sqrt(delta)
	d1 := ps - sd
	d2 := ps + sd
	if d2 < 0 {
		return NoIntersection
	}
	d := d1
	fromOutside := true
	if d1 < 0 {
		d = d2
		fromOutside = false
	}
	p := x.Add(u.Dilate(d))
	n := p.Minus(s.C).Normalize()
	if !fromOutside {
		n = Vector{}.Minus(n)
	}
	return IntersectRes{p, true, d, n, s.mat}
}
