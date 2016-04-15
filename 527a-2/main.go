package main

import "fmt"

func order(x, y *int64) {
	if *x < *y {
		*x, *y = *y, *x
	}
}

func main() {
	var a, b, sharps int64
	fmt.Scan(&a, &b)

	for b > 0 {
		sharps += a / b
		a, b = b, a%b
	}

	fmt.Println(sharps)
}
