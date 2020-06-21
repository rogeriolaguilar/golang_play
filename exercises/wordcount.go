package main

import (
	"fmt"
	"strings"
)

// WordCount counts the words occurrences in a map
func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		counter[word]++
	}
	return counter
}

func main() {
	s := "to be or not to be"
	fmt.Println()
	fmt.Println(s)
	result := WordCount(s)
	fmt.Println(result)
}
