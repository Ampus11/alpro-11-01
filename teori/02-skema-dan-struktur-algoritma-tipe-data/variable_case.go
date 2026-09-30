package main

import "fmt"

func main() {
	var name string

	name = "Rizky Fadlurrahman Ramadhan"
	fmt.Println("Nama : ", name)

	// name = 17
	// fmt.Println(name)

	var lastName = "Ramadhan"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Fadlurrahman"
	fmt.Println("Nama Tengah : ", middleName)

	var (
		fullName  = "Rizky Fadlurrahman Ramadhan"
		firstName = "Rizky"
	)
	fmt.Println("Nama Lengkap : ", fullName)
	fmt.Println("Nama Depan : ", firstName)
}
