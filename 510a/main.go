package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i%2 != 0 {
				fmt.Print("#")
			} else if j == 1 && i/2%2 == 0 {
				fmt.Print("#")
			} else if j == m && i/2%2 != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
