package main

import "fmt"

func urutDua(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}

func main() {
	for {
		var a, b int
		dibaca, _ := fmt.Scan(&a, &b)
		if dibaca != 2 || a == b {
			break
		}

		kecil, besar := urutDua(a, b)
		fmt.Println(kecil, besar)
	}
}