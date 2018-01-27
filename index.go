//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - BEGIN
package main
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -IMPORT
import (
  "fmt"
  "math"
)
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - FUNCTIONS
func add(x, y int) int {
	return x + y
}

func addMixed(x float32, y int) int {
	return int(x) + y
}

func swap(x, y string) (string, string) {
	return y, x
}
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -MAIN
func main() {
  fmt.Println("Hello World, 世界")
  fmt.Printf("Math, square root of 7 is: %g.\n", math.Sqrt(7))
  fmt.Printf("Math, exported name Pi: %g.\n", math.Pi)
  fmt.Printf("Function, result: %d.\n", add(4, 2))
  fmt.Printf("Function, result: %d.\n", addMixed(1.4142, 2))
  a, b := swap("results", "Multiple")
  fmt.Println(a, b)
}
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - END
