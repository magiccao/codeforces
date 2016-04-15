package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	left := a
	sum := a
	for left >= b {
		newmake := left / b
		sum += newmake
		left = newmake + left%b
	}

	fmt.Print(sum)
}
