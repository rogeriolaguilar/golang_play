package main

import (
	"fmt"
	"math"
)

// Sqrt finds square root of x using the Newton's method
func Sqrt(x float64, precision float64) (z float64) {
	z = 1.0
	for true {
		delta := (z*z - x) / (2 * z)
		if math.Abs(delta) < precision {
			return z
		}
		z -= delta
		fmt.Println(z)
	}
	return
}

func main() {
	number := 1000.0
	precision := 1e-15
	fmt.Printf("The square root of %f is %f, with precision %.20f \n", number, sqrt(number, precision), precision)
}
