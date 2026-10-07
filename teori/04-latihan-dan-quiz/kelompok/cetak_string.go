package main

import "fmt"

func main() {
	var n, i int
	var s string
	fmt.Scan(&n, &s)
	for i = 1; i <= n; i++ {
		fmt.Println(s)
	}
}
