package main

import "fmt"

func subtotal(harga, jumlah int) int {
	return harga * jumlah
}

func totalBayar(sub int) int {
	if sub >= 100000 {
		return sub - sub/10
	}
	return sub
}

func main() {
	var harga, jumlah int
	if dibaca, _ := fmt.Scan(&harga, &jumlah); dibaca != 2 {
		return
	}

	sub := subtotal(harga, jumlah)
	fmt.Println(totalBayar(sub))
}