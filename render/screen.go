package render

import (
	. "broengine/config"
	"image/color"
)

// Contains the screen pixels (matrice of color)
type Screen struct {
	// sizes of the screen in pixel
	PixelsX, PixelsY int
	// the pixels colors
	Pixels [][]color.Gray
}

// Init the screen sizes and set to black every pixel
// [PixelsX + 1][PixelsY + 1]color.Gray
func (s *Screen) Init(conf Config) {
	s.PixelsX, s.PixelsY = conf.PixelsX, conf.PixelsY
	s.Pixels = make([][]color.Gray, s.PixelsX+1)
	for I := 0; I <= conf.PixelsX; I++ {
		for J := 0; J <= conf.PixelsY; J++ {
			s.Pixels[I] = append(s.Pixels[I], color.Gray{0})
		}
	}
}

// Fill the pixel with the given intensity
func (s *Screen) FillPixel(i int, j int, color color.Gray) {
	I, J := convertIndexToScreenIndex(i, j, s.PixelsX, s.PixelsY)
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
