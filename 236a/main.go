package main

import "fmt"

func main() {
	var word string
	var chars [200]bool
	fmt.Scan(&word)

	num := 0
	for i := 0; i < len(word); i++ {
		if !chars[word[i]] {
			chars[word[i]] = true
			num += 1
		}
	}

	if num%2 == 0 {
		fmt.Print("CHAT WITH HER!")
	} else {
		fmt.Print("IGNORE HIM!")
	}

}
