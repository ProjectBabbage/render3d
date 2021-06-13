package render

import (
	// . "broengine/config"
	. "broengine/util"
)

var p1 = Vector{10, 10, 200}
var p2 = Vector{-10, 10, 200}
var p3 = Vector{-10, -10, 200}
var n = Vector{0, 0, -1}
var t = NewTriangle(p1, p2, p3, n, 5, 2, 3, 1)
var scene = SurfaceFromSurfaces([]Surface{t})

func Cast(r Ray) {

}
