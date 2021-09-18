package datatypes

import "image/color"

type Col struct {
	R, G, B float64
}

func IsoCol(channel float64) Col {
	return Col{channel, channel, channel}
}

func RedCol(channel float64) Col {
	return Col{channel, 0, 0}
}

func GreenCol(channel float64) Col {
	return Col{0, channel, 0}
}

func BlueCol(channel float64) Col {
	return Col{0, 0, channel}
}

func capfloat64(i float64) float64 {
	switch {
	case i <= 0:
		return 0.
	case i > 255:
		return 255.
	default:
		return i
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
	mulx := func(n float64) float64 {
		return x * float64(n)
	}
	return Col{mulx(c.R), mulx(c.G), mulx(c.B)}
}

func (c Col) DilateColorByChannels(c1 Col) Col {
	return Col{c.R * c1.R, c.G * c1.G, c.B * c1.B}
}

func (c Col) RGBA() (r, g, b, a uint32) {
	// This function call by SDL needs float32 as output
	r = uint32(capfloat64(c.R))
	r |= r << 8
	g = uint32(capfloat64(c.G))
	g |= g << 8
	b = uint32(capfloat64(c.B))
	b |= b << 8
	a = 255
	return
}

func (c Col) ColorRGBA() color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}
