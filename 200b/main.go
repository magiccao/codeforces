package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum, v := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		sum += v
	}

	fmt.Printf("%.12f\n", float64(sum)/float64(n))
}
