package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	caps := 0
	left := 0
	var u, v int
	for i := 0; i < n; i++ {
		fmt.Scan(&u, &v)
		left += (v - u)
		if left > caps {
			caps = left
		}
	}

	fmt.Print(caps)
}
