package util

import (
	"fmt"
)

type Triangle struct {
	P1, P2, P3 Vector
	N          Vector
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
	p1 := t.P1
	p2 := t.P2
	p3 := t.P3

	b := isSameSide(p, p1, p2, p3) &&
		isSameSide(p, p2, p3, p1) &&
		isSameSide(p, p3, p1, p2)
	return b
}

func (t Triangle) projection(v Vector) Vector {
	n := t.N.Normalize()
	return v.Minus(n.Dilate(v.Minus(t.P1).ProdScal(n)))
}

func (t Triangle) Intersect(v Vector) Inter {
	p := t.projection(v)
	b := t.contains(p)
	return Inter{b, p}
}

func (t Triangle) Translate(v Vector) {
	t.P1 = t.P1.Add(v)
	t.P2 = t.P2.Add(v)
	t.P3 = t.P3.Add(v)
}
