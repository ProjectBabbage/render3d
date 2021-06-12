package util

// When hasIntersection, we should have distance >= 0
type IntersectRes struct {
	hasIntersection bool
	distance        float64
	position        Vector
}

func (I1 IntersectRes) Update(I2 IntersectRes) {
	if !I1.hasIntersection || I2.hasIntersection && I2.distance < I1.distance {
		I1 = I2
	}
}

var NoIntersection = IntersectRes{false, 0, Vector{0, 0, 0}}

type Surface interface {
	Intersect(Ray) IntersectRes
	Translate(Vector)
	Print()
}

type scene struct {
	support []Surface
}

func (s scene) Intersect(r Ray) IntersectRes {
	res := NoIntersection
	for _, surf := range s.support {
		I := surf.Intersect(r)
		res.Update(I)
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
