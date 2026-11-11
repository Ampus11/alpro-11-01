package main

import "fmt"

func main() {
	var c string
	m, k, t := 0, 0, 0

	fmt.Scan(&c)
	for c == "M" || c == "K" || c == "T" {
		if c == "M" {
			m++
		} else if c == "K" {
			k++
		} else {
			t++
		}
		fmt.Scan(&c)
	}

	fmt.Println("Tipe M =", m)
	fmt.Println("Tipe K =", k)
	fmt.Println("Tipe T =", t)
}