package main

import "fmt"

func main() {
	var digit int64
	fmt.Scan(&digit)

	num := 0
	for digit > 0 {
		v := digit % 10
		if v == 4 || v == 7 {
			num += 1
		}
		digit /= 10
	}

	if num == 4 || num == 7 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
