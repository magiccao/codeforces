package main

import "fmt"

func main() {
	var n int
	var key string
	fmt.Scan(&n)
	fmt.Scan(&key)
	fmt.Scan(&key)

	sum := 0
	if key == "week" {
		sum += 363 / 7
		if n <= 363%7 {
			sum += 1
		}
		if n >= 5 {
			sum += 1
		}
	} else {
		if n <= 29 {
			sum = 12
		} else if n == 30 {
			sum = 11
		} else {
			sum = 7
		}
	}

	fmt.Println(sum)
}
