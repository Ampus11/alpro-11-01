package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	kpk := n
	for kpk%m != 0 {
		kpk += n
	}
	fmt.Println(kpk)
}