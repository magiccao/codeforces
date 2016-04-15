package main

import (
	"fmt"
	"math"
)

func main() {
	var m, n, a int
	fmt.Scanf("%d %d %d", &m, &n, &a)
	if m > n {
		m, n = n, m
	}

	x := math.Ceil(float64(m) / float64(a))
	y := math.Ceil(float64(n) / float64(a))
	fmt.Printf("%d", int64(x*y))
}
