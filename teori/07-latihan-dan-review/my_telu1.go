package main

import "fmt"

func main() {
	var poin int
	fmt.Scan(&poin)

	if poin > 200 {
		fmt.Println("Gold user")
	} else if poin >= 100 {
		fmt.Println("Platinum user")
	} else if poin >= 50 {
		fmt.Println("Silver user")
	}
}