package config

import (
	. "broengine/util"
)

// Eye is at origin
var Eye = Vector{0, 0, 0}

// z-position of the screen
const D float64 = 100

// Height and Width of the screen
const ScreenWidth float64 = 50
const ScreenHeight float64 = 50

// Position of the light
var Light = Vector{100, 0, 200}

// Pixel density
const PixelsWidth = 20
const Lx = -PixelsWidth / 2
const Hx = PixelsWidth / 2

const PixelsHeight = 10
const Ly = -PixelsHeight / 2
const Hy = PixelsHeight / 2

// Vector that comes from the eye and goes through the (i,j) pixel
func Pxy(i, j int) Vector {
	px := Vector{ScreenWidth / PixelsWidth, 0, 0}
	py := Vector{0, ScreenHeight / PixelsHeight, 0}
	d := Vector{0, 0, D}
	return d.Add(px.Dilate(float64(i))).Add(py.Dilate(float64(j)))
}
