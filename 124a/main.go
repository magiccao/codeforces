package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	pos := a + 1
	if pos < n-b {
		pos = n - b
	}

	fmt.Println(n - pos + 1)
}
