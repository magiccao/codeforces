package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	flag := 0
	max := 1

	var cur int32
	var prev int32
	fmt.Scan(&prev)

	for i := 1; i < n; i++ {
		fmt.Scan(&cur)
		if cur < prev {
			if max < i-flag {
				max = i - flag
			}
			flag = i
		}
		prev = cur
	}

	if max < (n - flag) {
		max = n - flag
	}
	fmt.Print(max)
}
