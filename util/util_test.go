package util

import (
	"testing"
)

func TestSameSide(t *testing.T) {
	a := Vector{4, 3, 0}
	b := Vector{0, 0, 0}
	c := Vector{1, 1, 0}
	d := Vector{1, 0, 0}
	if !isSameSide(a, b, c, d) {
		t.Errorf("erreur")
	}

}

func TestNotSameSide(t *testing.T) {
	a := Vector{3, 4, 0}
	b := Vector{0, 0, 0}
	c := Vector{1, 1, 0}
	d := Vector{1, 0, 0}
	if isSameSide(a, b, c, d) {
		t.Errorf("erreur")
	}

}

func TestNormalize(t *testing.T) {
	v := Vector{0, 0, -1}
	if !(v.Normalize() == v) {
		t.Errorf("erreur")
	}
}

func TestOrthogonalize(t *testing.T) {
	v1 := Vector{3, 4, 0}
	v2 := Vector{1, 0, 0}
	v := Vector{0, 4, 0}
	if v1.Orthogonalize(v2) != v {
		t.Errorf("erreur")
	}
}

func TestProdScal(t *testing.T) {
	x := Vector{-5, -5, 0}
	p := Vector{0, 0, 100}
	n := Vector{0, 0, -1}
	if x.Minus(p).ProdScal(n) != 100 {
		t.Error("erreur")
	}
}
