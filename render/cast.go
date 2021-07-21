// https://en.wikipedia.org/wiki/Phong_reflection_model
package render

import (
	. "broengine/config"
	. "broengine/datatypes"
	"math"
	"runtime"
	"sync"
)

// compute the pixel intensity associated with the ray that intersected something.
func compute_intensity(iR IntersectRes, r Ray, scene Scene) Col {
	var (
		ia = Col{}
		id = Col{}
		is = Col{}
	)
	// common
	p := iR.Vector
	n := iR.Normale

	// ambient
	ka := iR.Mat.Ka

	// diffuse
	kd := iR.Mat.Kd

	// specular
	ks := iR.Mat.Ks
	v := r.Direction()
	a := iR.Mat.A

	// reflection
	// kr := iR.Kr
	// new_direction := v.Minus(n.Dilate(2 * v.ProdScal(n)))
	// new_ray := Ray{p, new_direction}

	for _, light := range scene.Lights {
		ia = ia.AddColor(light.Ia)

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
		inShadow := iSR.HasIntersection && iSR.DistanceToOrigin < light.Distance(p)

		if !inShadow {
			if ps_diffuse > 0 {
				id = id.AddColor(imd.DilateColor(ps_diffuse))
			}

			if ps_specular > 0 && ps_diffuse >= 0 {
				is = is.AddColor(ims.DilateColor(math.Pow(ps_specular, a)))
			}
		}

		// reflection
	}

	return ia.MulColor(ka).AddColor(id.MulColor(kd)).AddColor(is.MulColor(ks))
}

// Cast a ray in the scene, return its intensity.
func Cast(r Ray, scene Scene) Col {
	iR := scene.Intersect(r)
	if !iR.HasIntersection {
		return Col{} // no intensity
	}
	i := compute_intensity(iR, r, scene)
	return i
}

func CastWorker(i1, i2 int, screen *Screen, scene Scene, conf Config, wg *sync.WaitGroup) {
	for i := i1; i < i2; i++ {
		for j := conf.Ly(); j <= conf.Hy(); j++ {
			ray := NewRay(conf.Eye, conf.Pxy(i, j))
			c := Cast(ray, scene)
			screen.FillPixel(i, j, c)
		}
	}
	wg.Done()
}

// Cast all rays.
func CastAll(scene Scene, conf Config) Screen {
	screen := NewScreen(conf.PixelsX, conf.PixelsY) // set every pixel to black

	var wg sync.WaitGroup

	nbSlices := runtime.GOMAXPROCS(0) * 10
	L := conf.PixelsX / nbSlices
	i1 := conf.Lx()
	i2 := i1 + L
	wg.Add(nbSlices)
	for k := 0; k < nbSlices; k++ {
		go CastWorker(i1, i2, &screen, scene, conf, &wg)
		i1 = i1 + L
		i2 = i2 + L
		if k == nbSlices-2 {
			i2 = conf.Hx() + 1
		}
	}
	wg.Wait()
	return screen
}
