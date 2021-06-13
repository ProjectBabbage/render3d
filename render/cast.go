package render

import (
	. "broengine/config"
	. "broengine/util"
	"math"
)

var p1 = Vector{10, 10, 200}
var p2 = Vector{-10, 10, 200}
var p3 = Vector{-10, -10, 200}
var n = Vector{0, 0, -1}
var t = NewTriangle(p1, p2, p3, n, 5, 2, 3, 1)
var scene = Scene{[]Surface{t}}

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
		i += kd * imd * lm.ProdScal(n)
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

func Cast(r Ray) float64 {
	inter := scene.Intersect(r)
	ia := inter.Ka * Ia
	id := calc_id(inter)
	is := calc_is(inter, r)
	i := ia + id + is
	return i
}
