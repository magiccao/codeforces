package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	cubes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cubes[i])
	}

	sort.Ints(cubes)
	for i := 0; i < len(cubes); i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", cubes[i])
	}
}
