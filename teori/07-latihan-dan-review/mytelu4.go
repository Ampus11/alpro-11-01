package main

import "fmt"

func main() {
	gold, platinum, silver, total := 0, 0, 0, 0

	for total < 500 {
		var poin int
		fmt.Scan(&poin)
		for poin < 50 {
			fmt.Scan(&poin)
		}
		total += poin
		if poin > 200 {
			gold++
		} else if poin >= 100 {
			platinum++
		} else {
			silver++
		}
	}
	fmt.Println(gold, platinum, silver)
}