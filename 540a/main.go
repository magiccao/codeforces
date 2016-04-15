package main

import "fmt"

func main() {
	var n, diff int
	var pos, dst string
	fmt.Scan(&n, &pos, &dst)

	sum := 0
	for i := 0; i < len(pos); i++ {
		if pos[i] > dst[i] {
			diff = int(pos[i] - dst[i])
		} else {
			diff = int(dst[i] - pos[i])
		}

		if diff > 5 {
			diff = 10 - diff
		}
		sum += diff
	}

	fmt.Print(sum)
}
