package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	convert := true
	for i := 1; i < len(word); i++ {
		if word[i] > 'Z' || word[i] < 'A' {
			convert = false
			break
		}
	}

	if !convert {
		fmt.Print(word)
	} else {
		for i := 0; i < len(word); i++ {
			if word[i] >= 'A' && word[i] <= 'Z' {
				fmt.Printf("%c", word[i]+32)
			} else {
				fmt.Printf("%c", word[i]-32)
			}
		}
	}
}
