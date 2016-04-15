package main

import "fmt"

var dp [4005]int

func main() {
	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	for i := 1; i < len(dp); i++ {
		dp[i] = -100000
	}
	for i := a; i <= n; i++ {
		dp[i] = max(dp[i], dp[i-a]+1)
		if i >= b {
			dp[i] = max(dp[i], dp[i-b]+1)
		}
		if i >= c {
			dp[i] = max(dp[i], dp[i-c]+1)
		}
	}

	fmt.Println(dp[n])
}

func max(x, y int) int {
	if x < y {
		x = y
	}

	return x
}
