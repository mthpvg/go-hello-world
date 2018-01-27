//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - BEGIN
package main

//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -IMPORT
import (
	"fmt"
	"math"
)

var packageVar bool

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

func nakedReturns(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -MAIN
func main() {
	var functionVar bool
	var golang, java = true, "no!"

	fmt.Println("Hello World, 世界")
	fmt.Printf("Math, square root of 7 is: %g.\n", math.Sqrt(7))
	fmt.Printf("Math, exported name Pi: %g.\n", math.Pi)
	fmt.Printf("Function, result: %d.\n", add(4, 2))
	fmt.Printf("Function, result: %d.\n", addMixed(1.4142, 2))
	a, b := swap("results", "Multiple")
	fmt.Println(a, b)
	c, d := nakedReturns(17)
	fmt.Printf("Naked return: %d %d\n", c, d)
	fmt.Printf("Variable declaration: %t, %t\n", packageVar, functionVar)
	fmt.Printf("Variable initializers: %t, %s\n", golang, java)
}

//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - END
