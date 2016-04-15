package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &p[i])
	}

	buy := a[0]
	price := p[0]
	cost := 0
	for i := 1; i < n; i++ {
		if p[i] >= price {
			buy += a[i]
		} else {
			cost += buy * price
			buy = a[i]
			price = p[i]
		}
	}

	cost += buy * price
	fmt.Print(cost)
}
