package main

import "fmt"

func main() {
	var n, t int
	var line string
	fmt.Scan(&n, &t, &line)

	queue := []byte(line)
	for i := 0; i < t; i++ {
		for j := 0; j < len(queue)-1; j++ {
			if queue[j] == 'B' && queue[j+1] == 'G' {
				queue[j], queue[j+1] = queue[j+1], queue[j]
				j++
			}
		}
	}

	fmt.Print(string(queue))

}
