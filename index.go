//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - BEGIN
package main
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -IMPORT
import (
  "fmt"
  "math"
)
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - FUNCTIONS
func add(x int, y int) int {
	return x + y
}
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -MAIN
func main() {
  fmt.Println("Hello World, 世界")
  fmt.Printf("Math, square root of 7 is: %g.\n", math.Sqrt(7))
  fmt.Printf("Math, exported name Pi: %g.\n", math.Pi)
}
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - END
