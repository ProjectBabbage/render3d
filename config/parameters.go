package config

import (
	. "broengine/util"
)

// Eye is at origin

// z-position of the screen
const D float64 = 100

// Height and Width of the screen
const HWidth float64 = 50
const HHeight float64 = 50

// Position of the light
var Light = Vector{100, 0, 200}

const PixelsWidth = 20
const Lx = -PixelsWidth / 2
const Hx = PixelsWidth / 2

const PixelsHeight = 10
const Ly = -PixelsHeight / 2
const Hy = PixelsHeight / 2

func Pxy(i, j int) Vector {
	px := Vector{HWidth / PixelsWidth, 0, 0}
	py := Vector{0, HHeight / PixelsHeight, 0}
	d := Vector{0, 0, D}
	return d.Add(px.Dilate(float64(i))).Add(py.Dilate(float64(j)))
}
