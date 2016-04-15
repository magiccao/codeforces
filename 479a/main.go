package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	max := binOp(binOp(a, b), c)
	tmp := binOp(a, binOp(b, c))
	if max > tmp {
		fmt.Print(max)
	} else {
		fmt.Print(tmp)
	}
}

func binOp(x, y int) int {
	if x+y > x*y {
		return x + y
	}

	return x * y
}
