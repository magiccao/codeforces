package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := 0
	var word string
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		if word == "++X" || word == "X++" {
			x += 1
		} else if word == "--X" || word == "X--" {
			x -= 1
		}
	}

	fmt.Print(x)
}
