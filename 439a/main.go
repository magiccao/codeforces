package main

import "fmt"

func main() {
	var n, d, v, jokes int
	fmt.Scan(&n, &d)

	cost := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		cost += v
	}

	cost += (n - 1) * 10
	jokes = (n - 1) * 2
	delta := d - cost
	if delta < 0 {
		fmt.Println(-1)
		return
	}
	jokes += delta / 5

	fmt.Println(jokes)
}
