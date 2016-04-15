package main

import "fmt"

const (
	MAX_CHARS = 10
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}

	for i, word := range words {
		length := len(word)
		if length > MAX_CHARS {
			fmt.Printf(fmt.Sprintf("%s%d%s", word[:1], length-2, word[length-1:]))
		} else {
			fmt.Printf(word)
		}

		if i != len(words)-1 {
			fmt.Println("")
		}
	}
}
