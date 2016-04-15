package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	max, min := -1, 1000000
	i, j := 0, 0
	for k := 0; k < n; k++ {
		if heights[k] > max {
			max = heights[k]
			i = k
		}
		if heights[k] <= min {
			min = heights[k]
			j = k
		}
	}

	if i > j {
		fmt.Print(i + n - j - 2)
	} else {
		fmt.Print(i + n - j - 1)
	}
}
