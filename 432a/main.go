package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	v := 0
	c := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if v <= 5-k {
			c += 1
		}
	}

	fmt.Println(c / 3)
}
