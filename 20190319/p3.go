package main

import (
	"fmt"
	"math"
)

type Itf interface {
	M() float64
}

func describe( i Itf ) {
	fmt.Printf("( %v, %T )\n", i, i)
}

type MyFloat float64

func (f MyFloat) M() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	var i Itf
	i = f
	describe(i)
	fmt.Println(i.M())
}

