package util

// When HasIntesection, we should have Distance >= 0
type IntersectRes struct {
	HasIntersection bool
	Distance        float64
	Position        Vector
	Normale         Vector
	Ka, Kd, Ks      float64
	A               float64
}

var NoIntersection = IntersectRes{false, 0, Vector{0, 0, 0}, Vector{0, 0, 0}, 0, 0, 0, 0}

type Surface interface {
	Intersect(Ray) IntersectRes
	Translate(Vector)
	Print()
}

type surfaces struct {
	support []Surface
}

func (s surfaces) Intersect(r Ray) IntersectRes {
	res := NoIntersection
	for _, surf := range s.support {
		I := surf.Intersect(r)
		if I.HasIntersection && (!res.HasIntersection || res.HasIntersection && I.Distance < res.Distance) {
			res.HasIntersection = I.HasIntersection
			res.Distance = I.Distance
			res.Position = I.Position
		}

	}
	return res
}

func (s surfaces) Translate(v Vector) {
	for _, surf := range s.support {
		surf.Translate(v)
	}
}

func (s surfaces) Print() {
	for _, surf := range s.support {
		surf.Print()
	}
}

func SurfaceFromSurfaces(surfs []Surface) Surface {
	return surfaces{surfs}
}
