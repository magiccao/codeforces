package main

import "fmt"

var flowers [200005]int

func main() {
	var n int
	fmt.Scan(&n)

	min := 10000000000
	max := 0
	minn, maxn := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&flowers[i])
		if flowers[i] > max {
			max = flowers[i]
		}

		if flowers[i] < min {
			min = flowers[i]
		}
	}

	diff := max - min
	for i := 0; i < n; i++ {
		if flowers[i] == min {
			minn += 1
		} else if flowers[i] == max {
			maxn += 1
		}
	}

	fmt.Println(diff, int64(minn*maxn))
}
