package main

import "fmt"

func main () {
	var pow = []int {1, 2, 3, 4, 5, 6}
	for i,v := range pow { 
		fmt.Printf("index : %d, value : %d\n", i ,v )
	}
}

