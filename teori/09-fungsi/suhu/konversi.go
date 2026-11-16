package main

import "fmt"

func konversi(c float64) (float64, float64, float64) {
	return 0.8 * c, 1.8*c + 32, c + 273.15
}

func main() {
	var c float64
	if dibaca, _ := fmt.Scan(&c); dibaca != 1 {
		return
	}

	r, f, k := konversi(c)
	fmt.Printf("%.2f %.2f %.2f\n", r, f, k)
}