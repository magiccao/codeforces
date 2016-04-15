package main

import "fmt"

var sign [55]bool

func main() {
	var n, u, v, k int
	fmt.Scan(&n, &u)

	var on []int
	for i := 1; i <= n; i++ {
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			fmt.Scan(&v)
			if u > v && !sign[i] {
				sign[i] = true
				on = append(on, i)
			}
		}
	}

	fmt.Println(len(on))
	for i := 0; i < len(on); i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(on[i])
	}
}
