package main

import "fmt"

const NMAX = 100

type book struct {
	code, title string
	borrowed    int
	pages       int
}

type books [NMAX]book

// membaca data buku sampai kode "SELESAI"
func bacaData(T *books, n *int) {
	var kode string
	*n = 0
	for {
		if _, err := fmt.Scan(&kode); err != nil || kode == "SELESAI" {
			break
		}
		var b book
		b.code = kode
		fmt.Scan(&b.title, &b.borrowed, &b.pages)
		if *n < NMAX {
			T[*n] = b
			*n++
		}
	}
}

func cetakData(T books, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", T[i].code, T[i].title, T[i].borrowed, T[i].pages)
	}
}

// IS : T[0..n-1] dan x terdefinisi
// FS : semua buku dengan borrowed < x terhapus dengan menggeser elemen
//      setelahnya ke kiri; urutan buku tersisa tetap; n berkurang
func hapusBuku(T *books, n *int, x int) {
	i := 0
	for i < *n {
		if T[i].borrowed < x {
			for j := i; j < *n-1; j++ {
				T[j] = T[j+1]
			}
			*n--
		} else {
			i++
		}
	}
}

func totalPinjam(T books, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].borrowed
	}
	return total
}

// indeks buku dengan peminjaman terbanyak; jika sama, pilih halaman lebih sedikit
func idxFavorit(T books, n int) int {
	idx := 0
	for i := 1; i < n; i++ {
		if T[i].borrowed > T[idx].borrowed ||
			(T[i].borrowed == T[idx].borrowed && T[i].pages < T[idx].pages) {
			idx = i
		}
	}
	return idx
}

func main() {
	var T books
	var n, x int

	bacaData(&T, &n)
	fmt.Scan(&x)

	fmt.Println("Data awal:")
	cetakData(T, n)

	hapusBuku(&T, &n, x)
	fmt.Printf("Data setelah penghapusan (peminjaman >= %d):\n", x)
	if n == 0 {
		fmt.Println("Tidak ada buku tersisa")
		return
	}
	cetakData(T, n)

	fmt.Printf("Total peminjaman: %d\n", totalPinjam(T, n))
	f := idxFavorit(T, n)
	fmt.Println("Buku paling sering dipinjam:")
	fmt.Printf("%s %s %d %d\n", T[f].code, T[f].title, T[f].borrowed, T[f].pages)
}