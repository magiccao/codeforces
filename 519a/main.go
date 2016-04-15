package main

import "fmt"

func main() {
	var line string

	var players [2]int
	for i := 0; i < 8; i++ {
		fmt.Scan(&line)
		for _, c := range line {
			switch c {
			case 'Q':
				players[0] += 9
			case 'q':
				players[1] += 9
			case 'R':
				players[0] += 5
			case 'r':
				players[1] += 5
			case 'B', 'N':
				players[0] += 3
			case 'b', 'n':
				players[1] += 3
			case 'P':
				players[0] += 1
			case 'p':
				players[1] += 1
			}
		}
	}

	if players[0] > players[1] {
		fmt.Print("White")
	} else if players[0] < players[1] {
		fmt.Print("Black")
	} else {
		fmt.Print("Draw")
	}
}
