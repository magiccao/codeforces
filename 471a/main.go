package main

import "fmt"

var sticks [10]int

func main() {
	v := 0
	for i := 0; i < 6; i++ {
		fmt.Scan(&v)
		sticks[v] += 1
	}

	legOk := false
	elephant := false
	for i := 0; i < len(sticks); i++ {
		if sticks[i] >= 4 {
			legOk = true
			if sticks[i] == 6 {
				elephant = true
			}
		} else if sticks[i] == 2 {
			elephant = true
		}
	}

	if !legOk {
		fmt.Println("Alien")
	} else if elephant {
		fmt.Println("Elephant")
	} else {
		fmt.Println("Bear")
	}
}
