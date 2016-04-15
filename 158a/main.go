package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &scores[i])
	}

	num := 0
	for _, score := range scores {
		if score != 0 && score >= scores[k-1] {
			num += 1
		}
	}

	fmt.Printf("%d", num)
}
