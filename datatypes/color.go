package datatypes

import (
	"image/color"
)

// private fields to prevent overflow
type Col struct {
	r, g, b, a uint8
}

func NewCol(r, g, b, a uint8) Col {
	return Col{r, g, b, a}
}

func FixedACol(c Col) Col {
	return Col{c.r, c.g, c.b, 255}
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

func adduint(u1, u2 uint8) uint8 {
	return capuint(int(u1) + int(u2))
}

func muluint(u1, u2 uint8) uint8 {
	return capuint(int(u1) * int(u2))
}

func (c1 Col) Equal(c2 Col) bool {
	if c1.r != c2.r || c1.g != c2.g || c1.b != c2.b || c1.a != c2.a {
		return false
	}
	return true
}

func (c1 Col) AddColor(c2 Col) Col {
	r := adduint(c1.r, c2.r)
	g := adduint(c1.g, c2.g)
	b := adduint(c1.b, c2.b)
	a := adduint(c1.a, c2.a)
	return Col{r, g, b, a}
}

func (c1 Col) MulColor(c2 Col) Col {
	r := muluint(c1.r, c2.r)
	g := muluint(c1.g, c2.g)
	b := muluint(c1.b, c2.b)
	a := muluint(c1.a, c2.a)
	return Col{r, g, b, a}
}

func (c Col) DilateColor(x float64) Col {
	mulx := func(u uint8) uint8 {
		return capuint(int(x * float64(u)))
	}
	return Col{mulx(c.r), mulx(c.g), mulx(c.b), mulx(c.a)}
}

func (c Col) RGBA() (r, g, b, a uint32) {
	return color.NRGBA{c.r, c.g, c.b, c.a}.RGBA()
}

func MeanColor(cs []Col) Col {
	var (
		n float64 = 0
		r float64 = 0
		g float64 = 0
		b float64 = 0
		a float64 = 0
	)
	for _, c := range cs {
		n += 1
		r += float64(c.r)
		g += float64(c.g)
		b += float64(c.b)
		a += float64(c.a)
	}
	return NewCol(uint8(r/n), uint8(g/n), uint8(b/n), uint8(a/n))
}
