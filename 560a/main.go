package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := 0
	ok := false
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if v == 1 {
			ok = true
		}
	}

	if ok {
		fmt.Println(-1)
	} else {
		fmt.Println(1)
	}

}
