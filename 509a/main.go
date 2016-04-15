package main

import "fmt"

func value(i, j int) int {
	if i == 1 || j == 1 {
		return 1
	}

	return value(i-1, j) + value(i, j-1)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Print(value(n, n))
}
