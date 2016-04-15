package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cashes := 1
	var p1, p2, v1, v2 int
	fmt.Scan(&p1, &p2)

	max := 0
	for i := 1; i < n; i++ {
		fmt.Scan(&v1, &v2)
		if v1 == p1 && v2 == p2 {
			cashes += 1
		} else {
			if cashes > max {
				max = cashes
			}
			p1, p2 = v1, v2
			cashes = 1
		}
	}

	fmt.Println(max)
}
