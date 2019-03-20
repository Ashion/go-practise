package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f);
}

type Vertex struct {
	X, Y float64
}

func (v * Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
//func (v * Vertex) Abs1() float64 {
//	return math.Sqrt(v.X * v.X + v.Y * v.Y)
//}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	x := Vertex { 3, 4 };
	fmt.Println(x.Abs())
	var y * Vertex
	y = &x;
	fmt.Println(y.Abs())
	
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	a = f
	fmt.Println("a=f,a.abs()", a.Abs())
	//a = x
	//fmt.Println("a=x,a.abs()", a.Abs())
	a = y
	fmt.Println("a=y,a.abs()", a.Abs())
}
	



