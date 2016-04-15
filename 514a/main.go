package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '5' {
			if c == '9' && i == 0 {
				fmt.Printf("%c", c)
			} else {
				fmt.Printf("%c", '9'-c+'0')
			}
		} else {
			fmt.Printf("%c", c)
		}
	}
}
