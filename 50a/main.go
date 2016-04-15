package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	fmt.Printf("%d", n/2*m+n%2*m/2)
}
