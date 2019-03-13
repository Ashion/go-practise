package main

import (
	"fmt"
)

func split(sum int) (x, y int, z string) {
	x = sum * 4/9;
	y = sum - x;
	z = "test string";
	return
}

func main () {
	fmt.Println(split(17));
}
