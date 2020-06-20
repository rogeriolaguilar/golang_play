package main

import (
	"fmt"
)

var currentX = 1
var currentY = 1

func printSteps(steps int) {
	for i := 0; i < steps; i++ {
		fmt.Println("Y")
	}
}

func walk(stepX, stepY int) (x, y int) {
	x = currentX + stepX
	y = currentX + stepY
	printSteps(stepX + stepY)
	return
}

func main() {
	stepX := 2
	stepY := 4

	currentX, currentY = walk(stepX, stepY)

	fmt.Println(currentX, currentY)
}
