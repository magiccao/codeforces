package main

import "fmt"

func main() {
	var first, second string
	fmt.Scan(&first, &second)

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
