package main

import "fmt"

func main() {
	var n int
	var sum [3]int64

	fmt.Scan(&n)

	var v int64
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		sum[0] += v
	}

	for i := 0; i < n-1; i++ {
		fmt.Scan(&v)
		sum[1] += v
	}

	for i := 0; i < n-2; i++ {
		fmt.Scan(&v)
		sum[2] += v
	}

	fmt.Printf("%d\n%d\n", sum[0]-sum[1], sum[1]-sum[2])
}
