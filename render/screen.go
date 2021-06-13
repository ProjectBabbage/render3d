package render

import (
	"broengine/config"
	"broengine/util"
	"image/color"
)

// Contains the screen pixels (matrice of color)
type Screen struct {
	Pixels [config.PixelsX + 1][config.PixelsY + 1]color.Color
}

// set to black every pixel
func (s *Screen) Init() {
	for I := 0; I <= config.PixelsX; I++ {
		for J := 0; J <= config.PixelsY; J++ {
			s.Pixels[I][J] = color.Black
		}
	}
}

// Fill the pixel with the given color
func (s *Screen) FillPixel(i int, j int, color color.Color) {
	I, J := convertIndexToScreenIndex(i, j, config.PixelsX, config.PixelsY)
	s.Pixels[I][J] = color
}

// (I right, J down) with origin top left corner with color
func convertIndexToScreenIndex(i int, j int, PixelsX int, PixelsY int) (int, int) {
	Hx := PixelsX / 2
	Hy := PixelsY / 2

	I := i + Hx
	J := j + Hy
	return I, J
}

func Render(scene *util.Scene) {
	var screen = new(Screen)
	screen.Init() // set to black every pixel

	for i := config.Lx; i <= config.Hx; i++ {
		for j := config.Ly; j <= config.Hy; j++ {
			ray := util.NewRay(config.Eye, config.Pxy(i, j))
			intensity := Cast(ray, *scene)
			c := color.Gray{uint8(intensity)}
			screen.FillPixel(i, j, c)
		}
	}
	Rendering(screen)
}
