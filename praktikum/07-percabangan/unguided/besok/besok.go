package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	hariBulan := 0
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		hariBulan = 31
	case 4, 6, 9, 11:
		hariBulan = 30
	case 2:
		if kabisat {
			hariBulan = 29
		} else {
			hariBulan = 28
		}
	}

	if tanggal < 1 || tanggal > hariBulan {
		fmt.Println("Tidak valid")
		return
	}

	tanggal++
	if tanggal > hariBulan {
		tanggal = 1
		bulan++
		if bulan > 12 {
			bulan = 1
			tahun++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}