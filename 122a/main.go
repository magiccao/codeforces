package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	for _, lucky := range []int{4, 7, 44, 47, 74, 77, 444, 447, 474, 744, 747, 774, 777} {
		if m%lucky == 0 {
			fmt.Print("YES")
			return
		}
	}

	for m > 0 {
		d := m % 10
		if d != 4 && d != 7 {
			fmt.Print("NO")
			return
		}
		m /= 10
	}

	fmt.Print("YES")
}
