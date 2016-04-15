package main

import "fmt"

func main() {
	var chars [26]int
	var guest, host, pile string
	fmt.Scan(&guest, &host, &pile)

	for i := 0; i < len(guest); i++ {
		chars[guest[i]-'A'] += 1
	}

	for i := 0; i < len(host); i++ {
		chars[host[i]-'A'] += 1
	}

	for i := 0; i < len(pile); i++ {
		if chars[pile[i]-'A'] == 0 {
			fmt.Print("NO")
			return
		} else {
			chars[pile[i]-'A'] -= 1
		}
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] != 0 {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}
