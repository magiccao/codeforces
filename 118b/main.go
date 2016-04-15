package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < 2*n+1; i++ {
		l := i
		if i > n {
			l = 2*n - i
		}

		for j := 0; j < 2*n+1; j++ {
			if j < n-l || j > n+l {
				if j < n-l {
					fmt.Print(" ")
				} else {
					continue
				}
			} else {
				k := l - dis(j, n)
				fmt.Print(l - dis(j, n))
				if j >= n && k == 0 {
					continue
				}
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func dis(x, y int) int {
	d := x - y
	if d < 0 {
		d = -1 * d
	}

	return d
}
