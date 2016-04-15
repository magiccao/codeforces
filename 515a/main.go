package main

import "fmt"

func main() {
	var a, b, s int64
	fmt.Scan(&a, &b, &s)

	delta := make([]int64, 4)
	delta[0] = s + a + b
	delta[1] = s - a - b
	delta[2] = s + a - b
	delta[3] = s - a + b
	for i := 0; i < 4; i++ {
		if delta[i] < 0 || delta[i]%2 != 0 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
