package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	vowels := []byte{'a', 'o', 'y', 'e', 'u', 'i'}
	var word string
	fmt.Scan(&word)
	word = strings.ToLower(word)

	for i := 0; i < len(word); i++ {
		if index := bytes.IndexByte(vowels, word[i]); index < 0 {
			fmt.Printf(".%c", word[i])
		}
	}
}
