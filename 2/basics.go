package main

import "fmt"

func main() {

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type %T\n", x, x)
	fmt.Printf("y=%v, type %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("mean: %v, type of %T\n", mean, mean)
	
}