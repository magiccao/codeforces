package main

import "fmt"

var a [100]int

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				if a[k] == 0 {
					sum += 1
				}
			}
			for k := 0; k < i; k++ {
				if a[k] == 1 {
					sum += 1
				}
			}
			for k := j + 1; k < n; k++ {
				if a[k] == 1 {
					sum += 1
				}
			}

			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
