package main

import "fmt"

var on [105]int

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	u, v := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&u)
		for j := 0; j < u; j++ {
			fmt.Scan(&v)
			on[v] += 1
		}
	}

	for i := 1; i <= m; i++ {
		if on[i] == 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")

}
