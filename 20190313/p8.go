package main
import "fmt"
import "math"

func test(x float64) string {
	if x < 0 {
		return test(-x) + "i";
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(test(9), test(-9));
}

