package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var nama1, nama2 string
	nilai1, nilai2 := math.MaxInt, math.MaxInt

	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < nilai1 {
			nilai2, nama2 = nilai1, nama1
			nilai1, nama1 = nilai, nama
		} else if nilai < nilai2 {
			nilai2, nama2 = nilai, nama
		}
		i++
	}

	fmt.Println(nama1)
	fmt.Println(nama2)
}