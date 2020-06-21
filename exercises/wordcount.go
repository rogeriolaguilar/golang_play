package main

import (
	"fmt"
)

// WordCount counts the characters occurrences in a map
func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	for _, char := range s {
		counter[string(char)]++
	}
	return counter
}

func main() {
	s := "amanha e hoje"
	fmt.Println(s)
	result := WordCount(s)
	fmt.Println(result)
	//wc.Test(WordCount)
}
