package main

import "fmt"

func gcd(a, b int) int {
	for {
		v := a % b
		a, b = b, v
		if b == 0 {
			break
		}
	}

	return a
}

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	for {
		gone := gcd(a, n)
		n -= gone
		if n == 0 {
			fmt.Print(0)
			break
		}

		gone = gcd(b, n)
		n -= gone
		if n == 0 {
			fmt.Print(1)
			break
		}
	}
}
