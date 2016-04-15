package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		l := 0
		if i <= (n+1)/2 {
			l = 2*i - 1
		} else {
			l = 2*(n-i) + 1
		}

		k := (n-l)/2 + 1
		for j := 1; j <= n; j++ {
			if j < k {
				fmt.Print("*")
			} else if j >= (k + l) {
				fmt.Print("*")
			} else {
				fmt.Print("D")
			}
		}
		fmt.Println("")
	}
}
