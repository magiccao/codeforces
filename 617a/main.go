package main

import "fmt"

func main() {
	var pos int
	fmt.Scan(&pos)

	step := (pos + 4) / 5
	fmt.Print(step)
}
