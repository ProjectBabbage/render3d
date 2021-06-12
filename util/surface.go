package util

// When hasIntersection, we should have distance >= 0
type IntersectRes struct {
	hasIntersection bool
	distance        float64
	position        Vector
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
		if I.hasIntersection && (!res.hasIntersection || res.hasIntersection && I.distance < res.distance) {
			res.hasIntersection = I.hasIntersection
			res.distance = I.distance
			res.position = I.position
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
