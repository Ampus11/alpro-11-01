package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fpb := 1
	for i := 1; i <= n && i <= m; i++ {
		if n%i == 0 && m%i == 0 {
			fpb = i
		}
	}
	fmt.Println(fpb)
}