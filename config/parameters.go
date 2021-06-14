package config

import (
	. "broengine/datatypes"
)

// Eye is at origin
var Eye = Vector{X: 0, Y: 0, Z: 0}

// z-Position of the screen
const D float64 = 2

// Height and Width of the screen
const ScreenX float64 = 1
const ScreenY float64 = 1

// Position of the light
var Lights = []Light{Light{25, 200, 3, Vector{-10, -50, -50}}}

const PixelsY = 1000
const Ly = -PixelsY / 2
const Hy = PixelsY / 2

const PixelsX = 1000
const Lx = -PixelsX / 2
const Hx = PixelsX / 2

// Vector that comes from the eye and goes through the (i,j) pixel
func Pxy(i, j int) Vector {
	px := Vector{ScreenX / float64(PixelsX), 0, 0}
	py := Vector{0, ScreenY / float64(PixelsY), 0}
	d := Vector{0, 0, D}
	return d.Add(px.Dilate(float64(i))).Add(py.Dilate(float64(j)))
}
