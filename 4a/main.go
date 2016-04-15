package main

import "fmt"

func main() {
	var w int
	fmt.Scanf("%d", &w)

	if w != 2 && w%2 == 0 {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
