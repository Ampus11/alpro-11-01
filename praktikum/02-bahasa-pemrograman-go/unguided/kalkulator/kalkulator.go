package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	sisa := a % b

	fmt.Printf("%d %d %d %d %d\n", tambah, kurang, kali, bagi, sisa)
}