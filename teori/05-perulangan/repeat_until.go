package main

import "fmt"

func main() {
	var n, iterasi int
	fmt.Scan(&n)

	for {
		iterasi++
		fmt.Print(iterasi, " ")
		if iterasi == n {
			break
		}
	}
}