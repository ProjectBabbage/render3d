package datatypes

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

func (I1 *IntersectRes) Update(I2 IntersectRes) {
	if !I1.HasIntersection || I2.HasIntersection && I2.Distance < I1.Distance {
		*I1 = I2
	}
}

type Surface interface {
	Intersect(Ray) IntersectRes
	Translate(Vector) Surface
	Print()
}

type Scene struct {
	Surfaces []Surface
}

func (s Scene) Intersect(r Ray) IntersectRes {
	res := NoIntersection
	for _, surf := range s.Surfaces {
		I := surf.Intersect(r)
		res.Update(I)
	}
	return res
}

func (s Scene) Translate(v Vector) Scene {
	surfacesList := []Surface{}
	for _, surf := range s.Surfaces {
		surfacesList = append(surfacesList, surf.Translate(v))
	}
	return Scene{surfacesList}
}

func (s Scene) Print() {
	for _, surf := range s.Surfaces {
		surf.Print()
	}
}
