package main

import "fmt"

func main() {
	var n1, n2, k1, k2 int
	fmt.Scan(&n1, &n2, &k1, &k2)

	if n1 > n2 {
		fmt.Print("First")
	} else {
		fmt.Print("Second")
	}
}
