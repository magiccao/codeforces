package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	r := abs(x2 - x1)
	tmp := abs(y2 - y1)

	if r == 0 {
		r = abs(y2 - y1)
	} else if tmp != 0 && tmp != r {
		fmt.Println(-1)
		return
	}

	if y1 == y2 {
		fmt.Println(x1, y1+r, x2, y1+r)
	} else if x1 == x2 {
		fmt.Println(x1+r, y1, x1+r, y2)
	} else {
		fmt.Println(x1, y2, x2, y1)
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
