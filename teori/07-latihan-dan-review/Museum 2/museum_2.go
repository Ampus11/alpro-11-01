package main

import "fmt"

func main() {
	total := 0
	for i := 1; i <= 5; i++ {
		var p int

		fmt.Printf("Hari ke-%d : ", i)
		fmt.Scan(&p)
		
		for p < 0 || p > 200 {
			fmt.Scan(&p)
		}
		total += p
	}
	fmt.Printf("Jumlah pengunjung : %d orang\n", total)
}