package util

import (
	"fmt"
)

// Assume P1, P2, P3 to be distinct and N to be normalized
type Triangle struct {
	p1, p2, p3 Vector
	n          Vector
}

func NewTriangle(p1, p2, p3, n Vector) Triangle {
	return Triangle{p1, p2, p3, n.Normalize()}
}

func (v Triangle) Print() {
	fmt.Println(v)
}

// Indicate if x is on the same side of p3 compared to the line p1-p2
func isSameSide(x, p1, p2, p3 Vector) bool {
	u := p2.Minus(p1)
	v := p3.Minus(p1)
	v = v.Orthogonalize(u)
	vx := x.Minus(p1)
	b := vx.ProdScal(v) >= 0
	return b
}

func (t Triangle) contains(p Vector) bool {
	p1 := t.p1
	p2 := t.p2
	p3 := t.p3

	b := isSameSide(p, p1, p2, p3) &&
		isSameSide(p, p2, p3, p1) &&
		isSameSide(p, p3, p1, p2)
	return b
}

func (t Triangle) Intersect(r Ray) IntersectRes {
	x := r.Origin()
	u := r.Direction()
	u_n := u.ProdScal(t.n)
	if u_n >= 0 {
		return IntersectRes{false, 0, Vector{0, 0, 0}}
	}
	{
		d := x.Minus(t.p1).ProdScal(t.n) / u_n
		p := x.Add(u.Dilate(d))
		b := t.contains(p)
		return IntersectRes{b, d, p}
	}
}

func (t Triangle) Translate(v Vector) {
	t.p1 = t.p1.Add(v)
	t.p2 = t.p2.Add(v)
	t.p3 = t.p3.Add(v)
}
