package main

import "fmt"

func main() {
	var n, dst int
	fmt.Scan(&n, &dst)

	ok := false
	path := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Scan(&path[i])
	}

	door := 1
	for door < n {
		if door+path[door] == dst {
			ok = true
			break
		}
		door = door + path[door]
	}

	if ok {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
