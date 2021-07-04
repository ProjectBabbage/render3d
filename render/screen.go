package render

import (
	"image/color"
)

// Contains the screen pixels (matrice of color)
type Screen struct {
	// sizes of the screen in pixel
	PixelsX, PixelsY int
	// the pixels colors
	Pixels [][]color.Gray
}

// Init the screen sizes and set to black every pixel.
// [PixelsX + 1][PixelsY + 1]color.Gray
func NewScreen(Px, Py int) Screen {
	s := Screen{}
	s.PixelsX, s.PixelsY = Px, Py
	s.Pixels = make([][]color.Gray, s.PixelsX+1)
	for I := 0; I <= s.PixelsX; I++ {
		for J := 0; J <= s.PixelsY; J++ {
			s.Pixels[I] = append(s.Pixels[I], color.Gray{0})
		}
	}
	return s
}

// Fill the pixel with the given intensity
// i, j are indexes from "natural" referential:
// x right, y up, screen centered on (0, 0)
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

// Returns a screen of size (S.PixelsX/msaa, S.PixelsY/msaa).
// Each pixel intensity is the mean of the pixels intensities around it
func (S *Screen) MeanScreen(msaa int) Screen {
	newS := NewScreen(S.PixelsX/msaa, S.PixelsY/msaa)

	for I := 0; I < newS.PixelsX; I++ {
		for J := 0; J < newS.PixelsY; J++ {
			// compute the mean intensity
			var intensity int = 0
			for i := 0; i < msaa; i++ {
				for j := 0; j < msaa; j++ {
					intensity += int(S.Pixels[I*msaa+i][J*msaa+j].Y)
				}
			}
			intensity /= (msaa * msaa)
			//

			newS.Pixels[I][J] = color.Gray{uint8(intensity)}
		}
	}
	return newS
}
