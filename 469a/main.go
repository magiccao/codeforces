package main

import "fmt"

func main() {
	var n, m, v int
	fmt.Scan(&n)

	pass := make([]bool, n+1)
	sum := 0
	for i := 0; i < 2; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&v)
			if !pass[v] {
				pass[v] = true
				sum += v
			}
		}
	}

	if sum == (1+n)*n/2 {
		fmt.Print("I become the guy.")
	} else {
		fmt.Print("Oh, my keyboard!")
	}
}
