package main

import (
	"fmt"
	"sort"
)

var group [300005]int

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&group[i])
	}

	if n == 1 {
		fmt.Println(group[0])
		return
	}

	b := group[:n]
	sort.Ints(b)
	var score int64
	for i := 0; i < n-2; i++ {
		score += int64((i + 2) * b[i])
	}

	score += int64(n * (b[n-2] + b[n-1]))
	fmt.Println(score)
}
