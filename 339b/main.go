package main

import "fmt"

func main() {
	var n, m, v int
	fmt.Scan(&n, &m)

	step := 0
	prev := 1
	for i := 0; i < m; i++ {
		fmt.Scan(&v)
		if v >= prev {
			step += (v - prev)
		} else {
			step += (n + v - prev)
		}
		prev = v
	}

	fmt.Print(step)
}
