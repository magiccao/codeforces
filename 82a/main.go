package main

import (
	"fmt"
	"math"
)

func main() {
	people := []string{"Sheldon", "Leonard", "Penny", "Rajesh", "Howard"}
	var k int
	fmt.Scan(&k)

	if k <= len(people) {
		fmt.Print(people[k-1])
		return
	}

	round := int(math.Ceil(math.Log2(float64(k)/float64(len(people)) + 1)))
	left := k - 5*(int(math.Pow(2, float64(round-1)))-1)
	pos := (left - 1) / int(math.Pow(2, float64(round-1)))

	fmt.Print(people[pos])
}
