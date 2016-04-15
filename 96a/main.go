package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	cont := 1
	flag := word[0]
	for i := 1; i < len(word); i++ {
		if word[i] == flag {
			cont += 1
		} else {
			flag = word[i]
			if cont >= 7 {
				break
			}
			cont = 1
		}
	}

	if cont >= 7 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
