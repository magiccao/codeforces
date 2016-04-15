package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	var nums [4]int
	for i := 0; i < len(word); i += 2 {
		nums[word[i]-'0'] += 1
	}

	ok := false
	for i := 1; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}

		if ok {
			fmt.Print("+")
		} else {
			ok = true
		}

		for j := 0; j < nums[i]; j++ {
			if j != 0 {
				fmt.Print("+")
			}
			fmt.Printf("%c", i+'0')
		}
	}
}
