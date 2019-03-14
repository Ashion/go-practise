package main
import "fmt"

func main() {
	a := [6] int {1, 2, 3, 4 }
	p := &a;
	fmt.Printf("%v %p %v\n", a, p, *p)
	fmt.Println(len(a), cap(a))
}
