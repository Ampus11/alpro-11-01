package main

import "fmt"

func main() {
	var kelas, hari string
	var usia int
	fmt.Scan(&kelas, &usia, &hari)

	harga := 0
	switch kelas {
	case "Ekonomi":
		harga = 100000
	case "Bisnis":
		harga = 200000
	case "Eksekutif":
		harga = 350000
	default:
		fmt.Println("Kelas tidak valid")
		return
	}

	if hari == "Sabtu" || hari == "Minggu" {
		harga = harga * 120 / 100
	}

	if usia < 3 {
		harga = 0
	} else if usia <= 12 {
		harga = harga * 50 / 100
	} else if usia >= 60 {
		harga = harga * 70 / 100
	}

	fmt.Println(harga)
}