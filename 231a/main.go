package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	solution := 0
	for i := 0; i < n; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		if u+v+w >= 2 {
			solution += 1
		}
	}

	fmt.Print(solution)
}
