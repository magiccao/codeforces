package main

import "fmt"

func main() {
	var n int
	var word string
	fmt.Scan(&n, &word)
	flag := word[0]

	stones := 0
	for i := 1; i < len(word); i++ {
		if word[i] == flag {
			stones += 1
		} else {
			flag = word[i]
		}
	}

	fmt.Print(stones)
}
