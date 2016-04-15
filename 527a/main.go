package main

import "fmt"

func main() {
	var n, v int
	fmt.Scan(&n)

	untreated := 0
	police := 0
	crime := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if v >= 0 {
			police += v
		} else {
			if police > 0 {
				police -= 1
			} else {
				crime += 1
				if crime > untreated {
					untreated = crime
				}
			}
		}
	}

	fmt.Print(untreated)
}
