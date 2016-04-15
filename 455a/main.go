package main

import (
	"fmt"
	"sort"
)

func points(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}

	max := 0
	b := make([]int, len(a))
	for i := 0; i < len(a); {
		pos := 0
		cursor := a[i]
		for j := 0; j < len(a); j++ {
			if j == i {
				continue
			}
			if a[j] != cursor-1 && a[j] != cursor+1 {
				if j != pos {
					a[j], a[pos] = a[pos], a[j]
				}
				b[pos] = j
				pos += 1
			}
		}
		ps := cursor + points(a[:pos])
		if max < ps {
			max = ps
		}

		// recover spot
		for j := 0; j < pos; j++ {
			a[j], a[b[j]] = a[b[j]], a[j]
		}

		for i += 1; i < len(a); i++ {
			if a[i] > cursor {
				break
			}
		}
	}

	return max
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	fmt.Print(points(a))
}
