package main

import "fmt"
import "math/cmplx"

var i, j int = 1, 2
var y int = 1<<32 - 1
var x complex128 = cmplx.Sqrt(-5 + 12i);

func main() {
	var c, python, java = true, false, "java"
	k := "xxxx"
	fmt.Println(i, j, c, python, java, k);
	fmt.Printf("y Type : %T, Value : %v\n", y, y); 
	fmt.Printf("x Type : %T, Value : %v\n", x, x); 
}
