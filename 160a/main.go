package main

import "fmt"

func main() {
	var n int
	var coins [101]int
	fmt.Scan(&n)

	var v int
	total := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		coins[v] += 1
		total += v
	}

	num := 0
	left := total/2 + 1
	for i := 100; i > 0; i-- {
		if coins[i] == 0 {
			continue
		}
		if coins[i]*i < left {
			left -= coins[i] * i
			num += coins[i]
		} else {
			for j := 1; j <= coins[i]; j++ {
				left -= i
				num += 1
				if left <= 0 {
					fmt.Print(num)
					return
				}
			}
		}
	}
}
