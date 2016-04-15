package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Scan(&line)

	words := strings.Split(line, "WUB")
	hasPrev := false
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if hasPrev {
			fmt.Print(" ")
		} else {
			hasPrev = true
		}
		fmt.Print(word)
	}
}
