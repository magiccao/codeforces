package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)
	c := word[0]
	if c >= 'a' && c <= 'z' {
		c -= 32
	}
	fmt.Printf("%c%s", c, word[1:])
}
