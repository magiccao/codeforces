package main

import (
	"fmt"
	"strings"
)

func main() {
	var words [2]string
	fmt.Scan(&words[0], &words[1])

	words[0] = strings.ToLower(words[0])
	words[1] = strings.ToLower(words[1])
	for i := 0; i < len(words[0]); i++ {
		if words[0][i] > words[1][i] {
			fmt.Print("1")
			return
		} else if words[0][i] < words[1][i] {
			fmt.Print("-1")
			return
		}
	}

	fmt.Print("0")
}
