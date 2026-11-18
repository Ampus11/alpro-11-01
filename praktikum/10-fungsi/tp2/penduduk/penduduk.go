package main

import (
	"fmt"
)

const NMAX int = 1000

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

type tabPenduduk [NMAX]penduduk

func cariTahun(T tabPenduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahun == th {
			*idx = i
			break
		}
	}
}

func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]
	for i := 1; i < n; i++ {
		if T[i].tahun < tua.tahun {
			*tua = T[i]
		}
		if T[i].tahun > muda.tahun {
			*muda = T[i]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabPenduduk
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}

	var thCari int
	fmt.Scan(&thCari)

	var idx int
	cariTahun(T, n, thCari, &idx)

	if idx != -1 {
		fmt.Printf("%s %d %s\n", T[idx].nama, T[idx].tahun, T[idx].kota)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	var tua, muda penduduk
	tertuaTermuda(T, n, &tua, &muda)

	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}