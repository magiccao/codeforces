package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	for i := 0; i < len(word); i++ {
		if word[i] == 'H' || word[i] == 'Q' || word[i] == '9' {
			fmt.Print("YES")
			return
		}
	}

	fmt.Print("NO")
}
