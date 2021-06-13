package render

import (
	"broengine/config"
	"image/color"
)

// Contains the screen pixels (matrice of color)
type Screen struct {
	Pixels [config.PixelsWidth][config.PixelsHeight]color.Color
}

// Fill the pixel with the given color
func (s *Screen) FillPixel(i int, j int, color color.Color) {
	I, J := convertIndexToScreenIndex(i, j, config.PixelsWidth, config.PixelsHeight)
	s.Pixels[I][J] = color
}

// (i right,  and j up)  with origin the center of the screen
// (I right, J down) with origin top left corner with color
func convertIndexToScreenIndex(i int, j int, PixelsWidth int, PixelsHeight int) (int, int) {
	Hx := PixelsWidth / 2
	Hy := PixelsHeight / 2

	I := (PixelsWidth - 1) - (i + Hx)
	J := j + Hy
	return I, J
}
