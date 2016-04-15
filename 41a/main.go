package main

import "fmt"

func main() {
	var ori, rev string
	fmt.Scan(&ori, &rev)

	if len(ori) != len(rev) {
		fmt.Print("NO")
		return
	}
	n := len(ori)

	for i := 0; i < n; i++ {
		if ori[i] != rev[n-i-1] {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}
