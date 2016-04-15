package main

import "fmt"

var cost [5]int

func main() {
	fmt.Scan(&cost[1], &cost[2], &cost[3], &cost[4])
	var game string
	fmt.Scan(&game)

	sum := 0
	for _, c := range game {
		id := c - '0'
		sum += cost[id]
	}

	fmt.Println(sum)
}
