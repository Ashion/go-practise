package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y * v.Y);
}

func Abs1( v Vertex ) float64 {
	return math.Sqrt(v.X*v.X + v.Y * v.Y);
}
	

func main() {
	x := Vertex{3, 4}
	fmt.Println(x.Abs())
	fmt.Println(Abs1(x))
}
