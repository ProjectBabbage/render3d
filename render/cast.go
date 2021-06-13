package render

import (
	. "broengine/config"
	. "broengine/datatypes"
	"math"
)

func calc_Ia() float64 {
	var i float64 = 0
	for _, light := range Lights {
		i += light.Ia
	}
	return i
}

var Ia = calc_Ia()

func calc_id(inter IntersectRes) float64 {
	var i float64 = 0
	p := inter.Position
	n := inter.Normale
	kd := inter.Kd
	for _, light := range Lights {
		lm := light.Pos.Minus(p).Normalize()
		imd := light.Id
		ps := lm.ProdScal(n)
		if ps > 0 {
			i += kd * imd * ps
		}
	}
	return i
}

func calc_is(inter IntersectRes, r Ray) float64 {
	var i float64 = 0
	p := inter.Position
	n := inter.Normale
	ks := inter.Ks
	v := r.Direction()
	a := inter.A
	for _, light := range Lights {
		lm := light.Pos.Minus(p).Normalize()
		rm := lm.Minus(n.Dilate(2 * n.ProdScal(lm)))
		ims := light.Is
		i += ks * ims * math.Pow(rm.ProdScal(v), a)
	}
	return i
}

func Cast(r Ray, scene Scene) float64 {
	inter := scene.Intersect(r)
	if !inter.HasIntersection {
		return 0 // no intensity
	}
	ia := inter.Ka * Ia
	id := calc_id(inter)
	i := ia + id
	return i
}
