package render

import (
	"broengine/config"
	"image/color"
)

type Screen struct {
	Pixels [config.PixelsWidth][config.PixelsHeight]color.Color // [r, g, b] (no alpha)
}

// (i ->,  and j ^) centered
// Save the color to the right (i, j) in screen.pixels
func (s *Screen) FillPixel(i int, j int, color color.Color) {
	I := i + config.Hx
	J := (config.PixelsHeight - 1) - (j + config.Hy)
	s.Pixels[I][J] = color
}
