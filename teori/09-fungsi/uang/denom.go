package main

import "fmt"

func denom(n int) (int, int, int) {
	return n / 10000, (n % 10000) / 2000, (n % 2000) / 1000
}

func main() {
	var n int
	if dibaca, _ := fmt.Scan(&n); dibaca != 1 {
		return
	}

	a, b, c := denom(n)
	fmt.Println(a, b, c)
}