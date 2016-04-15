package main

import "fmt"

func main() {
	var n, h, v int
	fmt.Scan(&n)

	var visit [2][]bool
	var day []int
	visit[0], visit[1] = make([]bool, n+1), make([]bool, n+1)

	for i := 1; i <= n*n; i++ {
		fmt.Scan(&h, &v)
		if !visit[0][h] && !visit[1][v] {
			day = append(day, i)
			visit[0][h] = true
			visit[1][v] = true
		}
	}

	for i := 0; i < len(day); i++ {
		if i != 0 {
			fmt.Print(" ")
		}

		fmt.Print(day[i])
	}
}
