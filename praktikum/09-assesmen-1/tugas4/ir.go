package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	hitung := 0

	for i := l; i <= r; i++ {
		n := i
		jumlah := 0
		kanan := 10
		tidakMenurun := true

		for n > 0 {
			d := n % 10
			if d > kanan {
				tidakMenurun = false
			}
			jumlah = jumlah + d
			kanan = d
			n = n / 10
		}

		if tidakMenurun {
			prima := true
			if jumlah < 2 {
				prima = false
			}
			for j := 2; j*j <= jumlah; j++ {
				if jumlah%j == 0 {
					prima = false
				}
			}
			if prima {
				hitung++
			}
		}
	}

	fmt.Println(hitung)
}