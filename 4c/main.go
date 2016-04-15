package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	m := make(map[string]int)

	str := ""
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		v, ok := m[str]
		if !ok {
			fmt.Println("OK")
			m[str] = 1
		} else {
			fmt.Printf("%s%d\n", str, v)
			m[str] = v + 1
		}
	}
}
