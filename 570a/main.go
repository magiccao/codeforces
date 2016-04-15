package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var ticket int64
	candidates := make([]int, n)
	for i := 0; i < m; i++ {
		max := int64(0)
		win := 0
		for j := 0; j < n; j++ {
			fmt.Scan(&ticket)
			if ticket > max {
				max = ticket
				win = j
			}
		}

		candidates[win] += 1
	}

	win := 0
	for i := 1; i < n; i++ {
		if candidates[i] > candidates[win] {
			win = i
		}
	}

	fmt.Print(win + 1)
}
