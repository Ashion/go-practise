package main 

import(
	"fmt"
)

func add( x , y int, z string) int {
	fmt.Println(z);
	return x + y;
}

func main() {
	fmt.Println(add(1,2, "test"));
}