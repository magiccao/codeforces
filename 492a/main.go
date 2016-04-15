package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	h := int(math.Sqrt(float64(2 * n)))
	sum := 0
	for i := 1; i <= h; i++ {
		sum += (i + 1) * i / 2
	}

	for sum > n {
		sum -= (1 + h) * h / 2
		h -= 1
	}

	fmt.Print(h)
}
