package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (z float64) {
	z = 1.0
	for true {
		delta := (z*z - x) / (2 * z)
		if math.Abs(delta) < 0.000001 {
			return z
		}
		z -= delta
		fmt.Println(z)
	}
	return
}

func main() {
	fmt.Println(sqrt(20.0))
}
