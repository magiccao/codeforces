package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	c := []byte{'h', 'e', 'l', 'o'}
	cont := []int{1, 1, 2, 1}
	flag := 0
	hitNum := 0

	for i := 0; i < len(word); i++ {
		if word[i] == c[flag] {
			hitNum += 1
			if hitNum >= cont[flag] {
				hitNum = 0
				flag += 1
			}
		}

		if flag >= len(c) {
			fmt.Print("YES")
			return
		}
	}

	fmt.Print("NO")
}
