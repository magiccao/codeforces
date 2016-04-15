package main

import "fmt"

func main() {
	var chars [26]int
	var n int
	var word string
	fmt.Scan(&n, &word)

	if n < 26 {
		fmt.Print("NO")
		return
	}

	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'A' && c <= 'Z' {
			chars[c-'A'] += 1
		} else if c >= 'a' && c <= 'z' {
			chars[c-'a'] += 1
		}
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] == 0 {
			fmt.Print("NO")
			return
		}
	}

	fmt.Print("YES")
}
