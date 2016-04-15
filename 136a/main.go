package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var to int
	giftFroms := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&to)
		giftFroms[to] = i
	}

	for i := 1; i < len(giftFroms); i++ {
		if i != 1 {
			fmt.Print(" ")
		}
		fmt.Print(giftFroms[i])
	}
}
