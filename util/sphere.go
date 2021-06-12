package util

import (
	"fmt"
)

type Sphere struct {
	C Vector
	R float64
}

func (s Sphere) Print() {
	fmt.Println(s)
}
