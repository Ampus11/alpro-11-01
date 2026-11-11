
package main

import (
	"fmt"
)

func main() {
	var angka, sebelumnya int
	var total, jumlah int
	var tertinggi, terendah int

	fmt.Scan(&angka)

	tertinggi = angka
	terendah = angka
	sebelumnya = angka
	total = angka
	jumlah = 1

	for {
		_, err := fmt.Scan(&angka)
		if err != nil {
			break
		}

		if angka == sebelumnya {
			break
		}

		if angka > tertinggi {
			tertinggi = angka
		}

		if angka < terendah {
			terendah = angka
		}

		total += angka
		jumlah++
		sebelumnya = angka
	}

	rataRata := float64(total) / float64(jumlah)

	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rataRata)
}