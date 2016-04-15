package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	persist := 0
	for n >= m {
		persist += m
		n = n - m + 1
	}
	persist += n

	fmt.Print(persist)
}
