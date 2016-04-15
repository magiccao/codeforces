package main

import (
	"fmt"
	"sort"
)

type fight struct {
	dragon int
	bonus  int
}

type fights []*fight

func (f fights) Len() int {
	return len(f)
}

func (f fights) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f fights) Less(i, j int) bool {
	return f[i].dragon < f[j].dragon
}

func main() {
	var s, n int
	fmt.Scan(&s, &n)

	var fs fights
	var u, v int
	for i := 0; i < n; i++ {
		fmt.Scan(&u, &v)
		fs = append(fs, &fight{u, v})
	}

	sort.Sort(fs)

	for _, f := range fs {
		if s <= f.dragon {
			fmt.Print("NO")
			return
		}

		s += f.bonus
	}

	fmt.Print("YES")
}
