package main
import "fmt"
import "time"

func sender(s string, c chan string) {
	c <- s;
}

func main() {
	c := make(chan string)
	go sender("hello", c)
	time.Sleep(100 * time.Millisecond)
	go sender("world", c)
	x,y := <- c, <- c
	fmt.Println(x,y)
}

