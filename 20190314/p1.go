package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Printf("p type : %T, value : %v, address : %p\n", p, *p, p)
	//fmt.Println(p)
	p = &j
	fmt.Println(p)
}
