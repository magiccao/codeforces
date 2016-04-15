package main

import "fmt"

var (
	onRow [15]bool
	onCol [15]bool
	rowN  int
	colN  int
)

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	var line string
	for i := 0; i < r; i++ {
		fmt.Scan(&line)
		for j := 0; j < c; j++ {
			v := line[j]
			if v == 'S' {
				if !onRow[i] {
					onRow[i] = true
					rowN += 1
				}
				if !onCol[j] {
					onCol[j] = true
					colN += 1
				}
			}
		}
	}

	fmt.Println((r-rowN)*c + (c-colN)*rowN)
}
