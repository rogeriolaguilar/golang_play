package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/rogeriolaguilar/golang_play/morestrings"
)

func sum(a, b int) int {
	return a + b
}

// return 2 values
func swap(a, b string) (string, string) {
	return b, a
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(sum(33, 39))

	s1, s2 := swap("a", "b")
	fmt.Println("swap", s1, s2)

	fmt.Println(split(20))

	// var declaration
	var i, j int = 1, 2
	//implicit type
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
