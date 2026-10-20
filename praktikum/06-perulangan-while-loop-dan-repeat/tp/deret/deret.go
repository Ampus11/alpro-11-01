package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	i := 1
	jumlah := 0
	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println(jumlah)
}