package datatypes

import "image/color"

type Col struct {
	R, G, B int
}

func IsoCol(channel int) Col {
	return Col{channel, channel, channel}
}

func RedCol(channel int) Col {
	return Col{channel, 0, 0}
}

func GreenCol(channel int) Col {
	return Col{0, channel, 0}
}

func BlueCol(channel int) Col {
	return Col{0, 0, channel}
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

// proportionally dilate c in regards to c2
func (c Col) PropDilateColor(c2 Col) Col {
	propDilate := func(channel1, channel2 int) int {
		// ex:
		// channel1 = 155
		// channel2 = 255 (the maximum)
		// cap(channel1, channel2) will return 155
		return int(float64(channel1) * (float64(channel2) / 255))
	}
	return Col{
		propDilate(c.R, c2.R),
		propDilate(c.G, c2.G),
		propDilate(c.B, c2.B),
	}
}

func (c Col) RGBA() (r, g, b, a uint32) {
	return color.NRGBA{capuint(c.R), capuint(c.G), capuint(c.B), 255}.RGBA()
}
