package main

import "fmt"

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	m := (n + 1) / 2
	if k <= m {
		fmt.Print(2*(k-1) + 1)
	} else {
		fmt.Print(2 * (k - m))
	}
}
