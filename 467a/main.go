package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n)

	num := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		if y-x >= 2 {
			num += 1
		}
	}

	fmt.Print(num)
}
