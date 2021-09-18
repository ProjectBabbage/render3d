package config

import (
	. "broengine/datatypes"
)

type Config struct {
	// The final render backend, "gpu" or "cpu"
	RenderBackend string
	// Anti-aliasing
	Msaa int
	// Where is the Eye
	Eye Vector
	// z-Position of the screen
	D float64
	// Height and Width of the screen
	ScreenX, ScreenY float64
	// Size of the screen in pixel
	PixelsX, PixelsY int
	// Reflection and Refraction bounce limit
	MaxBounces int
	// Whether or not to save the render as a png in output/
	SaveAsPNG bool
}

// Create a default config, overriding default config fields with
// the one present in override_conf.
func NewConfig(override_conf Config) Config {

	var (
		PixelsX = 500
		PixelsY = 500
	)

	config := Config{
		Msaa: 1,
		// Eye is at origin
		Eye:     Vector{X: 0, Y: 0, Z: 0},
		D:       2,
		ScreenX: 1, ScreenY: 1,
		PixelsX: PixelsX, PixelsY: PixelsY,
		MaxBounces: 2,
		SaveAsPNG:  false,
	}

	// TODO there must be a cleaner way to do this:
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
	}
	if override_conf.PixelsY != 0 {
		config.PixelsY = override_conf.PixelsY
	}
	if override_conf.Msaa != 0 {
		config.Msaa = override_conf.Msaa
		config.PixelsX *= override_conf.Msaa
		config.PixelsY *= override_conf.Msaa
	}
	if override_conf.MaxBounces != 0 {
		config.MaxBounces = override_conf.MaxBounces
	}
	if override_conf.SaveAsPNG != config.SaveAsPNG {
		config.SaveAsPNG = override_conf.SaveAsPNG
	}

	return config
}

func (c Config) Lx() int {
	return -c.PixelsX / 2
}

func (c Config) Hx() int {
	return c.PixelsX / 2
}

func (c Config) Ly() int {
	return -c.PixelsY / 2
}

func (c Config) Hy() int {
	return c.PixelsY / 2
}

// Vector that comes from the eye and goes through the (i,j) pixel
func (conf Config) Pxy(i, j int) Vector {
	px := Vector{X: conf.ScreenX / float64(conf.PixelsX), Y: 0, Z: 0}
	py := Vector{X: 0, Y: conf.ScreenY / float64(conf.PixelsY), Z: 0}
	d := Vector{X: 0, Y: 0, Z: conf.D}
	return d.Add(px.Dilate(float64(i))).Add(py.Dilate(float64(j)))
}
