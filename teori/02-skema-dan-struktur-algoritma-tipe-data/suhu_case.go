package main 

import "fmt"

func main() {
	var umur int8
	var suhu float32

	//suhu = 36.3
	//umur = 10
	fmt.Scanf("%f", &suhu)
	fmt.Scanf("%d", &umur)

	fmt.Println("Umur: ", umur)
	fmt.Println("Suhu: ", suhu)
	fmt.Println("Alamat memori dari Var suhu: ", &suhu)
	fmt.Println("Alamat memori dari Var umur: ", &umur)
}