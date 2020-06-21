package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib1 := 0
	fib2 := 1
	index := -1
	return func() (fib int) {
		index++
		switch index {
		case 0:
			fib = 0
		case 1:
			fib = 1
		default:
			fib = fib1 + fib2
			fib1 = fib2
			fib2 = fib
		}
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
