package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scan(&word)

	lower := 0
	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			lower += 1
		}
	}

	if lower >= (len(word)+1)/2 {
		fmt.Print(strings.ToLower(word))
	} else {
		fmt.Print(strings.ToUpper(word))
	}
}
