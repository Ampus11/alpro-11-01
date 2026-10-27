package main

import "fmt"

func main() {
	var a, b, c, maks int
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		maks = a
	} else if b >= a && b >= c {
		maks = b
	} else {
		maks = c
	}

	fmt.Println(maks)
}