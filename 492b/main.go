package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	lanterns := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lanterns[i])
	}

	sort.Ints(lanterns)
	radious := float64(lanterns[0])
	r := float64(l - lanterns[n-1])
	radious = math.Max(radious, r)

	for i := 0; i < n-1; i++ {
		r = float64(lanterns[i+1]-lanterns[i]) / 2
		radious = math.Max(radious, r)
	}

	fmt.Printf("%.10f", radious)
}
