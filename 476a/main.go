package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n < m {
		fmt.Print(-1)
		return
	}

	step := n / 2
	left := n % 2
	exchange := false
	for (step+left)%m != 0 {
		if exchange {
			step -= 1
			exchange = false
		}
		left += 1
	}

	fmt.Print(step + left)
}
