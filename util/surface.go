package util

// When HasIntesection, we should have Distance >= 0
type IntersectRes struct {
	HasIntesection bool
	Distance       float64
	Position       Vector
}

type Surface interface {
	Intersect(Ray) IntersectRes
	Translate(Vector)
	Print()
}

type scene struct {
	support []Surface
}

func (s scene) Intersect(r Ray) IntersectRes {
	res := IntersectRes{false, 0, Vector{0, 0, 0}}
	for _, surf := range s.support {
		I := surf.Intersect(r)
		if I.HasIntesection && (!res.HasIntesection || res.HasIntesection && I.Distance < res.Distance) {
			res.HasIntesection = I.HasIntesection
			res.Distance = I.Distance
			res.Position = I.Position
		}
	}
	return res
}

func (s scene) Translate(v Vector) {
	for _, surf := range s.support {
		surf.Translate(v)
	}
}

func (s scene) Print() {
	for _, surf := range s.support {
		surf.Print()
	}
}

func SurfaceFromSurfaces(surfs []Surface) Surface {
	return scene{surfs}
}
