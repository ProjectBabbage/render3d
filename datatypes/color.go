package datatypes

import "image/color"

type Col struct {
	R, G, B int
}

func capuint(i int) uint8 {
	switch {
	case i <= 0:
		return 0
	case i > 255:
		return 255
	default:
		return uint8(i)
	}
}

func (c1 Col) AddColor(c2 Col) Col {
	r := c1.R + c2.R
	g := c1.G + c2.G
	b := c1.B + c2.B
	return Col{r, g, b}
}

func (c1 Col) MulColor(c2 Col) Col {
	r := c1.R * c2.R
	g := c1.G * c2.G
	b := c1.B * c2.B
	return Col{r, g, b}
}

func (c Col) DilateColor(x float64) Col {
	mulx := func(n int) int {
		return int(x * float64(n))
	}
	return Col{mulx(c.R), mulx(c.G), mulx(c.B)}
}

func (c Col) RGBA() (r, g, b, a uint32) {
	return color.NRGBA{capuint(c.R), capuint(c.G), capuint(c.B), 255}.RGBA()
}
