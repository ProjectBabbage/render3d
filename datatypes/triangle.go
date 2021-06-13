package datatypes

import (
	"fmt"
)

// Assume P1, P2, P3 to be distinct and N to be normalized
type Triangle struct {
	p1, p2, p3    Vector
	n             Vector
	ka, kd, ks, a float64
}

func NewTriangle(p1, p2, p3, n Vector, ka, kd, ks, a float64) Triangle {
	return Triangle{p1, p2, p3, n.Normalize(), ka, kd, ks, a}
}

func (t Triangle) Print() {
	fmt.Printf("Triangle : %+v\n", t)
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
		return IntersectRes{false, 0, Vector{0, 0, 0}, Vector{0, 0, 0}, 0, 0, 0, 0}
	} else {
		d := t.p1.Minus(x).ProdScal(t.n) / u_n
		p := x.Add(u.Dilate(d))
		b := t.contains(p) && d >= 0
		return IntersectRes{b, d, p, t.n, t.ka, t.kd, t.ks, t.a}
	}
}

func (t Triangle) Translate(v Vector) Surface {
	t.p1 = t.p1.Add(v)
	t.p2 = t.p2.Add(v)
	t.p3 = t.p3.Add(v)
	return t
}

func ConvertTriangleListIntoSurfaceList(triangles []Triangle) []Surface {
	var surfaces = []Surface{}
	for _, triangle := range triangles {
		surfaces = append(surfaces, triangle)
	}
	return surfaces
}

func (t Triangle) RecomputeNormal() Triangle {
	u := t.p2.Minus(t.p1)
	v := t.p3.Minus(t.p1)
	t.n = Vector{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}.Normalize()
	return t
}
