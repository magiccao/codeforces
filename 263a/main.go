package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 0, 0

	var v int
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Scan(&v)
			if v == 1 {
				x, y = i, j
			}
		}
	}

	fmt.Print(int(math.Abs(float64(x-3)) + math.Abs(float64(y-3))))
}
