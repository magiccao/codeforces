package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Printf("%d ", b)
		fmt.Print((a - b) / 2)
	} else {
		fmt.Printf("%d ", a)
		fmt.Print((b - a) / 2)
	}
}
