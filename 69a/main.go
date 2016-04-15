package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x, y, z := 0, 0, 0
	var x1, y1, z1 int
	for i := 0; i < n; i++ {
		fmt.Scan(&x1, &y1, &z1)
		x += x1
		y += y1
		z += z1
	}

	if x != 0 || y != 0 || z != 0 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
