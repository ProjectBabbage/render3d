package config

import (
	. "broengine/datatypes"
)

type Config struct {
	// The final render backend, "gpu" or "cpu"
	RenderBackend string
	// Do the computations modulo Eps(ilon)
	Eps float64
	// Where is the Eye
	Eye Vector
	// z-Position of the screen
	D float64
	// Height and Width of the screen
	ScreenX, ScreenY float64
	// Size of the screen in pixel
	PixelsX, PixelsY int
	// (-/+ PixelY/2)
	Ly, Hy int
	// (-/+ PixelX/2)
	Lx, Hx int
}

// Create a default config, overriding default config fields with
// the one present in override_conf
func NewConfig(override_conf Config) Config {

	var (
		PixelsX = 500
		PixelsY = 500
	)

	config := Config{
		RenderBackend: "cpu",
		Eps:           0.0001,
		// Eye is at origin
		Eye:     Vector{X: 0, Y: 0, Z: 0},
		D:       2,
		ScreenX: 1, ScreenY: 1,
		PixelsX: PixelsX, PixelsY: PixelsY,

		Ly: -PixelsY / 2, Hy: PixelsY / 2,
		Lx: -PixelsX / 2, Hx: PixelsX / 2,
	}

	// TODO there must be a cleaner way to do this:
	if override_conf.RenderBackend != "" {
		config.RenderBackend = override_conf.RenderBackend
	}
	if override_conf.Eps != 0 {
		config.Eps = override_conf.Eps
	}
	if override_conf.Eye.X != 0 && override_conf.Eye.Y != 0 && override_conf.Eye.Z != 0 {
		config.Eye = override_conf.Eye
	}
	if override_conf.D != 0 {
		config.D = override_conf.D
	}
	if override_conf.ScreenX != 0 {
		config.ScreenX = override_conf.ScreenX
	}
	if override_conf.ScreenY != 0 {
		config.ScreenY = override_conf.ScreenY
	}
	if override_conf.PixelsX != 0 {
		config.PixelsX = override_conf.PixelsX
		config.Lx = -config.PixelsX / 2
		config.Hx = config.PixelsX / 2
	}
	if override_conf.PixelsY != 0 {
		config.PixelsY = override_conf.PixelsY
		config.Ly = -config.PixelsY / 2
		config.Hy = config.PixelsY / 2
	}

	return config
}

// Vector that comes from the eye and goes through the (i,j) pixel
func (conf Config) Pxy(i, j int) Vector {
	px := Vector{X: conf.ScreenX / float64(conf.PixelsX), Y: 0, Z: 0}
	py := Vector{X: 0, Y: conf.ScreenY / float64(conf.PixelsY), Z: 0}
	d := Vector{X: 0, Y: 0, Z: conf.D}
	return d.Add(px.Dilate(float64(i))).Add(py.Dilate(float64(j)))
}
