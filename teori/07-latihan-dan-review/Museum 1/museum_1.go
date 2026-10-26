package main

import "fmt"

func main() {
	var p, prev, count, naik int
	var total float64

	fmt.Scan(&p)

	for p >= 0 && p <= 200  {
		total += float64(p)
		count++
		
		if count > 1 && p > prev {
			naik++
		}
		prev = p
		fmt.Scan(&p)
	}
	if count == 0 {
		fmt.Println(0, 0)
	} else {
		fmt.Println(naik, total/float64(count))
	}
}