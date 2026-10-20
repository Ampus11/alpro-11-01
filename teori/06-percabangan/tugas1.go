package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai  < 0 || nilai > 100 {
		fmt.Println("Nilai tidak valid")
	} else if nilai < 50 {
		fmt.Println("Nilai E")
	} else if nilai < 70 {
		fmt.Println("Nilai C")
	} else if nilai < 85 {
		fmt.Println("Nilai B")
	} else if nilai < 100 {
		fmt.Println("Nilai A")
	} else {
		fmt.Println("Nilai A+")
	}
}