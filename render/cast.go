// https://en.wikipedia.org/wiki/Phong_reflection_model
package render

import (
	"broengine/config"
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

		ps := rm.ProdScal(v)
		shadowRay := NewRay(p, lm)
		inShadow := scene.Intersect(shadowRay).HasIntersection
		if ps > 0 && !inShadow {
			i += ks * ims * math.Pow(ps, a)
		}
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
	if id >= 0 {
		is := calc_is(iR, r, scene)
		i += is
	}
	return i
}

func CastAll(scene Scene) Screen {
	var screen = new(Screen)
	screen.Init() // set every pixel to black

	for i := config.Lx; i <= config.Hx; i++ {
		for j := config.Ly; j <= config.Hy; j++ {
			ray := NewRay(config.Eye, config.Pxy(i, j))
			intensity := Cast(ray, scene)
			c := ConvertIntensityToGrayScale(intensity)
			screen.FillPixel(i, j, c)
		}
	}
	return *screen
}
