package main

import "fmt"

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m, &a, &b)

	if b/m >= a {
		fmt.Print(a * n)
	} else if a*(n%m) > b {
		fmt.Print((n + m - 1) / m * b)
	} else {
		fmt.Print(b*(n/m) + a*(n%m))
	}
}
