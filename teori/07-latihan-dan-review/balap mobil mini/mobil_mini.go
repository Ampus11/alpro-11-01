package main

import "fmt"

func main() {
	var mobil string
	countA, countB := 0, 0
	winA, winB := 0, 0
	pemenang := ""

	for i := 0; i < 10; i++ {
		fmt.Scan(&mobil)
		if mobil == "A" {
			countA++
			winA++
		} else if mobil == "B" {
			countB++
			winB++
		}
		if winA == 4 && pemenang == "" {
			pemenang = "A"
		} else if winB == 4 && pemenang == "" {
			pemenang = "B"
		}
	}

	if countA != 5 || countB != 5 {
		fmt.Println("tidak valid")
	} else {
		fmt.Println(pemenang)
	}
}