package main

import (
	"fmt"
	"sort"
)

var month [12]int

func main() {
	var k int
	fmt.Scan(&k)

	for i := 0; i < 12; i++ {
		fmt.Scan(&month[i])
	}

	if k == 0 {
		fmt.Println(0)
		return
	}

	sort.Ints(month[:])
	i := 11
	for ; i >= 0; i-- {
		k -= month[i]
		if k <= 0 {
			break
		}
	}

	if i < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(12 - i)
	}
}
