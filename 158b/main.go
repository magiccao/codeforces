package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var groups [5]int
	var m int
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		groups[m] += 1
	}

	cars := groups[4] + groups[3]
	groups[1] -= groups[3]

	cars += groups[2] / 2
	if groups[2]%2 == 1 {
		cars += 1
		groups[1] -= 2
	}

	if groups[1] > 0 {
		cars += (groups[1] + 3) / 4
	}

	fmt.Print(cars)
}
