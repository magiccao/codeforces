package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var subjects [3][]int

	v := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&v)
		subjects[v-1] = append(subjects[v-1], i)
	}

	min := len(subjects[0])
	if min > len(subjects[1]) {
		min = len(subjects[1])
	}

	if min > len(subjects[2]) {
		min = len(subjects[2])
	}

	fmt.Println(min)
	for i := 0; i < min; i++ {
		fmt.Printf("%d %d %d\n", subjects[0][i], subjects[1][i], subjects[2][i])
	}
}
