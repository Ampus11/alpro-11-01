package main

import "fmt"

func main() {
	var p, q int
	fmt.Scan(&p, &q)
	ganjilDanGanjil := (p%2 != 0) && (q%2 != 0)
	genapAtauGenap := !ganjilDanGanjil
	tidakSama := !(p == q)
	fmt.Println(genapAtauGenap, ganjilDanGanjil, tidakSama)
}