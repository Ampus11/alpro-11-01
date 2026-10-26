package main

import "fmt"

func main() {
	gold, platinum, silver := 0, 0, 0
	totG, totP, totS := 0, 0, 0
	total := 0

	for total < 500 {
		var poin int
		fmt.Scan(&poin)
		for poin < 50 {
			fmt.Scan(&poin)
		}
		total += poin
		if poin > 200 {
			gold++
			totG += poin
		} else if poin >= 100 {
			platinum++
			totP += poin
		} else {
			silver++
			totS += poin
		}
	}
	avgG, avgP, avgS := 0.0, 0.0, 0.0
	if gold > 0 {
		avgG = float64(totG) / float64(gold)
	}
	if platinum > 0 {
		avgP = float64(totP) / float64(platinum)
	}
	if silver > 0 {
		avgS = float64(totS) / float64(silver)
	}
	fmt.Println(avgG, avgP, avgS)
}