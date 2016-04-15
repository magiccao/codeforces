package main

import "fmt"

func main() {
	var n, v, min, max int
	fmt.Scan(&n)

	amazing := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if i == 0 {
			min, max = v, v
		} else {
			if v > max {
				amazing += 1
				max = v
			}
			if v < min {
				amazing += 1
				min = v
			}
		}
	}

	fmt.Print(amazing)
}
