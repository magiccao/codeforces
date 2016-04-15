package main

import "fmt"

var rating [2001][]int
var students [2000]int

func main() {
	var n, v int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		rating[v] = append(rating[v], i)
	}

	pos := 1
	for i := 2000; i > 0; i-- {
		if len(rating[i]) == 0 {
			continue
		}

		for _, r := range rating[i] {
			students[r] = pos
		}
		pos += len(rating[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", students[i])
	}
}
