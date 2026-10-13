package main

import "fmt"

func main() {
	var n, iterasi int
	fmt.Scan(&n)

	for iterasi != n {
		iterasi++
		fmt.Print(iterasi, " ")
	}
}