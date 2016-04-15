package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for ; n > 0; count++ {
		n &= (n - 1)
	}

	fmt.Println(count)
}
