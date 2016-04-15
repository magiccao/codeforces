package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	coordinats := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&coordinats[i])
	}

	for i := 0; i < n; i++ {
		min := int64(1e11 + 5)
		left := i - 1
		right := i + 1
		if left >= 0 {
			if coordinats[i]-coordinats[left] < min {
				min = coordinats[i] - coordinats[left]
			}
		}
		if right < n {
			if coordinats[right]-coordinats[i] < min {
				min = coordinats[right] - coordinats[i]
			}
		}

		max := coordinats[i] - coordinats[0]
		if max < coordinats[n-1]-coordinats[i] {
			max = coordinats[n-1] - coordinats[i]
		}

		fmt.Printf("%d %d\n", min, max)
	}
}
