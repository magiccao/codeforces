package main

import "fmt"

func main() {
	colors := make(map[int]struct{})
	var v int
	for i := 0; i < 4; i++ {
		fmt.Scan(&v)
		colors[v] = struct{}{}
	}

	fmt.Print(4 - len(colors))
}
