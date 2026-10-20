package main

import "fmt"

func main() {
	lamp_status := "Kunings"

	if lamp_status == "Merah" {
		fmt.Println("Berhenti")
	} else if lamp_status == "Kuning" {
		fmt.Println("Hati-hati / lebih pelan")
	} else if lamp_status == "Hijau" {
		fmt.Println("Jalan")
	} else {
		fmt.Print("Perintah tidak terdaftar disini")
	}
}