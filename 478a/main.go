package main

import "fmt"

func main() {
	var v int
	sum := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&v)
		sum += v
	}

	if sum%5 != 0 || sum == 0 {
		fmt.Print(-1)
	} else {
		fmt.Print(sum / 5)
	}
}
