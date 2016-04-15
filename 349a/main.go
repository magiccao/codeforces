package main

import "fmt"

var kinds [3]int

func main() {
	var n int
	fmt.Scan(&n)

	v := 0
	ok := true
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if ok {
			if i == 0 && v != 25 {
				ok = false
				continue
			}

			switch v / 25 {
			case 1:
				kinds[1] += 1
			case 2:
				if kinds[1] == 0 {
					ok = false
				} else {
					kinds[1] -= 1
					kinds[2] += 1
				}
			case 4:
				if kinds[1] == 0 {
					ok = false
				} else {
					if kinds[2] > 0 {
						kinds[2] -= 1
						kinds[1] -= 1
					} else if kinds[1] >= 3 {
						kinds[1] -= 3
					} else {
						ok = false
					}
				}
			}
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
