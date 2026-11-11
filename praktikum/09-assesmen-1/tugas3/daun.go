package main

import "fmt"

func main() {
	var n, lebar int
	fmt.Scan(&n)
	fmt.Scan(&lebar)
	tersempit := lebar

	for i := 1; i < n; i++ {
		fmt.Scan(&lebar)
		if lebar < tersempit {
			tersempit = lebar
		}
	}

	fmt.Println(tersempit)
} 
