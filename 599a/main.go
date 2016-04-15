package main

import "fmt"

func main() {
	var d [3]int
	fmt.Scan(&d[0], &d[1], &d[2])

	id := 0
	sum := d[0]
	for i := 1; i < 3; i++ {
		if d[id] < d[i] {
			id = i
		}
		sum += d[i]
	}

	l := 0
	if 2*d[id] > sum {
		l = 2 * (sum - d[id])
	} else {
		l = sum
	}

	fmt.Println(l)
}
