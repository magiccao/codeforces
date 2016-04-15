package main

import "fmt"

func distinct(year int) bool {
	var on [10]bool
	for year > 0 {
		v := year % 10
		if on[v] {
			return false
		}
		on[v] = true
		year /= 10
	}

	return true
}

func main() {
	var year int
	fmt.Scan(&year)

	for {
		year += 1
		if distinct(year) {
			fmt.Print(year)
			return
		}
	}
}
