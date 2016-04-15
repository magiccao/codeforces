package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	puzzles := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&puzzles[i])
	}

	sort.Ints(puzzles)
	min := puzzles[len(puzzles)-1]
	for i := 0; i <= len(puzzles)-n; i++ {
		diff := puzzles[i+n-1] - puzzles[i]
		if diff < min {
			min = diff
		}
	}

	fmt.Print(min)
}
