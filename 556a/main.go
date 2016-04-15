package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n int
	var line string
	var c, t byte
	fmt.Scan(&n)
	fmt.Scan(&line)

	lst := list.New()
	for i := 0; i < len(line); i++ {
		c = line[i]
		if lst.Len() == 0 {
			lst.PushBack(c)
		} else {
			t = lst.Back().Value.(byte)
			if t != c {
				lst.Remove(lst.Back())
			} else {
				lst.PushBack(c)
			}
		}
	}

	fmt.Println(lst.Len())
}
