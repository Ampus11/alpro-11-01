package main

import "fmt"

func hitungLuas(alas, tinggi int, luas *float64) {
	*luas = float64(alas) * float64(tinggi) / 2
}

func main() {
	var n, a, t int
	var hasil float64

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &t)
		hitungLuas(a, t, &hasil)
		fmt.Println(hasil)
	}
}