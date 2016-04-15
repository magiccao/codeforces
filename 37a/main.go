package main

import "fmt"

var height [1001]int

func main() {
	var n int
	fmt.Scan(&n)

	v := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		height[v] += 1
	}

	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}

	num := n
	for i := 0; i < len(height); i++ {
		if height[i] != 0 {
			num = num - height[i] + 1
		}
	}

	fmt.Println(max, num)
}
