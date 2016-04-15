package main

import "fmt"

func main() {
	var colors [][2]int
	var n int
	fmt.Scan(&n)

	colors = make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&colors[i][0], &colors[i][1])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if colors[i][0] == colors[j][1] {
				cnt += 1
			}
		}
	}

	fmt.Print(cnt)
}
