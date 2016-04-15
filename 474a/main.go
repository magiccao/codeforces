package main

import "fmt"

type Site struct {
	row int
	pos int
}

var keyboards [3][]byte
var sites map[byte]*Site

func init() {
	keyboards[0] = []byte("qwertyuiop")
	keyboards[1] = []byte("asdfghjkl;")
	keyboards[2] = []byte("zxcvbnm,./")

	sites = make(map[byte]*Site)
	for i := 0; i < 3; i++ {
		for j := 0; j < len(keyboards[i]); j++ {
			sites[keyboards[i][j]] = &Site{i, j}
		}
	}
}

func main() {
	var c, word string
	fmt.Scan(&c, &word)

	for i := 0; i < len(word); i++ {
		site := sites[word[i]]
		if c == "R" {
			fmt.Printf("%c", keyboards[site.row][site.pos-1])
		} else {
			fmt.Printf("%c", keyboards[site.row][site.pos+1])
		}
	}
}
