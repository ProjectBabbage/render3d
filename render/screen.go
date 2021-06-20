package render

import (
	"broengine/config"
	"image/color"
)

// Contains the screen pixels (matrice of color)
type Screen struct {
	Pixels [config.PixelsX + 1][config.PixelsY + 1]color.Gray
}

// set to black every pixel
func (s *Screen) Init() {
	for I := 0; I <= config.PixelsX; I++ {
		for J := 0; J <= config.PixelsY; J++ {
			s.Pixels[I][J] = color.Gray{0} // That's black
		}
	}
}

// Fill the pixel with the given intensity
func (s *Screen) FillPixel(i int, j int, color color.Gray) {
	I, J := convertIndexToScreenIndex(i, j, config.PixelsX, config.PixelsY)
	s.Pixels[I][J] = color
}

// (I right, J down) with origin top left corner
func convertIndexToScreenIndex(i int, j int, PixelsX int, PixelsY int) (int, int) {
	Hx := PixelsX / 2
	Hy := PixelsY / 2

	I := i + Hx
	J := j + Hy

	return I, J
}

// intensity of the pixel is a float, sometimes > 255
func ConvertIntensityToGrayScale(i float64) color.Gray {
	if i > 255 {
		return color.Gray{255} // white
	}
	if i < 0 {
		return color.Gray{0} // black
	}
	return color.Gray{byte(i)} // byte is an alias for uint8
}
