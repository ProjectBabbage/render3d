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

// var NoIntersection = IntersectRes{false, 0, Vector{0, 0, 0}, Vector{0, 0, 0}, 0, 0, 0, 0}

type Surface interface {
	Intersect(Ray) IntersectRes
	Translate(Vector) Surface
	Print()
}

type Scene struct {
	Surfaces []Surface
}

func (s Scene) Intersect(r Ray) IntersectRes {
	res := IntersectRes{false, 0, Vector{0, 0, 0}, Vector{0, 0, 0}, 0, 0, 0, 0}
	for _, surf := range s.Surfaces {
		I := surf.Intersect(r)
		if I.HasIntersection && (!res.HasIntersection || res.HasIntersection && I.Distance < res.Distance) {
			res = I
		}

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
