package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	digits := 0
	m := t
	for m > 0 {
		digits += 1
		m /= 10
	}

	if digits > n {
		fmt.Print(-1)
		return
	}

	fmt.Print(t)
	n -= digits

	for i := 0; i < n; i++ {
		fmt.Print(0)
	}
}
