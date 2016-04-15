package main

import "fmt"

func main() {
	var account int64
	fmt.Scan(&account)

	if account >= 0 || account/10 == 0 {
		fmt.Print(account)
	} else {
		account *= -1
		last := account % 10
		account /= 10
		ten := account % 10
		account -= ten

		if last > ten {
			fmt.Print(-1 * (account + ten))
		} else {
			fmt.Print(-1 * (account + last))
		}
	}
}
