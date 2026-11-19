package main

import "fmt"

const NMAX int = 10000

type Arr [NMAX]int

func bacaData(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func geserKanan(T *Arr, n int) {
	terakhir := T[n-1]
	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}
	T[0] = terakhir
}

func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
	}
	fmt.Println()
}

func main() {
	var T Arr
	var n, k int
	fmt.Scan(&n, &k)
	bacaData(&T, n)

	for i := 0; i < k%n; i++ {
		geserKanan(&T, n)
	}
	cetak(T, n)
}