package main

import "fmt"

func main() {
	var r float64
	var luas float64
	pi := 22.0 / 7.0

	fmt.Scan(&r)
	luas = 4 * pi * r * r
	fmt.Println(luas)
}