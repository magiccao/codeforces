package main

import "fmt"

func sumOne(x int) int {
	sum := 0
	for x > 0 {
		sum += 1
		x &= (x - 1)
	}

	return sum
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	players := make([]int, m+1)

	friends := 0
	for i := 0; i <= m; i++ {
		fmt.Scan(&players[i])
	}

	tom := players[m]
	for i := 0; i < m; i++ {
		v := tom ^ players[i]
		if sumOne(v) <= k {
			friends += 1
		}
	}

	fmt.Println(friends)
}
