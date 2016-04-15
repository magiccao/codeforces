package main

import (
	"fmt"
	"math"
)

func prime(n int) bool {
	if n <= 2 {
		return true
	}

	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n, x, y int
	fmt.Scan(&n)
	if n%2 == 0 {
		x = n / 2
		y = x
	} else {
		x = n / 2
		y = x + 1
	}

	for {
		if !prime(x) && !prime(y) {
			fmt.Printf("%d %d", x, y)
			return
		}
		x -= 1
		y += 1
	}

}
