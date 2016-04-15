package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	v := 0
	for i := 0; i < t; i++ {
		fmt.Scan(&v)
		if 360%(180-v) != 0 {
			fmt.Println("NO")
		} else {
			n := 360 / (180 - v)
			if n < 3 {
				fmt.Println("NO")
			} else {
				fmt.Println("YES")
			}
		}
	}
}
