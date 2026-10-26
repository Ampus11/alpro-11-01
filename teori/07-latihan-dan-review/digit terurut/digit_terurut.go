package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	asc, desc := true, true
	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			asc = false
		}
		if str[i] < str[i+1] {
			desc = false
		}
	}
	if asc {
		fmt.Println("ascending")
	} else if desc {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}
}