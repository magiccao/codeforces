package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	count := 0
	for i := 1; i <= n; i++ {
		if x%i == 0 && x/i <= n {
			count += 1
		}
	}

	fmt.Println(count)
}
