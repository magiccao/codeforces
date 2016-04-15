package main

import "fmt"

func main() {
	var n, v int
	fmt.Scan(&n)

	envInd := 0
	oddInd := 0
	find := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&v)
		if v%2 == 0 {
			if envInd != 0 {
				find = 1
			} else {
				envInd = i
			}
		} else {
			oddInd = i
		}
	}

	if find == 0 {
		fmt.Print(envInd)
	} else {
		fmt.Print(oddInd)
	}
}
