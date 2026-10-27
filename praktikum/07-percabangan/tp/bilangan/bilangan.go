package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Nol")
	} else {
		var jenis, tanda string

		if n%2 == 0 {
			jenis = "Genap"
		} else {
			jenis = "Ganjil"
		}

		if n > 0 {
			tanda = "Positif"
		} else {
			tanda = "Negatif"
		}

		fmt.Println(jenis + ", " + tanda)
	}
}