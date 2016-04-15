package main

import "fmt"

func main() {
	var a1, a2, a3, b1, b2, b3 int
	var n int
	fmt.Scan(&a1, &a2, &a3, &b1, &b2, &b3, &n)

	cpus := a1 + a2 + a3
	medals := b1 + b2 + b3

	shelves := (cpus+4)/5 + (medals+9)/10
	if shelves <= n {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
