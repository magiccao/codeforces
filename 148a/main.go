package main

import "fmt"

func main() {
	var k, l, m, n, d int
	fmt.Scan(&k, &l, &m, &n, &d)

	concerns := []int{k, l, m, n}
	concerns = append(concerns, lcm(k, l))
	concerns = append(concerns, lcm(k, m))
	concerns = append(concerns, lcm(k, n))
	concerns = append(concerns, lcm(l, m))
	concerns = append(concerns, lcm(l, n))
	concerns = append(concerns, lcm(m, n))

	concerns = append(concerns, lcm(lcm(k, l), m))
	concerns = append(concerns, lcm(lcm(k, l), n))
	concerns = append(concerns, lcm(lcm(k, m), n))
	concerns = append(concerns, lcm(lcm(l, m), n))

	concerns = append(concerns, lcm(lcm(k, l), lcm(m, n)))

	num := 0
	for _, concern := range concerns[:4] {
		num += d / concern
	}
	for _, concern := range concerns[4:10] {
		num -= d / concern
	}
	for _, concern := range concerns[10:14] {
		num += d / concern
	}
	num -= d / concerns[14]

	fmt.Print(num)
}

func lcm(x, y int) int {
	a, b := x, y
	for {
		r := a % b
		a, b = b, r
		if b == 0 {
			break
		}
	}
	return x * y / a
}
