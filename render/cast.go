package render

import (
	. "broengine/datatypes"
	"math"
)

func calc_Ia(scene Scene) float64 {
	var i float64 = 0
	for _, light := range scene.Lights {
		i += light.Ia
	}
	return i
}

func calc_id(iR IntersectRes, scene Scene) float64 {
	var i float64 = 0
	p := iR.Vector
	n := iR.Normale
	kd := iR.Kd
	for _, light := range scene.Lights {
		lm := light.Minus(p).Normalize()
		imd := light.Id
		ps := lm.ProdScal(n)
		SR := NewRay(iR.Vector, lm) // Shadow Ray
		iSR := scene.Intersect(SR)
		inShadow := iSR.HasIntersection &&
			iSR.DistanceToOrigine < light.Distance(p)
		if ps > 0 && !inShadow {
			i += kd * imd * ps
		}
	}
	return i
}

func calc_is(iR IntersectRes, r Ray, scene Scene) float64 {
	var i float64 = 0
	p := iR.Vector
	n := iR.Normale
	ks := iR.Ks
	v := r.Direction()
	a := iR.A
	for _, light := range scene.Lights {
		lm := light.Minus(p).Normalize()
		rm := lm.Minus(n.Dilate(2 * n.ProdScal(lm)))
		ims := light.Is
		i += ks * ims * math.Pow(rm.ProdScal(v), a)
	}
	return i
}

func Cast(r Ray, scene Scene) float64 {
	iR := scene.Intersect(r)
	if !iR.HasIntersection {
		return 0 // no intensity
	}
	ia := iR.Ka * calc_Ia(scene)
	id := calc_id(iR, scene)
	i := ia + id
	return i
}
