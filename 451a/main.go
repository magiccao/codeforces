package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n > m {
		n, m = m, n
	}

	if n%2 != 0 {
		fmt.Print("Akshat")
	} else {
		fmt.Print("Malvika")
	}
}
