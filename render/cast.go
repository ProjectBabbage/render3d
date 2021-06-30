// https://en.wikipedia.org/wiki/Phong_reflection_model
package render

import (
	. "broengine/config"
	. "broengine/datatypes"
	"math"
)

// Compute the pixel intensity associated with the ray that intersected something
func compute_intensity(iR IntersectRes, r Ray, scene Scene, Eps float64) float64 {
	var (
		ia float64 = 0
		id float64 = 0
		is float64 = 0
	)
	// common
	p := iR.Vector
	n := iR.Normale

	// ambient
	ka := iR.Ka

	// diffuse
	kd := iR.Kd

	// specular
	ks := iR.Ks
	v := r.Direction()
	a := iR.A

	for _, light := range scene.Lights {
		ia += light.Ia

		// diffuse
		lm := light.Minus(p).Normalize()
		imd := light.Id
		ps_diffuse := lm.ProdScal(n)

		// specular
		rm := lm.Minus(n.Dilate(2 * n.ProdScal(lm)))
		ims := light.Is
		ps_specular := rm.ProdScal(v)

		// Shadow Ray
		SR := NewRay(p.Add(lm.Dilate(Eps)), lm)
		iSR := scene.Intersect(SR)
		inShadow := iSR.HasIntersection && iSR.DistanceToOrigine < light.Distance(p)

		if !inShadow {
			if ps_diffuse > 0 {
				id += imd * ps_diffuse
			}

			if ps_specular > 0 && ps_diffuse >= 0 {
				is += ims * math.Pow(ps_specular, a)
			}
		}
	}

	return ia*ka + id*kd + is*ks
}

// Cast a ray in the scene, return its intensity
func Cast(r Ray, scene Scene, Eps float64) float64 {
	iR := scene.Intersect(r)
	if !iR.HasIntersection {
		return 0 // no intensity
	}
	i := compute_intensity(iR, r, scene, Eps)
	return i
}

// Cast all rays
func CastAll(scene Scene, conf Config) Screen {
	var screen = new(Screen)
	screen.Init(conf) // set every pixel to black

	for i := conf.Lx; i <= conf.Hx; i++ {
		for j := conf.Ly; j <= conf.Hy; j++ {
			ray := NewRay(conf.Eye, conf.Pxy(i, j))
			intensity := Cast(ray, scene, conf.Eps)
			c := ConvertIntensityToGrayScale(intensity)
			screen.FillPixel(i, j, c)
		}
	}
	return *screen
}
