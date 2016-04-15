package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Print(n / 2)
	} else {
		fmt.Print(-1 + n/2*(-1))
	}
}
