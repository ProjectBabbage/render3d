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
