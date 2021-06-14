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

type Object struct {
	Surfaces []Surface
}

func (obj Object) Intersect(r Ray) IntersectRes {
	res := NoIntersection
	for _, surf := range obj.Surfaces {
		I := surf.Intersect(r)
		res.Update(I)
	}
	return res
}

func (obj Object) Translate(v Vector) Object {
	surfaces := []Surface{}
	for _, surf := range obj.Surfaces {
		surfaces = append(surfaces, surf.Translate(v))
	}
	return Object{surfaces}
}

func (obj Object) Print() {
	for _, surf := range obj.Surfaces {
		surf.Print()
	}
}

type Scene struct {
	objects Object
	Lights  []Light
}

func NewEmptyScene() Scene {
	return Scene{Object{[]Surface{}}, []Light{}}
}

func (s *Scene) AddObjects(objs ...Object) {
	surfs := s.objects.Surfaces
	for _, obj := range objs {
		surfs = append(surfs, obj.Surfaces...)
	}
	s.objects = Object{surfs}
}

func (s *Scene) AddLights(lights ...Light) {
	s.Lights = append(s.Lights, lights...)
}

func (s *Scene) TranslateObjects(v Vector) {
	(*s).objects = s.objects.Translate(v)
}

func (s Scene) Intersect(r Ray) IntersectRes {
	return s.objects.Intersect(r)
}
