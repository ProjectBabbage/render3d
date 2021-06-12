package main

import (
	. "broengine/config"
	. "broengine/util"
	"fmt"
)

func main() {
	p1 := Vector{20, 20, 200}
	p2 := Vector{-15, -30, 200}
	p3 := Vector{40, -10, 200}
	n := Vector{0, 0, 1}
	t := Triangle{p1, p2, p3, n}

	for i := Lx; i <= Hx; i++ {
		for j := Ly; j <= Hy; j++ {
			point := t.Intersect(Pxy(i, j))
			fmt.Println(point)
		}
	}
	// render()
}
